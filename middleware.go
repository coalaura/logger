package logger

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type MiddlewareAdapter interface {
	Next()

	Request() http.Request
	ClientIP() string
	StatusCode() int
	TimeTaken() time.Duration
}

var (
	MethodColors = map[string]int{
		"GET":     33,
		"HEAD":    141,
		"POST":    36,
		"PUT":     34,
		"DELETE":  160,
		"CONNECT": 45,
		"OPTIONS": 209,
		"TRACE":   162,
		"PATCH":   202,
	}

	// Integer divided by 100 (floored)
	StatusCodeColors = map[int]int{
		1: 159,
		2: 34,
		3: 184,
		4: 202,
		5: 124,
	}
)

// LogHTTPRequest logs the given request with the given adapter.
func (l *Logger) LogHTTPRequest(adp MiddlewareAdapter) {
	// Get the request information
	request := adp.Request()
	ip := adp.ClientIP()
	statusCode := adp.StatusCode()
	timeTaken := adp.TimeTaken()

	// Resolve colors
	method, ok := MethodColors[request.Method]
	if !ok {
		method = 52
	}

	status, ok := StatusCodeColors[statusCode/100]
	if !ok {
		/*
		* 218 This is fine
		* 418 I'm a teapot
		* 420 Enhance Your Calm
		 */
		if statusCode == 218 || statusCode == 418 || statusCode == 420 {
			status = 210
		} else {
			status = 111
		}
	}

	// Format strings
	methodStr := fmt.Sprintf("%-7s", strings.ToLower(request.Method))
	timeStr := fmt.Sprintf("%-5s ", _fmtDuration(timeTaken))
	ipStr := fmt.Sprintf("%-15s ", ip)
	statusStr := fmt.Sprintf("%-3d ", statusCode)

	// Log the request
	l.mx.Lock()
	defer l.mx.Unlock()

	l.CPrint(l.date(), 243, 0)

	l.CPrint("[", 243, 0)
	l.CPrint(methodStr, method, 0)
	l.CPrint("] ", 243, 0)

	l.CPrint("[", 243, 0)
	l.CPrint(timeStr, 115, 0)
	l.CPrint("] ", 243, 0)

	l.CPrint("[", 243, 0)
	l.CPrint(ipStr, 248, 0)
	l.CPrint("] ", 243, 0)

	l.CPrint("[", 243, 0)
	l.CPrint(statusStr, status, 0)
	l.CPrint("] ", 243, 0)

	l.CPrint(request.URL.Path, 248, 0)
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
