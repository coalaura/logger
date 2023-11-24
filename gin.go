package logger

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Middleware returns a logger middleware for gin-gonic
func (l *Logger) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		timeTaken := time.Since(start)

		l._logRequest(c.Request.Method, c.ClientIP(), c.Request.URL.Path, c.Request.Header.Get("user-agent"), c.Writer.Status(), timeTaken)
	}
}

func (l *Logger) _logRequest(method, clientIp, path, userAgent string, statusCode int, timeTaken time.Duration) {
	l.mx.Lock()
	defer l.mx.Unlock()

	l.CPrint(l.date(), 243, 0)

	col := 52
	switch method {
	case "GET":
		col = 33
	case "HEAD":
		col = 141
	case "POST":
		col = 36
	case "PUT":
		col = 34
	case "DELETE":
		col = 160
	case "CONNECT":
		col = 45
	case "OPTIONS":
		col = 209
	case "TRACE":
		col = 162
	case "PATCH":
		col = 202
	}

	l.CPrint("[", 243, 0)
	l.CPrint(fmt.Sprintf("%-7s", strings.ToLower(method)), col, 0)
	l.CPrint("] ", 243, 0)

	l.CPrint("[", 243, 0)
	l.CPrint(fmt.Sprintf("%-5s ", _fmtDuration(timeTaken)), 115, 0)
	l.CPrint("] ", 243, 0)

	l.CPrint("[", 243, 0)
	l.CPrint(fmt.Sprintf("%-15s", clientIp), 248, 0)
	l.CPrint("] ", 243, 0)

	col = 111
	if statusCode >= 100 && statusCode < 200 {
		col = 159
	} else if statusCode >= 200 && statusCode < 300 {
		col = 34
	} else if statusCode >= 300 && statusCode < 400 {
		col = 184
	} else if statusCode >= 400 && statusCode < 500 {
		col = 202
	} else if statusCode >= 500 {
		col = 124
	}

	/*
	 * 218 This is fine
	 * 418 I'm a teapot
	 * 420 Enhance Your Calm
	 */
	if statusCode == 218 || statusCode == 418 || statusCode == 420 {
		col = 210
	} else if http.StatusText(statusCode) == "" {
		col = 111
	}

	l.CPrint("[", 243, 0)
	l.CPrint(fmt.Sprintf("%-3d", statusCode), col, 0)
	l.CPrint("] ", 243, 0)

	l.CPrint(path, 248, 0)
	l.CPrint("\n", 0, 0)
}

func _fmtDuration(d time.Duration) string {
	if d < time.Microsecond {
		return fmt.Sprintf("%dns", d.Nanoseconds())
	} else if d < time.Millisecond {
		return fmt.Sprintf("%dµs", d.Microseconds())
	} else if d < time.Second {
		return fmt.Sprintf("%dms", d.Milliseconds())
	} else if d < time.Minute {
		return fmt.Sprintf("%ds", d.Milliseconds()/1000)
	}

	return fmt.Sprintf("%dm", d.Milliseconds()/1000/60)
}
