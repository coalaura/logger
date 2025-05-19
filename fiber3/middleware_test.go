package fiber3

import (
	"fmt"
	"math/rand"
	"net/http/httptest"
	"testing"

	"github.com/coalaura/logger"
	"github.com/gofiber/fiber/v2"
)

func TestMiddleware(t *testing.T) {
	l := logger.New()

	app := fiber.New()
	app.Use(Middleware(l))

	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"}

	for _, method := range methods {
		t.Run(method, func(t *testing.T) {
			req := httptest.NewRequest(method, "http://localhost/", nil)
			req.RemoteAddr = fmt.Sprintf("%d.%d.%d.%d:%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(65535))

			_, err := app.Test(req, -1)
			if err != nil {
				t.Fatalf("Failed to execute test request: %v", err)
			}
		})
	}
}
