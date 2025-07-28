package main

import (
	"context"
	"fmt"
	"github.com/winterant/gox/pkg/xlog"
	"log"
	"os"
)

func main() {
	os.MkdirAll("./log/", 0755)
	xlog.InitDefault(xlog.Option{
		Path: "./log/main.log",
	})
	xlog.HookStdout(context.TODO(), "./log/stdout-")
	xlog.HookStderr(context.TODO(), "./log/stderr-") // include panic

	xlog.Info(nil, "hello, world! I am xlog.Info") // --> ./log/test.log
	log.Println("hello, log.Println")              // --> ./log/test.log
	fmt.Println("hello, world! fmt.Println")       // --> ./log/stdout-*.log
	fmt.Fprintln(os.Stdout, "hello, stdout")       // --> ./log/stdout-*.log
	panic("hello, world! panic")                   // --> ./log/stderr-*.log
}
