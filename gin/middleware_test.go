package gin

import (
	"fmt"
	"math/rand"
	"net/http/httptest"
	"testing"

	"github.com/coalaura/logger"
	"github.com/gin-gonic/gin"
)

func TestMiddleware(t *testing.T) {
	l := logger.New()

	app := gin.New()
	app.Use(Middleware(l))

	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"}

	for _, method := range methods {
		t.Run(method, func(t *testing.T) {
			req := httptest.NewRequest(method, "http://localhost/test/path", nil)
			req.RemoteAddr = fmt.Sprintf("%d.%d.%d.%d:%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(65535))

			app.ServeHTTP(httptest.NewRecorder(), req)
		})
	}
}
