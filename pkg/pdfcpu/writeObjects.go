package pdfcpu

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

const (
	ObjectStreamMaxObjects = 100
)

func writeCommentLine(w *model.WriteContext, comment string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func writeHeader(w *model.WriteContext, v model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func writeTrailer(w *model.WriteContext) error { _ = "STUB: not implemented"; return nil }

func objectHeader(objNr, genNr int, eol string) string { _ = "STUB: not implemented"; return "" }

func writeObjectHeader(w *model.WriteContext, objNr, genNr int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func writeObjectTrailer(w *model.WriteContext) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func startObjectStream(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func stopObjectStream(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func addObjectStreamObject(osd types.ObjectStreamDict, objNumber int, obj types.Object) (types.ObjectStreamDict, error) {
	_ = "STUB: not implemented"
	return *new(types.ObjectStreamDict), nil
}

func writeToObjectStream(ctx *model.Context, objNumber, genNumber int) (ok bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func writeObject(ctx *model.Context, objNumber, genNumber int, s string) error {
	_ = "STUB: not implemented"
	return nil
}

func writePDFNullObject(ctx *model.Context, objNumber, genNumber int) error {
	_ = "STUB: not implemented"
	return nil
}

func writeBooleanObject(ctx *model.Context, objNumber, genNumber int, boolean types.Boolean) error {
	_ = "STUB: not implemented"
	return nil
}

func writeNameObject(ctx *model.Context, objNumber, genNumber int, name types.Name) error {
	_ = "STUB: not implemented"
	return nil
}

func writeStringLiteralObject(ctx *model.Context, objNumber, genNumber int, sl types.StringLiteral) error {
	_ = "STUB: not implemented"
	return nil
}

func writeHexLiteralObject(ctx *model.Context, objNumber, genNumber int, hl types.HexLiteral) error {
	_ = "STUB: not implemented"
	return nil
}

func writeIntegerObject(ctx *model.Context, objNumber, genNumber int, integer types.Integer) error {
	_ = "STUB: not implemented"
	return nil
}

func writeFloatObject(ctx *model.Context, objNumber, genNumber int, float types.Float) error {
	_ = "STUB: not implemented"
	return nil
}

func sigDict(d types.Dict) bool { _ = "STUB: not implemented"; return false }

func sigDictPDFString(ctx *model.Context, d types.Dict, objNr, genNr int) string {
	_ = "STUB: not implemented"
	return ""
}

func writeDictObject(ctx *model.Context, objNr, genNr int, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func writeArrayObject(ctx *model.Context, objNumber, genNumber int, a types.Array) error {
	_ = "STUB: not implemented"
	return nil
}

func writeStream(w *model.WriteContext, sd types.StreamDict) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func handleIndirectLength(ctx *model.Context, ir *types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func writeStreamObject(ctx *model.Context, objNr, genNr int, sd types.StreamDict, pdfString string) (int, int64, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, nil
}

func writeStreamDictObject(ctx *model.Context, objNr, genNr int, sd types.StreamDict) error {
	_ = "STUB: not implemented"
	return nil
}

func writeDirectObject(ctx *model.Context, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func writeNullObject(ctx *model.Context, objNumber, genNumber int) error {
	_ = "STUB: not implemented"
	return nil
}

func writeDeepDict(ctx *model.Context, d types.Dict, objNr, genNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func writeDeepStreamDict(ctx *model.Context, sd *types.StreamDict, objNr, genNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func writeDeepArray(ctx *model.Context, a types.Array, objNr, genNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func writeLazyObjectStreamObject(ctx *model.Context, objNr, genNr int, o types.LazyObjectStreamObject) error {
	_ = "STUB: not implemented"
	return nil
}

func writeObjectGeneric(ctx *model.Context, o types.Object, objNr, genNr int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func writeIndirectObject(ctx *model.Context, ir types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func writeDeepObject(ctx *model.Context, objIn types.Object) (objOut types.Object, written bool, err error) {
	_ = "STUB: not implemented"
	return *new(types.Object), false, nil
}

func writeEntry(ctx *model.Context, d types.Dict, dictName, entryName string) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func writeFlatObject(ctx *model.Context, objNr int) error { _ = "STUB: not implemented"; return nil }
