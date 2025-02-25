# gox

golang extend tools, including config, logger, try-catch, slice shortcut operation and so on.

## Installation

```
go get -u github.com/winterant/gox
```
- Required go 1.21+

## Usage & Samples

Use config:
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

Use log:
```go
package main
import (
    "context"
    "github.com/winterant/gox/pkg/xlog"
)
func main() {
	// Use default logger
    xlog.InitDefault("./log/main.log", 128, 100, 90, "DEBUG")
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
    logger := xlog.New(logWriter, "DEBUG")
    logger.Info(ctx, "hello, world. I am %s", userName)
}
```

Use try-catch:
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

Check context canceled:
```go
package main
import (
    "github.com/winterant/gox/pkg/x"
)
func main() {
    // Use x to check context done
    if x.CtxDone(ctx) {
        fmt.Println("context canceled")
    }
}
```

other shortcut operations:
```go
package main
import (
    "github.com/winterant/gox/pkg/x"
)
func main() {
    m := map[string]int{"a": 1, "b": 2}
    // Use x to get map keys
    keys := x.MapKeys(m)  // ["a", "b"]
}
```

More samples in [./examples/](./examples/)


## Recommended packages

| Function    | Package                                                        | Install                                  |
|-------------|----------------------------------------------------------------|------------------------------------------|
| Http Client | [github.com/go-resty/resty](https://github.com/go-resty/resty) | `go get -u github.com/go-resty/resty/v3` |
