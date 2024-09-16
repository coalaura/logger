package logger

import "net/http"

// DiscardWriter is an io.Writer that discards all data written to it.
type DiscardWriter struct{}

func (d DiscardWriter) Write(b []byte) (int, error) {
	return len(b), nil
}

func (d DiscardWriter) WriteHeader(_ int) {}

func (d DiscardWriter) Header() http.Header {
	return http.Header{}
}
