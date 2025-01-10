package main

import (
	"context"
	"flag"
	"github.com/winterant/gox/pkg/x"
	"github.com/winterant/gox/pkg/xconfig"
	"github.com/winterant/gox/pkg/xlog"
	"os"
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

var AppConfPath = flag.String("conf", "./conf/app.yaml", "app config path")

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
