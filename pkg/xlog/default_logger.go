package xlog

import (
	"context"
)

var defaultLogger *Logger

func InitDefault(opt Option) {
	opt.callerDepth = 3
	defaultLogger = New(opt)
}

func Debug(ctx context.Context, format string, args ...any) {
	defaultLogger.Debug(ctx, format, args...)
}

func Info(ctx context.Context, format string, args ...any) {
	defaultLogger.Info(ctx, format, args...)
}

func Warn(ctx context.Context, format string, args ...any) {
	defaultLogger.Warn(ctx, format, args...)
}

func Error(ctx context.Context, format string, args ...any) {
	defaultLogger.Error(ctx, format, args...)
}
