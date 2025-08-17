# gox

golang extend tools, including config, logger, try-catch, slice shortcut operation and so on.

## Installation

- Required go 1.21+

```
go get -u github.com/winterant/gox
```

## Usage & Samples

**Use config**:
```go
package main
import (
    "fmt"
    "github.com/winterant/gox/pkg/xconfig"
)
func main() {
    conf := xconfig.LoadYaml("./conf/app.yaml", &App, "APP")
    fmt.Println(App.Log.Level)
    fmt.Println(conf.GetInt("log.maxDays"))
}
```

**Use logger**:
```go
package main
import (
    "context"
    "github.com/winterant/gox/pkg/xlog"
)
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
		Writer: logWriter,  // Optional
	})
    logger.Info(ctx, "hello, world. I am %s", userName)

	// Hook std log to file
	xlog.HookStdout(context.TODO(), "./log/stdout-")
	xlog.HookStderr(context.TODO(), "./log/stderr-") // include panic

	fmt.Println("hello, world! fmt.Println")       // --> ./log/stdout-*.log
	fmt.Fprintln(os.Stdout, "hello, stdout")       // --> ./log/stdout-*.log
	panic("hello, world! panic")                   // --> ./log/stderr-*.log
}
```

**Use try-catch**:
```go
package main
import (
    "github.com/winterant/gox/pkg/x"
)
func main() {
    // Use x to catch panic with runtime stack info just like try-catch in other language
    defer x.TryCatch(func(e error) {
        xlog.Error(ctx, "panic: %+v", e)
    })
    // ... other codes which may panic ...
}
```

**Check context canceled**:
```go
    // Use x to check context done
    if x.CtxDone(ctx) {
        fmt.Println("context canceled")
    }
```

**other shortcut operations**:
```go
    m := map[string]int{"a": 1, "b": 2}
    // Use x to get map keys
    keys := x.MapKeys(m)  // ["a", "b"]
```

More samples in [./examples/](./examples/)


## Recommended packages

| Function        | Package                                                                          | Install                                            |
|-----------------|----------------------------------------------------------------------------------|----------------------------------------------------|
| struct default  | [github.com/creasty/defaults](https://github.com/creasty/defaults)               | `go get -u github.com/creasty/defaults`            |
| struct validate | [github.com/go-playground/validator](https://github.com/go-playground/validator) | `go get -u github.com/go-playground/validator/v10` |
| Http Client     | [github.com/go-resty/resty](https://github.com/go-resty/resty)                   | `go get -u github.com/go-resty/resty/v3`           |
