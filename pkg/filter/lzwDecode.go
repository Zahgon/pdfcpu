package filter

import (
	"io"
)

type lzwDecode struct {
	baseFilter
}

func (f lzwDecode) Encode(r io.Reader) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (f lzwDecode) Decode(r io.Reader) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (f lzwDecode) DecodeLength(r io.Reader, maxLen int64) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}
