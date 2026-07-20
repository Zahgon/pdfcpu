package filter

import (
	"io"
)

type ascii85Decode struct {
	baseFilter
}

const eodASCII85 = "~>"

func (f ascii85Decode) Encode(r io.Reader) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (f ascii85Decode) Decode(r io.Reader) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (f ascii85Decode) DecodeLength(r io.Reader, maxLen int64) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}
