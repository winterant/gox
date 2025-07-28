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
	Level string `default:"info"`

	Writer     io.Writer ``                            // Custom writer for log output
	Path       string    `default:"./log/default.log"` // Effective only when Writer is nil
	MaxSizeMB  int       `default:"64"`                // Effective only when Writer is nil
	MaxBackups int       `default:"500"`               // Effective only when Writer is nil
	MaxDays    int       `default:"90"`                // Effective only when Writer is nil

	Stdout bool `default:"false"` // Output the log to stdout at the same time. Don't set to true if you use xlog.HookStdout

	callerDepth int `default:"2"`
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
	}
	if opt.Stdout {
		opt.Writer = io.MultiWriter(opt.Writer, os.Stdout)
	}

	baseLogger := slog.New(newPrettyHandler(withWriter(opt.Writer), withLever(opt.Level), withCallerDepth(opt.callerDepth), withCodeSource(true)))

	// 截获 标准库 log 打的日志
	log.SetFlags(0)
	log.SetPrefix("[go-std-log] ")
	log.SetOutput(&redirectWriter{Logger: baseLogger, level: slog.LevelInfo})

	return &Logger{Logger: baseLogger}
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
	l.Log(ctx, level, strings.TrimSpace(format))
}
