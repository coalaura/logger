package fiber3

import (
	"time"

	"github.com/coalaura/logger"
	"github.com/gofiber/fiber/v3"
)

type FiberAdapter struct {
	ctx fiber.Ctx

	timeTaken time.Duration
}

func Middleware(log *logger.Logger) fiber.Handler {
	return func(c fiber.Ctx) (err error) {
		start := time.Now()

		err = c.Next()

		log.LogHTTPRequest(&FiberAdapter{
			ctx:       c,
			timeTaken: time.Since(start),
		})

		return
	}
}

func (a *FiberAdapter) Method() string {
	return a.ctx.Method()
}

func (a *FiberAdapter) Path() string {
	return a.ctx.Path()
}

func (a *FiberAdapter) ClientIP() string {
	return a.ctx.IP()
}

func (a *FiberAdapter) StatusCode() int {
	return a.ctx.Response().StatusCode()
}

func (a *FiberAdapter) TimeTaken() time.Duration {
	return a.timeTaken
}
