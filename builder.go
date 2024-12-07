package logger

import (
	"fmt"
	"strings"
)

type colorBuilder struct {
	strings.Builder

	color      bool
	background string

	hasForeground bool
	hasBackground bool
}

func newColorBuilder(logger *Logger) *colorBuilder {
	return &colorBuilder{
		color:      !logger.options.NoColor,
		background: logger.background,
	}
}

func (b *colorBuilder) ForceColor() {
	b.color = true
}

func (b *colorBuilder) String() string {
	if b.hasBackground || b.hasForeground {
		b.WriteString("\033[0m")
	}

	return b.Builder.String()
}

func (b *colorBuilder) StringLn() string {
	b.WriteRune('\n')

	return b.String()
}

func (b *colorBuilder) Write(foreground, str string) {
	if !b.color {
		b.WriteString(str)

		return
	}

	if b.background != "" {
		b.WriteString("\033[48;5;")
		b.WriteString(b.background)
		b.WriteString("m")

		b.hasBackground = true
	} else if b.hasBackground {
		b.WriteString("\033[49m")

		b.hasBackground = false
	}

	if foreground != "" {
		b.WriteString("\033[38;5;")
		b.WriteString(foreground)
		b.WriteString("m")

		b.hasForeground = true
	} else if b.hasForeground {
		b.WriteString("\033[39m")

		b.hasForeground = false
	}

	b.WriteString(str)
}

func (b *colorBuilder) WriteF(foreground, format string, data ...interface{}) {
	b.Write(foreground, fmt.Sprintf(format, data...))
}
