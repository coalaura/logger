package logger

import (
	"fmt"
	"time"

	"github.com/gookit/color"
)

const (
	LevelDebug = iota
	LevelNotice
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

// level returns the color and level name
func (l *Logger) level(level int) (string, string) {
	switch level {
	case LevelDebug:
		return "188", "debug"
	case LevelNotice:
		return "216", "note"
	case LevelInfo:
		return "117", "info"
	case LevelWarning:
		return "202", "warn"
	case LevelError:
		return "124", "error"
	case LevelFatal:
		return "196", "fatal"
	}

	return "", ""
}

// date returns the date string
func (l *Logger) date() string {
	return fmt.Sprintf("[%-23s] ", time.Now().Format("2006-01-02T15:04:05 MST"))
}

// write writes the message to the output
func (l *Logger) write(msg string) {
	l.mx.Lock()
	defer l.mx.Unlock()

	if l.forceNoColor {
		l.out.Write([]byte(msg))

		return
	}

	color.Fprint(l.out, msg)
	color.Reset()
}

// print colors the message and then writes it
func (l *Logger) print(level int, data ...interface{}) {
	builder := newColorBuilder(l)

	if !l.options.NoTime {
		builder.Write("243", l.date())
	}

	fg, lvl := l.level(level)

	if !l.options.NoLevel {
		if lvl != "" {
			builder.Write("243", "[")
			builder.WriteF(fg, "%-7s", lvl)
			builder.Write("243", "] ")
		}

		fg = l.foreground
	}

	if l.options.ParseCodes {
		l.parseColorCodes(builder, fg, fmt.Sprint(data...))
	} else {
		builder.Write(fg, fmt.Sprint(data...))
	}

	l.write(builder.String())
}

// printf formats the data and then prints it
func (l *Logger) printf(level int, format string, data ...interface{}) {
	l.print(level, fmt.Sprintf(format, data...))
}

// printLn formats the data with a new line and then prints it
func (l *Logger) println(level int, data ...interface{}) {
	data = append(data, "\n")

	l.print(level, data...)
}
