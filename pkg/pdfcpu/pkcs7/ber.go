package pkcs7

import (
	"bytes"
)

type asn1Object interface {
	EncodeTo(writer *bytes.Buffer) error
}

type asn1Structured struct {
	tagBytes []byte
	content  []asn1Object
}

func (s asn1Structured) EncodeTo(out *bytes.Buffer) error { _ = "STUB: not implemented"; return nil }

type asn1Primitive struct {
	tagBytes []byte
	content  []byte
}

func (p asn1Primitive) EncodeTo(out *bytes.Buffer) error { _ = "STUB: not implemented"; return nil }

func ber2der(ber []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func allZero(bb []byte) bool { _ = "STUB: not implemented"; return false }

func marshalLongLength(out *bytes.Buffer, i int) (err error) { _ = "STUB: not implemented"; return nil }

func lengthLength(i int) (numBytes int) { _ = "STUB: not implemented"; return 0 }

func encodeLength(out *bytes.Buffer, length int) (err error) { _ = "STUB: not implemented"; return nil }

func readObject(ber []byte, offset int) (asn1Object, int, error) {
	_ = "STUB: not implemented"
	return *new(asn1Object), 0, nil
}

func readTag(ber []byte, offset int) (start, end int, constructed bool, next int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, false, 0, nil
}

func readLength(ber []byte, offset int) (length int, indefinite bool, next int, err error) {
	_ = "STUB: not implemented"
	return 0, false, 0, nil
}

func readStructuredContent(
	ber []byte,
	offset, contentEnd int,
	indefinite bool,
) ([]asn1Object, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func isIndefiniteTermination(ber []byte, offset int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
