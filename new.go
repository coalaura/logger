package logger

import (
	"io"
	"os"
	"sync"
)

type Logger struct {
	options Options
	out     io.Writer

	mx sync.Mutex
}

type Options struct {
	NoColor bool
	NoLevel bool
	NoTime  bool
}

// New creates a new logger instance with the given options
func New() *Logger {
	return &Logger{
		out: os.Stdout,
	}
}

// WithOptions sets the options of the logger
func (l *Logger) WithOptions(options Options) *Logger {
	l.options = options

	return l
}

// WithOutput sets the output of the logger
func (l *Logger) WithOutput(out io.Writer) *Logger {
	l.out = out

	return l
}
