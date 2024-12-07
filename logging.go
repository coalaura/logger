package logger

// Printf formats the data and prints it without any logging information
func (l *Logger) Printf(msg string, data ...interface{}) {
	l.printf(-1, msg, data...)
}

// Print prints the data and without any logging information
func (l *Logger) Print(data ...interface{}) {
	l.print(-1, data...)
}

// Println prints the data and without any logging information
func (l *Logger) Println(data ...interface{}) {
	l.println(-1, data...)
}

// Debugf formats the data and prints it with the log level debug
func (l *Logger) Debugf(msg string, data ...interface{}) {
	l.printf(LevelDebug, msg, data...)
}

// Debug prints the data with the log level debug
func (l *Logger) Debug(data ...interface{}) {
	l.println(LevelDebug, data...)
}

// Notef formats the data and prints it with the log level notice
func (l *Logger) Notef(msg string, data ...interface{}) {
	l.printf(LevelNotice, msg, data...)
}

// Note prints the data with the log level notice
func (l *Logger) Note(data ...interface{}) {
	l.println(LevelNotice, data...)
}

// Infof formats the data and prints it with the log level info
func (l *Logger) Infof(msg string, data ...interface{}) {
	l.printf(LevelInfo, msg, data...)
}

// Info prints the data with the log level info
func (l *Logger) Info(data ...interface{}) {
	l.println(LevelInfo, data...)
}

// Warningf formats the data and prints it with the log level warning
func (l *Logger) Warningf(msg string, data ...interface{}) {
	l.printf(LevelWarning, msg, data...)
}

// Warning prints the data with the log level warning
func (l *Logger) Warning(data ...interface{}) {
	l.println(LevelWarning, data...)
}

// Errorf formats the data and prints it with the log level error
func (l *Logger) Errorf(msg string, data ...interface{}) {
	l.printf(LevelError, msg, data...)
}

// Error prints the data with the log level error
func (l *Logger) Error(data ...interface{}) {
	l.println(LevelError, data...)
}

// Fatalf formats the data and prints it with the log level fatal
func (l *Logger) Fatalf(msg string, data ...interface{}) {
	l.printf(LevelFatal, msg, data...)
}

// Fatal prints the data with the log level fatal
func (l *Logger) Fatal(data ...interface{}) {
	l.println(LevelFatal, data...)
}
