package http

import (
	"net/http"
	"time"

	"github.com/coalaura/logger"
)

type HTTPAdapter struct {
	request   *http.Request
	response  *StatusRecorder
	timeTaken time.Duration
}

func Middleware(log *logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wrapped, recorder := NewStatusRecorder(w)

			start := time.Now()

			next.ServeHTTP(wrapped, r)

			log.LogHTTPRequest(&HTTPAdapter{
				request:   r,
				response:  recorder,
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
	return a.response.StatusCode()
}

func (a *HTTPAdapter) TimeTaken() time.Duration {
	return a.timeTaken
}
