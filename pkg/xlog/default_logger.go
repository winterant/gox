package xlog

import (
	"context"
	"io"
	"log/slog"
	"os"

	"github.com/creasty/defaults"
	"gopkg.in/natefinch/lumberjack.v2"
)

var defaultLogger *Logger

func InitDefault(opt Option) {
	_ = defaults.Set(&opt)

	var logWriter io.Writer = &lumberjack.Logger{
		Filename:   opt.Path,       // defaultLog file path
		MaxSize:    opt.MaxSizeMB,  // file max size in MB
		MaxBackups: opt.MaxBackups, // max number of backup defaultLog files
		MaxAge:     opt.MaxDays,    // max number of days to keep old files
		Compress:   false,          // whether to compress/archive old files
		LocalTime:  true,           // Use local time or not
	}
	if opt.Stdout {
		logWriter = io.MultiWriter(logWriter, os.Stdout)
	}
	sLevel := getSlogLevel(opt.Level)
	handler := newPrettyHandler(withWriter(logWriter), withLever(sLevel), withCallerDepth(3), withCodeSource(true))
	defaultLogger = &Logger{
		Logger: slog.New(handler),
	}
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
