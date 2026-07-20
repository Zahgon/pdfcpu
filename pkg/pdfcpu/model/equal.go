package model

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func appendPair(pairs []int, a, b int) []int { _ = "STUB: not implemented"; return nil }

func containsPair(pairs []int, a, b int) bool { _ = "STUB: not implemented"; return false }

func EqualObjects(o1, o2 types.Object, xRefTable *XRefTable, pairs []int) (ok bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func equalArrays(a1, a2 types.Array, xRefTable *XRefTable, pairs []int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func equalStreamDicts(sd1, sd2 *types.StreamDict, xRefTable *XRefTable, pairs []int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func equalFontNames(v1, v2 types.Object, xRefTable *XRefTable) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func equalDicts(d1, d2 types.Dict, xRefTable *XRefTable, pairs []int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
