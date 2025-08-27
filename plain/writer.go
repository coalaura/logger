package plain

type writer struct {
	code  []byte
	plain *PlainLogger
}

func wrap(pl *PlainLogger, code []byte) *writer {
	return &writer{
		code:  code,
		plain: pl,
	}
}

// Write implements [io.Writer].
func (w *writer) Write(b []byte) (int, error) {
	return w.plain.write(w.code, b)
}
