package echo

import (
	"time"

	"github.com/coalaura/logger"
	"github.com/labstack/echo/v4"
)

type EchoAdapter struct {
	ctx       echo.Context
	timeTaken time.Duration
}

func EchoMiddleWare(log *logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			start := time.Now()

			err = next(c)

			log.LogHTTPRequest(&EchoAdapter{
				ctx:       c,
				timeTaken: time.Since(start),
			})

			return
		}
	}
}

func (e *EchoAdapter) Method() string {
	return e.ctx.Request().Method
}

func (e *EchoAdapter) Path() string {
	return e.ctx.Request().URL.Path
}

func (e *EchoAdapter) ClientIP() string {
	return e.ctx.RealIP()
}

func (e *EchoAdapter) StatusCode() int {
	return e.ctx.Response().Status
}

func (e *EchoAdapter) TimeTaken() time.Duration {
	return e.timeTaken
}
