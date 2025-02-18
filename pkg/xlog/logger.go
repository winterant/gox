package xlog

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"
)

type Logger struct {
	*slog.Logger
}

func New(writer io.Writer, level string) *Logger {
	sLevel := getSlogLevel(level)
	return &Logger{
		Logger: slog.New(newPrettyHandler(withWriter(writer), withLever(sLevel), withCallerDepth(2), withCodeSource(true))),
	}
}

func (l *Logger) Debug(ctx context.Context, format string, args ...any) {
	l.log(ctx, slog.LevelDebug, format, args...)
}

func (l *Logger) Info(ctx context.Context, format string, args ...any) {
	l.log(ctx, slog.LevelInfo, format, args...)
}

func (l *Logger) Warn(ctx context.Context, format string, args ...any) {
	l.log(ctx, slog.LevelWarn, format, args...)
}

func (l *Logger) Error(ctx context.Context, format string, args ...any) {
	l.log(ctx, slog.LevelError, format, args...)
}

func (l *Logger) log(ctx context.Context, level slog.Level, format string, args ...any) {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args...)
	}
	l.Log(ctx, level, format)
}

func getSlogLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		panic("log level must be one of debug, info, warn, error. But got " + level)
	}
}
