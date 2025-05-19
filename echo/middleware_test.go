package echo

import (
	"fmt"
	"math/rand"
	"net/http/httptest"
	"testing"

	"github.com/coalaura/logger"
	"github.com/labstack/echo/v4"
)

func TestEchoMiddleware(t *testing.T) {
	l := logger.New()

	e := echo.New()
	e.Use(MiddleWare(l))

	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"}

	for _, method := range methods {
		t.Run(method, func(t *testing.T) {
			req := httptest.NewRequest(method, "http://localhost/test/path", nil)
			req.RemoteAddr = fmt.Sprintf("%d.%d.%d.%d:%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(65535))

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("/test/path")

			e.ServeHTTP(rec, req)
		})
	}
}
