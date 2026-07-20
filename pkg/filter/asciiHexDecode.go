package filter

import (
	"io"
)

type asciiHexDecode struct {
	baseFilter
}

const eodHexDecode = '>'

func (f asciiHexDecode) Encode(r io.Reader) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (f asciiHexDecode) Decode(r io.Reader) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (f asciiHexDecode) DecodeLength(r io.Reader, maxLen int64) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}
