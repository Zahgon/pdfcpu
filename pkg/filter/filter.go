package filter

import (
	"bytes"
	"errors"
	"io"
)

const (
	ASCII85   = "ASCII85Decode"
	ASCIIHex  = "ASCIIHexDecode"
	RunLength = "RunLengthDecode"
	LZW       = "LZWDecode"
	Flate     = "FlateDecode"
	CCITTFax  = "CCITTFaxDecode"
	JBIG2     = "JBIG2Decode"
	DCT       = "DCTDecode"
	JPX       = "JPXDecode"
)

var ErrUnsupportedFilter = errors.New("filter not supported")

var ErrDecodeLimitExceeded = errors.New("filter decode limit exceeded")

const DefaultMaxDecodeBytes int64 = 512 << 20

const maxInt = int(^uint(0) >> 1)
const maxInt64 = int64(^uint64(0) >> 1)

type Filter interface {
	Encode(r io.Reader) (io.Reader, error)
	Decode(r io.Reader) (io.Reader, error)

	DecodeLength(r io.Reader, maxLen int64) (io.Reader, error)
}

func NewFilter(filterName string, parms map[string]int, maxDecodeBytes ...int64) (filter Filter, err error) {
	_ = "STUB: not implemented"
	return *new(Filter), nil
}

func List() []string { _ = "STUB: not implemented"; return nil }

type baseFilter struct {
	parms          map[string]int
	maxDecodeBytes int64
}

func SupportsDecodeParms(f string) bool { _ = "STUB: not implemented"; return false }

func getReaderBytes(r io.Reader) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (f baseFilter) decodeLimit(maxLen int64) int64 { _ = "STUB: not implemented"; return 0 }

func (f baseFilter) copyDecoded(r io.Reader, maxLen int64) (*bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
