package logger

// DebugE logs an error with the log level debug
func (l *Logger) DebugE(err error) {
	l.println(LevelDebug, err.Error())
}

// NoteE logs an error with the log level note
func (l *Logger) NoteE(err error) {
	l.println(LevelNotice, err.Error())
}

// InfoE logs an error with the log level info
func (l *Logger) InfoE(err error) {
	l.println(LevelInfo, err.Error())
}

// WarningE logs an error with the log level warning
func (l *Logger) WarningE(err error) {
	l.println(LevelWarning, err.Error())
}

// ErrorE logs an error with the log level error
func (l *Logger) ErrorE(err error) {
	l.println(LevelError, err.Error())
}

// FatalE logs an error with the log level fatal
func (l *Logger) FatalE(err error) {
	l.println(LevelFatal, err.Error())
}

// Panic logs and panics with the given error
func (l *Logger) Panic(err error) {
	l.FatalE(err)

	panic(err)
}

// MustPanic logs and panics if the error is not nil
func (l *Logger) MustPanic(err error) {
	if err == nil {
		return
	}

	l.Panic(err)
}
