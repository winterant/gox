package xlog

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/creasty/defaults"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*slog.Logger
}

type Option struct {
	Level      string `default:"info"`
	Writer     io.Writer
	Stdout     bool   `default:"false"`             // 仅当 Writer 为空时生效
	Path       string `default:"./log/default.log"` // 仅当 Writer 为空时生效
	MaxSizeMB  int    `default:"64"`                // 仅当 Writer 为空时生效
	MaxBackups int    `default:"100"`               // 仅当 Writer 为空时生效
	MaxDays    int    `default:"90"`                // 仅当 Writer 为空时生效
}

func New(opt Option) *Logger {
	_ = defaults.Set(&opt)

	if opt.Writer == nil {
		opt.Writer = &lumberjack.Logger{
			Filename:   opt.Path,       // defaultLog file path
			MaxSize:    opt.MaxSizeMB,  // file max size in MB
			MaxBackups: opt.MaxBackups, // max number of backup defaultLog files
			MaxAge:     opt.MaxDays,    // max number of days to keep old files
			Compress:   false,          // whether to compress/archive old files
			LocalTime:  true,           // Use local time or not
		}
		if opt.Stdout {
			opt.Writer = io.MultiWriter(opt.Writer, os.Stdout)
		}
	}

	//  收集log打的日志 2025/06/19 11:03:06.242132 /data/main.go:21: hello world
	log.SetOutput(opt.Writer)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)

	sLevel := getSlogLevel(opt.Level)
	return &Logger{
		Logger: slog.New(newPrettyHandler(withWriter(opt.Writer), withLever(sLevel), withCallerDepth(2), withCodeSource(true))),
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
