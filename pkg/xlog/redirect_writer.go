package xlog

import (
	"context"
	"log/slog"
	"strings"
)

type redirectWriter struct {
	*slog.Logger
	level slog.Level
}

func (w *redirectWriter) Write(p []byte) (n int, err error) {
	w.Logger.Log(context.Background(), w.level, strings.TrimSpace(string(p)))
	return len(p), nil
}
