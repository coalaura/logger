package http

import (
	"net/http"
	"time"

	"github.com/coalaura/logger"
	"github.com/felixge/httpsnoop"
)

type HTTPAdapter struct {
	request   *http.Request
	metrics   httpsnoop.Metrics
	timeTaken time.Duration
}

func Middleware(log *logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			metrics := httpsnoop.CaptureMetrics(next, w, r)

			log.LogHTTPRequest(&HTTPAdapter{
				request:   r,
				metrics:   metrics,
				timeTaken: time.Since(start),
			})
		})
	}
}

func (a *HTTPAdapter) Method() string {
	return a.request.Method
}

func (a *HTTPAdapter) Path() string {
	return a.request.URL.Path
}

func (a *HTTPAdapter) ClientIP() string {
	return a.request.RemoteAddr
}

func (a *HTTPAdapter) StatusCode() int {
	return a.metrics.Code
}

func (a *HTTPAdapter) TimeTaken() time.Duration {
	return a.timeTaken
}
