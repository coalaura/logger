package logger

import (
	"io"
	"os"
	"strconv"
	"sync"
)

type Logger struct {
	options Options
	out     io.Writer

	foreground string
	background string

	mx sync.Mutex
}

// Options are the options for the logger
// NoColor:      Disable colors
// NoLevel:      Disable log levels
// NoTime:       Disable timestamps
// ParseCodes:   Parse color codes (e.g. ~39~blue text ~160~red text ~r~reset text ~15~white text)
type Options struct {
	NoColor    bool
	NoLevel    bool
	NoTime     bool
	ParseCodes bool
}

// New creates a new logger instance with the default options
func New() *Logger {
	return (&Logger{}).WithOutput(os.Stdout).WithNoBackground().WithForeground(248)
}

// WithOptions sets the options of the logger
func (l *Logger) WithOptions(options Options) *Logger {
	l.options = options

	// Disable foreground and background if NoColor is enabled
	if l.options.NoColor {
		l.foreground = ""
		l.background = ""
	}

	return l
}

// WithForeground sets the foreground color of the logger
func (l *Logger) WithForeground(code uint8) *Logger {
	l.foreground = strconv.Itoa(int(code))

	return l
}

// WithNoForeground disables the foreground color of the logger
func (l *Logger) WithNoForeground() *Logger {
	l.foreground = ""

	return l
}

// WithBackground sets the background color of the logger
func (l *Logger) WithBackground(code uint8) *Logger {
	l.background = strconv.Itoa(int(code))

	return l
}

// WithNoBackground disables the background color of the logger
func (l *Logger) WithNoBackground() *Logger {
	l.background = ""

	return l
}

// WithOutput sets the output of the logger
func (l *Logger) WithOutput(out io.Writer) *Logger {
	l.out = out

	return l
}
