package logger_v2

// PrintF formats the data and prints it without any logging information
func (l *Logger) PrintF(msg string, data ...interface{}) {
	l.printF(-1, msg, data...)
}

// Print prints the data and without any logging information
func (l *Logger) Print(data ...interface{}) {
	l.print(-1, data...)
}

// PrintLn prints the data and without any logging information
func (l *Logger) PrintLn(data ...interface{}) {
	l.printLn(-1, data...)
}

// DebugF formats the data and prints it with the log level debug
func (l *Logger) DebugF(msg string, data ...interface{}) {
	l.printF(LevelDebug, msg, data...)
}

// Debug prints the data with the log level debug
func (l *Logger) Debug(data ...interface{}) {
	l.printLn(LevelDebug, data...)
}

// NoteF formats the data and prints it with the log level notice
func (l *Logger) NoteF(msg string, data ...interface{}) {
	l.printF(LevelNotice, msg, data...)
}

// Note prints the data with the log level notice
func (l *Logger) Note(data ...interface{}) {
	l.printLn(LevelNotice, data...)
}

// InfoF formats the data and prints it with the log level info
func (l *Logger) InfoF(msg string, data ...interface{}) {
	l.printF(LevelInfo, msg, data...)
}

// Info prints the data with the log level info
func (l *Logger) Info(data ...interface{}) {
	l.printLn(LevelInfo, data...)
}

// WarningF formats the data and prints it with the log level warning
func (l *Logger) WarningF(msg string, data ...interface{}) {
	l.printF(LevelWarning, msg, data...)
}

// Warning prints the data with the log level warning
func (l *Logger) Warning(data ...interface{}) {
	l.printLn(LevelWarning, data...)
}

// ErrorF formats the data and prints it with the log level error
func (l *Logger) ErrorF(msg string, data ...interface{}) {
	l.printF(LevelError, msg, data...)
}

// Error prints the data with the log level error
func (l *Logger) Error(data ...interface{}) {
	l.printLn(LevelError, data...)
}

// FatalF formats the data and prints it with the log level fatal
func (l *Logger) FatalF(msg string, data ...interface{}) {
	l.printF(LevelFatal, msg, data...)
}

// Fatal prints the data with the log level fatal
func (l *Logger) Fatal(data ...interface{}) {
	l.printLn(LevelFatal, data...)
}
