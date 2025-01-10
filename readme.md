# gox

golang extend tools, including config, logger, try-catch, slice shortcut operation and so on.

## Installation

```
go get -u github.com/winterant/gox
```

## Usage & Samples

```go
package main

import (
    "context"
    "github.com/winterant/gox/pkg/x"
    "github.com/winterant/gox/pkg/xconfig"
    "github.com/winterant/gox/pkg/xlog"
)

func main() {
    // Load config file
    _ = xconfig.LoadYaml("./conf/app.yaml", &App, "APP")

    // Init logger
    xlog.InitDefault(App.Log.Path, App.Log.MaxSizeMB, App.Log.MaxBackups, App.Log.MaxDays, App.Log.Level)
    ctx := xlog.ContextWithArgs(context.Background(), "app-name", "example") // add context args which will print in log

    // Use logger
    xlog.Info(ctx, "hello %s", "world")

    // Use x to catch panic with runtime stack info just like try-catch in other language
    defer x.TryCatch(func(e error) {
        xlog.Error(ctx, "panic: %+v", e)
    })
    // ... other codes which may panic ...

    // Use x to get map keys
    keys := x.MapKeys(map[string]int{"a": 1, "b": 2})
    xlog.Info(ctx, "%+v", keys)

    // Use x to check context done
    if x.CtxDone(ctx) {
        xlog.Warn(ctx, "context canceled")
    }
}
```

More samples in [./examples/](./examples/)
