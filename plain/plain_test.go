package plain

import (
	"fmt"
	"testing"
)

func TestPlainLogger(t *testing.T) {
	pl := New()

	pl.Printf("This is a %s message with %d numbers and %.2f floats\n", "plain", 123, 12.3456)
	pl.Warnf("This is a %s message with %d numbers and %.2f floats\n", "warning", 123, 12.3456)
	pl.Errorf("This is a %s message with %d numbers and %.2f floats\n", "error", 123, 12.3456)

	fmt.Fprintln(pl, "This is a message from fmt")
}
