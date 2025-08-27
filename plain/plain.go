package plain

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

type PlainLogger struct {
	target  *os.File
	code    []byte
	colored bool
}

// New creates a new plain logger with the given options.
func New(opts ...option) *PlainLogger {
	p := PlainLogger{
		target: os.Stdout,
		code:   []byte("\x1b[38;5;188m"),
	}

	for _, opt := range opts {
		opt(&p)
	}

	p.colored = term.IsTerminal(int(p.target.Fd()))

	return &p
}

// Write implements [io.Writer].
func (p *PlainLogger) Write(b []byte) (int, error) {
	if !p.colored {
		return p.target.Write(b)
	}

	if _, err := p.target.Write(p.code); err != nil {
		return 0, err
	}

	n, err := p.target.Write(b)
	if err != nil {
		return n, err
	}

	if _, err := p.target.Write([]byte("\x1b[0m")); err != nil {
		return n, err
	}

	return n, nil
}

// Printf formats according to a format specifier and writes to the target output.
func (p *PlainLogger) Printf(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(p, format, a...)
}

// Println formats using the default formats for its operands and writes to the target output.
func (p *PlainLogger) Println(a ...any) (n int, err error) {
	return fmt.Fprintln(p, a...)
}
