package gin

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"

	"github.com/coalaura/logger"
	"github.com/gin-gonic/gin"
)

func TestMiddleware(t *testing.T) {
	l := logger.New()

	gin.SetMode(gin.ReleaseMode)

	m := GinMiddleware(l)
	c, _ := gin.CreateTestContext(logger.DiscardWriter{})
	c.Request, _ = http.NewRequest("GET", "http://localhost/test/path", nil)

	codes := []int{
		100,
		200,
		300,
		400,
		500,
		218,
		418,
		420,
		999,
	}

	for i, method := range []string{"GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"} {
		c.Request.Method = method

		c.Request.RemoteAddr = fmt.Sprintf("%d.%d.%d.%d:%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(65535))
		c.Status(codes[i])

		m(c)
	}
}
