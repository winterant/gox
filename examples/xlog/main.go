package main

import (
	"context"
	"io"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/winterant/gox/pkg/xlog"
)

const userName = "Winterant"

func main() {
	// Use default logger
	xlog.InitDefault(xlog.Option{
		Level:  "debug",
		Stdout: true,
		Path:   "./log/main.log",
	})
	ctx := xlog.ContextWithArgs(context.Background(), "appName", "my-example-app") // add context args which will print in log
	xlog.Info(ctx, "hello, world")
	xlog.Error(ctx, "I am %s", userName)

	// Use custom logger
	logWriter := io.MultiWriter(&lumberjack.Logger{
		Filename:   "./log/my.log", // defaultLog file path
		MaxSize:    128,            // file max size in MB
		MaxBackups: 100,            // max number of backup defaultLog files
		MaxAge:     90,             // max number of days to keep old files
		Compress:   false,          // whether to compress/archive old files
		LocalTime:  true,           // Use local time or not
	}, os.Stdout)
	logger := xlog.New(xlog.Option{
		Writer: logWriter,
	})
	logger.Info(ctx, "hello, world. I am %s", userName)
}

/* Stdout:
2025-02-24 11:38:33.628920 INFO /Users/zhaojinglong01/Personal/Projects/gox/examples/xlog/main.go:17 [appName=my-example-app] hello, world
2025-02-24 11:38:33.629338 ERROR /Users/zhaojinglong01/Personal/Projects/gox/examples/xlog/main.go:18 [appName=my-example-app] I am Winterant
2025-02-24 11:38:33.629359 INFO /Users/zhaojinglong01/Personal/Projects/gox/examples/xlog/main.go:30 [appName=my-example-app] hello, world. I am Winterant
*/
