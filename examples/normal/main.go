package main

import (
	"context"
	"flag"
	"os"

	"github.com/winterant/gox/pkg/x"
	"github.com/winterant/gox/pkg/xconfig"
	"github.com/winterant/gox/pkg/xlog"
)

type app struct {
	Log log
}

type log struct {
	Level      string
	Path       string
	MaxSizeMB  int
	MaxBackups int
	MaxDays    int
}

var AppConfPath = flag.String("conf", "./examples/conf/app.yaml", "app config path")

var App app

func MustInit() context.Context {
	flag.Parse()

	_ = xconfig.LoadYaml(*AppConfPath, &App, "APP")

	xlog.InitDefault(App.Log.Path, App.Log.MaxSizeMB, App.Log.MaxBackups, App.Log.MaxDays, App.Log.Level)
	ctx := xlog.ContextWithArgs(context.Background(), "app-name", "my-first-app")

	return ctx
}

func main() {
	ctx := MustInit()

	xlog.Debug(ctx, "global config: %+v", App)

	// custom logger
	customLogger := xlog.New(os.Stdout, "debug")
	customLogger.Info(ctx, "hello this is a custom logger")

	// catch panic, must defer it
	defer x.TryCatch(func(e error) {
		xlog.Error(ctx, "panic: %+v", e)
	})

	keys := x.MapKeys(map[string]int{"a": 1, "b": 2})
	xlog.Info(ctx, "%+v", keys)

	panic("panic test")
}

/*
2025-02-10 19:44:51.801049 DEBUG /Users/zhaojinglong01/Personal/Projects/gox/examples/normal/main.go:43 [app-name=my-first-app] global config: {Log:{Level:debug Path:./log/main.log MaxSizeMB:128 MaxBackups:30 MaxDays:90}}
2025-02-10 19:44:51.801480 INFO /Users/zhaojinglong01/Personal/Projects/gox/examples/normal/main.go:47 [app-name=my-first-app] hello this is a custom logger
2025-02-10 19:44:51.801491 INFO /Users/zhaojinglong01/Personal/Projects/gox/examples/normal/main.go:55 [app-name=my-first-app] [a b]
2025-02-10 19:44:51.801564 ERROR /Users/zhaojinglong01/Personal/Projects/gox/examples/normal/main.go:51 [app-name=my-first-app] panic: panic test
        /Users/zhaojinglong01/Personal/Projects/gox/examples/normal/main.go:57 main.main
        /Library/env/goenv/versions/1.21.13/src/runtime/proc.go:267 runtime.main
        /Library/env/goenv/versions/1.21.13/src/runtime/asm_arm64.s:1197 runtime.goexit
*/
