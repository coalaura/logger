package gin

import (
	"net/http"
	"time"

	"github.com/coalaura/logger"
	"github.com/gin-gonic/gin"
)

type GinAdapter struct {
	ctx *gin.Context

	timeTaken time.Duration
}

func GinMiddleware(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		log.LogHTTPRequest(&GinAdapter{
			ctx:       c,
			timeTaken: time.Since(start),
		})
	}
}

func (a *GinAdapter) ClientIP() string {
	return a.ctx.ClientIP()
}

func (a *GinAdapter) StatusCode() int {
	return a.ctx.Writer.Status()
}

func (a *GinAdapter) TimeTaken() time.Duration {
	return a.timeTaken
}

func (a *GinAdapter) Request() *http.Request {
	return a.ctx.Request
}
