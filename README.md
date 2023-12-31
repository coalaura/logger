![banner](.github/banner.png)

# Logger

Simple easy to use logger for golang. Supports colors and different log levels.

```go
package main

import (
	"github.com/coalaura/logger"
)

func main() {
    // Default Logger
    log := logger.New()

    // Custom output
    log.WithOutput(os.Stdout)

    // Options
    log.WithOptions(logger.Options{
        NoColor: false,
        NoLevel: false,
        NoTime:  false,
    })

    log.Debug("This is a Debug Message")
    log.Note("This is an Note Message")
    log.Info("This is an Info Message")
    log.Warning("This is a Warning Message")
    log.Error("This is an Error Message")
    log.Fatal("This is a Fatal Message")

    // Gin Middleware
    // r := gin.Default()
    r.Use(log.Middleware())
}
```

![test](.github/test.png)