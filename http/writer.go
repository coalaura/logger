package http

import (
	"net/http"
	"sync/atomic"

	"github.com/felixge/httpsnoop"
)

type StatusRecorder struct {
	http.ResponseWriter
	status int32
}

func NewStatusRecorder(w http.ResponseWriter) (http.ResponseWriter, *StatusRecorder) {
	recorder := &StatusRecorder{
		ResponseWriter: w,
		status:         0,
	}

	wrapped := httpsnoop.Wrap(w, httpsnoop.Hooks{
		WriteHeader: func(next httpsnoop.WriteHeaderFunc) httpsnoop.WriteHeaderFunc {
			return func(code int) {
				atomic.CompareAndSwapInt32(&recorder.status, 0, int32(code))

				next(code)
			}
		},
		Write: func(next httpsnoop.WriteFunc) httpsnoop.WriteFunc {
			return func(b []byte) (int, error) {
				atomic.CompareAndSwapInt32(&recorder.status, 0, http.StatusOK)

				return next(b)
			}
		},
	})

	return wrapped, recorder
}

func (r *StatusRecorder) StatusCode() int {
	return int(atomic.LoadInt32(&r.status))
}
