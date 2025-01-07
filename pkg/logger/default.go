package logger

import (
	"context"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log/slog"
	"os"
)

var defaultLogger *Logger

func InitDefault(path string, maxSizeMB, maxBackups, maxDays int, level string) {
	var logWriter io.Writer = &lumberjack.Logger{
		Filename:   path,       // defaultLog file path
		MaxSize:    maxSizeMB,  // file max size in MB
		MaxBackups: maxBackups, // max number of backup defaultLog files
		MaxAge:     maxDays,    // max number of days to keep old files
		Compress:   false,      // whether to compress/archive old files
		LocalTime:  true,       // Use local time or not
	}
	sLevel := getSlogLevel(level)
	if sLevel == slog.LevelDebug {
		logWriter = io.MultiWriter(logWriter, os.Stdout)
	}
	handler := NewPrettyHandler(withWriter(logWriter), withLever(sLevel), withCallerDepth(3), withCodeSource(true))
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
