package plain

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

type PlainLogger struct {
	target  *os.File
	defCode []byte
	errCode []byte
	colored bool
}

// New creates a new plain logger with the given options.
func New(opts ...option) *PlainLogger {
	p := PlainLogger{
		target:  os.Stdout,
		defCode: []byte("\x1b[38;5;188m"),
		errCode: []byte("\x1b[38;5;196m"),
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
func (p *PlainLogger) Printf(format string, a ...any) (int, error) {
	if _, err := p.target.Write(p.defCode); err != nil {
		return 0, err
	}

	return fmt.Fprintf(p, format, a...)
}

// Println formats using the default formats for its operands and writes to the target output.
func (p *PlainLogger) Println(a ...any) (int, error) {
	if _, err := p.target.Write(p.defCode); err != nil {
		return 0, err
	}

	return fmt.Fprintln(p, a...)
}

// Errorf formats according to a format specifier and writes to the target output as an error.
func (p *PlainLogger) Errorf(format string, a ...any) (int, error) {
	if _, err := p.target.Write(p.errCode); err != nil {
		return 0, err
	}

	return fmt.Fprintf(p, format, a...)
}

// Errorln formats using the default formats for its operands and writes to the target output as an error.
func (p *PlainLogger) Errorln(a ...any) (int, error) {
	if _, err := p.target.Write(p.errCode); err != nil {
		return 0, err
	}

	return fmt.Fprintln(p, a...)
}

// MustFail logs and exits with code 1 if the error is not nil.
func (p *PlainLogger) MustFail(err error) {
	if err == nil {
		return
	}

	p.Errorln(err)

	os.Exit(1)
}
