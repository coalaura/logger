package logger

import (
	"fmt"
	"strings"
	"time"
)

type MiddlewareAdapter interface {
	Method() string
	Path() string
	ClientIP() string
	StatusCode() int
	TimeTaken() time.Duration
}

var (
	MethodColors = map[string]string{
		"GET":     "33",
		"HEAD":    "141",
		"POST":    "36",
		"PUT":     "34",
		"DELETE":  "160",
		"CONNECT": "45",
		"OPTIONS": "209",
		"TRACE":   "162",
		"PATCH":   "202",
	}

	// Integer divided by 100 (floored)
	StatusCodeColors = map[int]string{
		1: "159",
		2: "34",
		3: "184",
		4: "202",
		5: "124",
	}

	Methods = []string{"GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"}
)

// LogHTTPRequest logs the given request with the given adapter.
func (l *Logger) LogHTTPRequest(adp MiddlewareAdapter) {
	// Get the request information
	method := adp.Method()
	path := adp.Path()
	ip := adp.ClientIP()
	statusCode := adp.StatusCode()
	timeTaken := adp.TimeTaken()

	// Resolve colors
	methodColor, ok := MethodColors[method]
	if !ok {
		methodColor = "52"
	}

	statusColor, ok := StatusCodeColors[statusCode/100]
	if !ok {
		/*
		* 218 This is fine
		* 418 I'm a teapot
		* 420 Enhance Your Calm
		 */
		if statusCode == 218 || statusCode == 418 || statusCode == 420 {
			statusColor = "210"
		} else {
			statusColor = "111"
		}
	}

	// Format strings
	methodStr := fmt.Sprintf("%-7s", strings.ToLower(method))
	timeStr := fmt.Sprintf("%-5s ", _fmtDuration(timeTaken))
	ipStr := fmt.Sprintf("%-15s ", ip)
	statusStr := fmt.Sprintf("%-3d ", statusCode)

	// Log the request
	builder := newColorBuilder(l)

	builder.Write("243", l.date())

	builder.Write("243", "[")
	builder.Write(methodColor, methodStr)
	builder.Write("243", "] ")

	builder.Write("243", "[")
	builder.Write("115", timeStr)
	builder.Write("243", "] ")

	builder.Write("243", "[")
	builder.Write("248", ipStr)
	builder.Write("243", "] ")

	builder.Write("243", "[")
	builder.Write(statusColor, statusStr)
	builder.Write("243", "] ")

	builder.Write("248", path)

	l.write(builder.StringLn())
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
