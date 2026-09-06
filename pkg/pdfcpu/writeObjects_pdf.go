package pdfcpu

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func appendPDFObject(dst []byte, obj types.Object) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func appendLazyObjectStreamObject(dst []byte, obj types.LazyObjectStreamObject) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func appendPDFArray(dst []byte, a types.Array) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func arrayObjectNeedsSpace(obj types.Object) bool { _ = "STUB: not implemented"; return false }

func appendPDFDict(dst []byte, d types.Dict) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sortedDictKeys(d types.Dict) []string { _ = "STUB: not implemented"; return nil }

func appendPDFDictEntry(dst []byte, key string, obj types.Object) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dictObjectNeedsSpace(obj types.Object) bool { _ = "STUB: not implemented"; return false }
