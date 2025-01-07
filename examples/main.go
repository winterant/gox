package main

import (
	"context"
	"flag"
	"github.com/winterant/gox/pkg/config"
	"github.com/winterant/gox/pkg/logger"
	"github.com/winterant/gox/pkg/x"
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

	_ = config.LoadConfig(*AppConfPath, &App, "APP")

	logger.InitDefault(App.Log.Path, App.Log.MaxSizeMB, App.Log.MaxBackups, App.Log.MaxDays, App.Log.Level)
	ctx := logger.ContextWithArgs(context.Background(), "app-name", "my-first-app")

	return ctx
}

func main() {
	ctx := MustInit()

	logger.Debug(ctx, "global config: %+v", App)

	x.TryCatch(func(e error) {
		logger.Error(ctx, "panic: %+v", e)
	})

	keys := x.MapKeys(map[string]int{"a": 1, "b": 2})
	logger.Info(ctx, "%+v", keys)

	log2 := logger.New(os.Stdout, "debug")
	log2.Info(ctx, "hello this is a custom logger")

	panic("panic test")
}
