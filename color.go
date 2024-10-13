package logger

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gookit/color"
)

func (l *Logger) _printColor(msg string, foreground, background int) {
	if !l.options.NoColor && foreground > 0 {
		msg = color.RenderCode("38;5;"+strconv.Itoa(foreground), msg)
	}

	if !l.options.NoColor && background > 0 {
		msg = color.RenderCode("48;5;"+strconv.Itoa(background), msg)
	}

	color.Fprint(l.out, msg)

	_, _ = color.Reset()
}

func (l *Logger) _printWithCodes(text string) {
	rgx := regexp.MustCompile(`~(\d+)~`)
	matches := rgx.FindAllStringSubmatchIndex(text, -1)

	var (
		index = 0
		code  = "248"

		chunk  string
		result strings.Builder
	)

	for _, match := range matches {
		chunk = text[index:match[0]]

		if code != "" {
			chunk = color.RenderCode("38;5;"+code, chunk)
		}

		result.WriteString(chunk)

		code = text[match[2]:match[3]]
		index = match[1]
	}

	if index < len(text) {
		chunk = text[index:]

		if code != "" {
			chunk = color.RenderCode("38;5;"+code, chunk)
		}

		result.WriteString(chunk)
	}

	color.Fprint(l.out, result.String())
	_, _ = color.Reset()
}
