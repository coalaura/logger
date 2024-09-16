package http

import "net/http"

type statusWriter struct {
	status int
	w      http.ResponseWriter
}

func (w *statusWriter) Header() http.Header {
	return w.w.Header()
}

func (w *statusWriter) Write(b []byte) (int, error) {
	return w.w.Write(b)
}

func (w *statusWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.w.WriteHeader(statusCode)
}
