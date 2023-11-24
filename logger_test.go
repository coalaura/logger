package logger_v2

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLogger(t *testing.T) {
	l := New()

	l.PrintLn()

	fmt.Println("- Log Levels -")
	l.Debug("This is a Debug Message")
	l.Note("This is an Note Message")
	l.Info("This is an Info Message")
	l.Warning("This is a Warning Message")
	l.Error("This is an Error Message")
	l.Fatal("This is a Fatal Message")
	l.PrintLn("Just a normal message")

	l.PrintLn()

	fmt.Println("- Gin Middleware -")
	cycleRequests(l)

	l.PrintLn()
}

func cycleRequests(l *Logger) {
	gin.SetMode(gin.ReleaseMode)

	m := l.Middleware()
	c, _ := gin.CreateTestContext(discardWriter{})
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

type discardWriter struct{}

func (d discardWriter) Write(b []byte) (int, error) {
	return len(b), nil
}
func (d discardWriter) WriteHeader(_ int) {}
func (d discardWriter) Header() http.Header {
	return http.Header{
		"user-agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0"},
	}
}
