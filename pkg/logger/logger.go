package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *slog.Logger

func MustInit(path string, maxSizeMB, maxBackups, maxDays int, level string) {
	var logWriter io.Writer = &lumberjack.Logger{
		Filename:   path,       // log file path
		MaxSize:    maxSizeMB,  // file max size in MB
		MaxBackups: maxBackups, // max number of backup log files
		MaxAge:     maxDays,    // max number of days to keep old files
		Compress:   false,      // whether to compress/archive old files
		LocalTime:  true,       // Use local time or not
	}
	sLevel := getSlogLevel(level)
	if sLevel == slog.LevelDebug {
		logWriter = io.MultiWriter(logWriter, os.Stdout)
	}
	handler := newPrettyHandler(withWriter(logWriter), withLever(sLevel), withCallerDepth(2), withCodeSource(true))
	logger = slog.New(handler)
}

func Debug(ctx context.Context, format string, args ...any) {
	log(ctx, slog.LevelDebug, format, args...)
}

func Info(ctx context.Context, format string, args ...any) {
	log(ctx, slog.LevelInfo, format, args...)
}

func Warn(ctx context.Context, format string, args ...any) {
	log(ctx, slog.LevelWarn, format, args...)
}

func Error(ctx context.Context, format string, args ...any) {
	log(ctx, slog.LevelError, format, args...)
}

func log(ctx context.Context, level slog.Level, format string, args ...any) {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args...)
	}
	logger.Log(ctx, level, format)
}

func getSlogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
