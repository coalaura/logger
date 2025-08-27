package plain

import (
	"fmt"
	"os"
)

type option func(*PlainLogger)

// WithColor sets the foreground color code.
func WithColor(code int) option {
	return func(p *PlainLogger) {
		p.defCode = []byte(fmt.Sprintf("\x1b[38;5;%dm", code))
	}
}

// WithError sets the foreground color code for errors.
func WithError(code int) option {
	return func(p *PlainLogger) {
		p.errCode = []byte(fmt.Sprintf("\x1b[38;5;%dm", code))
	}
}

// WithTarget sets the target writer.
func WithTarget(target *os.File) option {
	return func(p *PlainLogger) {
		p.target = target
	}
}
