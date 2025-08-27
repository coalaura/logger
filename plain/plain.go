package plain

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

type PlainLogger struct {
	target  *os.File
	defCode []byte
	wrnCode []byte
	errCode []byte
	colored bool
}

var rstCode = []byte("\x1b[0m")

// New creates a new plain logger with the given options.
func New(opts ...option) *PlainLogger {
	p := PlainLogger{
		target:  os.Stdout,
		defCode: []byte("\x1b[37m"),
		wrnCode: []byte("\x1b[33m"),
		errCode: []byte("\x1b[31m"),
	}

	for _, opt := range opts {
		opt(&p)
	}

	p.colored = term.IsTerminal(int(p.target.Fd()))

	return &p
}

func (p *PlainLogger) write(code []byte, b []byte) (int, error) {
	if !p.colored {
		return p.target.Write(b)
	}

	var (
		n   int
		buf = make([]byte, len(code)+len(b)+len(rstCode))
	)

	n += copy(buf[n:], code)
	n += copy(buf[n:], b)
	n += copy(buf[n:], rstCode)

	return p.target.Write(buf)
}

// Write implements [io.Writer].
func (p *PlainLogger) Write(b []byte) (int, error) {
	if !p.colored {
		return p.target.Write(b)
	}

	return p.write(p.defCode, b)
}

// Printf formats according to a format specifier and writes to the target output.
func (p *PlainLogger) Printf(format string, a ...any) (int, error) {
	return fmt.Fprintf(p, format, a...)
}

// Println formats using the default formats for its operands and writes to the target output.
func (p *PlainLogger) Println(a ...any) (int, error) {
	return fmt.Fprintln(p, a...)
}

// Warnf formats according to a format specifier and writes to the target output as a warning.
func (p *PlainLogger) Warnf(format string, a ...any) (int, error) {
	wr := wrap(p, p.wrnCode)

	return fmt.Fprintf(wr, format, a...)
}

// Warnln formats using the default formats for its operands and writes to the target output as a warning.
func (p *PlainLogger) Warnln(a ...any) (int, error) {
	wr := wrap(p, p.wrnCode)

	return fmt.Fprintln(wr, a...)
}

// Errorf formats according to a format specifier and writes to the target output as an error.
func (p *PlainLogger) Errorf(format string, a ...any) (int, error) {
	wr := wrap(p, p.errCode)

	return fmt.Fprintf(wr, format, a...)
}

// Errorln formats using the default formats for its operands and writes to the target output as an error.
func (p *PlainLogger) Errorln(a ...any) (int, error) {
	wr := wrap(p, p.errCode)

	return fmt.Fprintln(wr, a...)
}

// MustFail logs and exits with code 1 if the error is not nil.
func (p *PlainLogger) MustFail(err error) {
	if err == nil {
		return
	}

	p.Errorln(err)

	os.Exit(1)
}
