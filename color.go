package logger_v2

import (
	"strconv"

	"github.com/gookit/color"
)

// CPrint prints a message with the given colors, to the logger's output.
func (l *Logger) CPrint(msg string, foreground, background int) {
	if !l.options.NoColor && foreground > 0 {
		msg = color.RenderCode("38;5;"+strconv.Itoa(foreground), msg)
	}

	if !l.options.NoColor && background > 0 {
		msg = color.RenderCode("48;5;"+strconv.Itoa(background), msg)
	}

	color.Fprint(l.out, msg)

	_, _ = color.Reset()
}
