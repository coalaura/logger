package http

import (
	"net/http"
	"sync/atomic"
)

type StatusRecorder struct {
	http.ResponseWriter
	status int32
}

func NewStatusRecorder(w http.ResponseWriter) *StatusRecorder {
	return &StatusRecorder{
		ResponseWriter: w,
		status:         0,
	}
}

func (r *StatusRecorder) WriteHeader(code int) {
	atomic.CompareAndSwapInt32(&r.status, 0, int32(code))

	r.ResponseWriter.WriteHeader(code)
}

func (r *StatusRecorder) Write(b []byte) (int, error) {
	atomic.CompareAndSwapInt32(&r.status, 0, http.StatusOK)

	return r.ResponseWriter.Write(b)
}

func (r *StatusRecorder) StatusCode() int {
	return int(atomic.LoadInt32(&r.status))
}

func (r *StatusRecorder) Unwrap() http.ResponseWriter {
	return r.ResponseWriter
}
