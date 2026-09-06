package filter

import (
	"io"
)

type runLengthDecode struct {
	baseFilter
}

func (f runLengthDecode) decode(w io.ByteWriter, src []byte, maxLen int64) error {
	_ = "STUB: not implemented"
	return nil
}

func detect(i, start, maxLen int, b byte, src []byte) int { _ = "STUB: not implemented"; return 0 }

func (f runLengthDecode) encode(w io.ByteWriter, src []byte) { _ = "STUB: not implemented"; return }

func (f runLengthDecode) Encode(r io.Reader) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (f runLengthDecode) Decode(r io.Reader) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (f runLengthDecode) DecodeLength(r io.Reader, maxLen int64) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}
