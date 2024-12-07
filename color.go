package logger

import (
	"regexp"
)

func (l *Logger) parseColorCodes(builder *colorBuilder, fg, text string) {
	rgx := regexp.MustCompile(`~(\d+|r)~`)
	matches := rgx.FindAllStringSubmatchIndex(text, -1)

	var (
		index int
		chunk string

		code = fg
	)

	for _, match := range matches {
		chunk = text[index:match[0]]

		if code == "r" {
			builder.Write("", chunk)
		} else {
			builder.Write(code, chunk)
		}

		code = text[match[2]:match[3]]
		index = match[1]

		builder.ForceColor()
	}

	if index < len(text) {
		chunk = text[index:]

		if code == "r" {
			builder.Write("", chunk)
		} else {
			builder.Write(code, chunk)
		}
	}
}
