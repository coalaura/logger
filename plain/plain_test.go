package plain

import "testing"

func TestPlainLogger(t *testing.T) {
	pl := New()

	pl.Printf("This is a %s message with %d numbers and %.2f floats", "plain", 123, 12.3456)
}
