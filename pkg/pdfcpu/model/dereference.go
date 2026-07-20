package model

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func processDictRefCounts(xRefTable *XRefTable, d types.Dict, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func processArrayRefCounts(xRefTable *XRefTable, a types.Array, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func processRefCounts(xRefTable *XRefTable, o types.Object, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func ProcessRefCountsWithError(xRefTable *XRefTable, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func ProcessRefCounts(xRefTable *XRefTable, o types.Object) { _ = "STUB: not implemented"; return }

func (xRefTable *XRefTable) indRefToObject(ir *types.IndirectRef, decodeLazy bool) (types.Object, int, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), 0, nil
}

func (xRefTable *XRefTable) Dereference(o types.Object) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func (xRefTable *XRefTable) DereferenceWithIncr(o types.Object) (types.Object, int, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), 0, nil
}

func (xRefTable *XRefTable) DereferenceForWrite(o types.Object) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func (xRefTable *XRefTable) DereferenceBoolean(o types.Object, sinceVersion Version) (*types.Boolean, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) DereferenceInteger(o types.Object) (*types.Integer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) DereferenceNumber(o types.Object) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (xRefTable *XRefTable) DereferenceName(o types.Object, sinceVersion Version, validate func(string) bool) (n types.Name, err error) {
	_ = "STUB: not implemented"
	return *new(types.Name), nil
}

func (xRefTable *XRefTable) DereferenceStringLiteral(o types.Object, sinceVersion Version, validate func(string) bool) (s types.StringLiteral, err error) {
	_ = "STUB: not implemented"
	return *new(types.StringLiteral), nil
}

func (xRefTable *XRefTable) DereferenceStringOrHexLiteral(obj types.Object, sinceVersion Version, validate func(string) bool) (s string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func Text(o types.Object) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (xRefTable *XRefTable) DereferenceText(o types.Object) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func CSVSafeString(s string) string { _ = "STUB: not implemented"; return "" }

func (xRefTable *XRefTable) DereferenceCSVSafeText(o types.Object) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (xRefTable *XRefTable) DereferenceArray(o types.Object) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func (xRefTable *XRefTable) DereferenceDict(o types.Object) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func (xRefTable *XRefTable) DereferenceDictWithIncr(o types.Object) (types.Dict, int, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), 0, nil
}

func (xRefTable *XRefTable) DereferenceFontDict(indRef types.IndirectRef) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func (xRefTable *XRefTable) DereferencePageNodeDict(indRef types.IndirectRef) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func (xRefTable *XRefTable) dereferenceDestArray(o types.Object) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func (xRefTable *XRefTable) DereferenceDestArray(key string) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func (xRefTable *XRefTable) DereferenceDictEntry(d types.Dict, key string) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func (xRefTable *XRefTable) DereferenceStringEntryBytes(d types.Dict, key string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) DestName(obj types.Object) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
