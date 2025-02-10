# gox

golang extend tools, including config, logger, try-catch, slice shortcut operation and so on.

## Installation

```
go get -u github.com/winterant/gox
```

## Usage & Samples

Use config:
```go
package main
import (
    "fmt"
    "github.com/winterant/gox/pkg/xconfig"
)
func main() {
    // Load config file
    _ = xconfig.LoadYaml("./conf/app.yaml", &App, "APP")
    fmt.Println(App.Log.Name)
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
    // Init logger
    xlog.InitDefault("./log/main.log", 128, 100, 90, "DEBUG")
    ctx := xlog.ContextWithArgs(context.Background(), "app-name", "example") // add context args which will print in log

    // Use logger
    xlog.Info(ctx, "hello %s", "world")
    xlog.Error(ctx, "error %s", "example")
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
