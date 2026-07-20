package filter

import (
	"io"
)

type dctDecode struct {
	baseFilter
}

func (f dctDecode) Encode(r io.Reader) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (f dctDecode) Decode(r io.Reader) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (f dctDecode) DecodeLength(r io.Reader, maxLen int64) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}
