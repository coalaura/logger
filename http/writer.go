package http

import (
	"net/http"
	"sync/atomic"
)

type StatusWriter struct {
	http.ResponseWriter
	status int64
}

func (w *StatusWriter) WriteHeader(statusCode int) {
	atomic.SwapInt64(&w.status, int64(statusCode))

	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *StatusWriter) GetStatusCode() int {
	return int(atomic.LoadInt64(&w.status))
}

func NewStatusWriter(w http.ResponseWriter) http.ResponseWriter {
	var result http.ResponseWriter = &StatusWriter{ResponseWriter: w}

	// http.Flusher compliance
	if flusher, ok := w.(http.Flusher); ok {
		result = &struct {
			http.ResponseWriter
			http.Flusher
		}{
			ResponseWriter: result,
			Flusher:        flusher,
		}
	}

	// http.Hijacker compliance
	if hijacker, ok := w.(http.Hijacker); ok {
		result = &struct {
			http.ResponseWriter
			http.Hijacker
		}{
			ResponseWriter: result,
			Hijacker:       hijacker,
		}
	}

	// http.Pusher compliance
	if pusher, ok := w.(http.Pusher); ok {
		result = &struct {
			http.ResponseWriter
			http.Pusher
		}{
			ResponseWriter: result,
			Pusher:         pusher,
		}
	}

	return result
}
