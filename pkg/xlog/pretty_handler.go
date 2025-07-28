package xlog

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime"
	"strings"
	"sync"
)

// prettyHandler implements a slog handler with pretty format as easy to view.
type prettyHandler struct {
	slog.Handler
	addSource   bool
	level       slog.Level
	callerDepth int
	w           io.Writer
	logAttrs    []slog.Attr
	mu          *sync.Mutex
}

type handlerOption func(*prettyHandler)

func withWriter(writer io.Writer) handlerOption {
	return func(handler *prettyHandler) {
		handler.w = writer
	}
}

func withLever(level string) handlerOption {
	var sLevel slog.Level
	switch strings.ToLower(level) {
	case "debug":
		sLevel = slog.LevelDebug
	case "info":
		sLevel = slog.LevelInfo
	case "warn":
		sLevel = slog.LevelWarn
	case "error":
		sLevel = slog.LevelError
	default:
		panic("log level must be one of debug, info, warn, error. But got " + level)
	}
	return func(handler *prettyHandler) {
		handler.level = sLevel
	}
}

func withCodeSource(addSource bool) handlerOption {
	return func(handler *prettyHandler) {
		handler.addSource = addSource
	}
}

func withCallerDepth(depth int) handlerOption {
	return func(handler *prettyHandler) {
		handler.callerDepth = depth
	}
}

func newPrettyHandler(options ...handlerOption) *prettyHandler {
	handler := prettyHandler{
		addSource:   true,
		level:       slog.LevelInfo,
		callerDepth: 0,
		w:           os.Stdout,
		mu:          &sync.Mutex{},
	}
	for _, option := range options {
		option(&handler)
	}
	handler.Handler = slog.NewJSONHandler(handler.w, &slog.HandlerOptions{AddSource: handler.addSource, Level: handler.level})
	return &handler
}

func (h *prettyHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.Handler.Enabled(ctx, level)
}

func (h *prettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h.logAttrs = append(h.logAttrs, attrs...)
	return h
}

func (h *prettyHandler) WithGroup(name string) slog.Handler {
	h.Handler = h.Handler.WithGroup(name)
	return h
}

func (h *prettyHandler) Handle(ctx context.Context, r slog.Record) error {
	buf := bytes.NewBuffer(make([]byte, 0, len(r.Message)+256))
	buf.WriteString(r.Time.Format("2006-01-02 15:04:05.000000"))
	buf.WriteString(" ")
	buf.WriteString(r.Level.String())

	if h.addSource {
		if _, file, line, ok := runtime.Caller(3 + h.callerDepth); ok {
			buf.WriteString(fmt.Sprintf(" %s:%d", file, line))
		}
	}

	formatArg := func(k, v any) {
		buf.WriteString(" [")
		if key, ok := k.(string); ok {
			buf.WriteString(key)
		} else {
			buf.WriteString(fmt.Sprintf("%v", k))
		}
		buf.WriteString("=")
		if val, ok := v.(string); ok {
			buf.WriteString(val)
		} else {
			buf.WriteString(fmt.Sprintf("%v", v))
		}
		buf.WriteString("]")
	}
	for _, attr := range h.logAttrs { // 创建slog.Logger时添加的参数
		formatArg(attr.Key, attr.Value.String())
	}
	ctxArgs := getContextArgs(ctx) // context中的参数
	for i := 0; i+1 < len(ctxArgs); i += 2 {
		formatArg(ctxArgs[i], ctxArgs[i+1])
	}
	if r.NumAttrs() > 0 {
		r.Attrs(func(attr slog.Attr) bool { // 打印日志时临时添加的参数
			formatArg(attr.Key, attr.Value.String())
			return true
		})
	}

	buf.WriteString(" ")
	buf.WriteString(r.Message)
	buf.WriteString("\n")

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.w.Write(buf.Bytes())
	return err
}
