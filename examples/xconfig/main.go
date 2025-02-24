package main

import (
	"flag"
	"fmt"
	"github.com/winterant/gox/pkg/xconfig"
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

var App app

var AppConfPath = flag.String("conf", "examples/xconfig/app.yaml", "app config path")

func main() {
	flag.Parse()

	conf := xconfig.LoadYaml(*AppConfPath, &App, "APP")

	fmt.Println(App.Log.Level)
	fmt.Println(conf.GetString("log.path"))
	fmt.Println(conf.GetInt("log.maxDays"))
}

/*
debug
./log/main.log
90
*/
