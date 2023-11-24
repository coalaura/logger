package logger

import (
	"fmt"
	"time"
)

const (
	LevelDebug = iota
	LevelNotice
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

func (l *Logger) _level(level int) (int, string) {
	switch level {
	case LevelDebug:
		return 188, "debug"
	case LevelNotice:
		return 216, "note"
	case LevelInfo:
		return 117, "info"
	case LevelWarning:
		return 202, "warn"
	case LevelError:
		return 124, "error"
	case LevelFatal:
		return 196, "fatal"
	}

	return 0, ""
}

func (l *Logger) date() string {
	return fmt.Sprintf("[%-23s] ", time.Now().Format("2006-01-02T15:04:05 MST"))
}

func (l *Logger) print(level int, data ...interface{}) {
	l.mx.Lock()
	defer l.mx.Unlock()

	if !l.options.NoTime {
		l.CPrint(l.date(), 243, 0)
	}

	if !l.options.NoLevel {
		fg, lvl := l._level(level)

		if lvl != "" {
			l.CPrint("[", 243, 0)
			l.CPrint(fmt.Sprintf("%-7s", lvl), fg, 0)
			l.CPrint("] ", 243, 0)
		}
	}

	l.CPrint(fmt.Sprint(data...), 248, 0)
}

func (l *Logger) printF(level int, format string, data ...interface{}) {
	l.print(level, fmt.Sprintf(format, data...))
}

func (l *Logger) printLn(level int, data ...interface{}) {
	if len(data) == 0 {
		l.mx.Lock()
		defer l.mx.Unlock()

		l.CPrint("\n", 0, 0)

		return
	}

	data = append(data, "\n")

	l.print(level, data...)
}

func (l *Logger) printFln(level int, format string, data ...interface{}) {
	l.printF(level, format+"\n", data...)
}
