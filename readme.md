# gox

golang extend tools

## Installation

```
go get -u github.com/winterant/gox
```

## Usage

```go
package main

import (
	"context"
	"github.com/winterant/gox/pkg/config"
	"github.com/winterant/gox/pkg/logger"
	"github.com/winterant/gox/pkg/x"
)

func main() {
    // Load config file
    _ = config.LoadConfig("./conf/app.yaml", &App, "APP")

    // Init logger
    logger.MustInit(App.Log.Path, App.Log.MaxSizeMB, App.Log.MaxBackups, App.Log.MaxDays, App.Log.Level)
    ctx := logger.ContextWithArgs(context.Background(), "app-name", "example") // add context args which will print in log

    // Use logger
    logger.Info(ctx, "hello %s", "world")

    // Use x to catch panic with runtime stack info just like try-catch in other language
    defer x.TryCatch(func(e error) {
        logger.Error(ctx, "panic: %+v", e)
    })
    // ... other codes which may panic ...

}
```

## Samples

[./examples/main.go](./examples/main.go)

## License