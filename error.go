package logger

// DebugE logs an error with the log level debug
func (l *Logger) DebugE(err error) {
	l.printFln(LevelDebug, err.Error())
}

// NoteE logs an error with the log level note
func (l *Logger) NoteE(err error) {
	l.printFln(LevelNotice, err.Error())
}

// InfoE logs an error with the log level info
func (l *Logger) InfoE(err error) {
	l.printFln(LevelInfo, err.Error())
}

// WarningE logs an error with the log level warning
func (l *Logger) WarningE(err error) {
	l.printFln(LevelWarning, err.Error())
}

// ErrorE logs an error with the log level error
func (l *Logger) ErrorE(err error) {
	l.printFln(LevelError, err.Error())
}

// FatalE logs an error with the log level fatal
func (l *Logger) FatalE(err error) {
	l.printFln(LevelFatal, err.Error())
}

// MustPanic logs and panics if the error is not nil
func (l *Logger) MustPanic(err error) {
	if err != nil {
		l.FatalE(err)

		panic(err)
	}
}
