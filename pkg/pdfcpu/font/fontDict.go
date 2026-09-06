package font

import (
	"bufio"
	"bytes"
	"errors"

	"github.com/pdfcpu/pdfcpu/pkg/font"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type cjk struct {
	encoding   string
	ordering   string
	supplement int
}

var cjkParms = map[string]cjk{

	"HANS": {"UniGB-UTF16-H", "GB1", 5},
	"HANT": {"UniCNS-UTF16-H", "CNS1", 7},

	"HIRA": {"UniJIS-UTF16-H", "Japan1", 7},
	"KANA": {"UniJIS-UTF16-H", "Japan1", 7},
	"JPAN": {"UniJIS-UTF16-H", "Japan1", 7},

	"HANG": {"UniKS-UTF16-H", "Korea1", 1},
	"KORE": {"UniKS-UTF16-H", "Korea1", 1},
}

func SupportedScript(s string) bool { _ = "STUB: not implemented"; return false }

func CJKEncoding(s string) bool { _ = "STUB: not implemented"; return false }

func ScriptForEncoding(enc string) string { _ = "STUB: not implemented"; return "" }

func fontDescriptorIndRefs(fd types.Dict, lang, fontName string, font *model.FontResource) error {
	_ = "STUB: not implemented"
	return nil
}

func fontDictName(d types.Dict) string { _ = "STUB: not implemented"; return "" }

func IndRefsForUserfontUpdate(xRefTable *model.XRefTable, d types.Dict, lang string, font *model.FontResource) error {
	_ = "STUB: not implemented"
	return nil
}

func fontPhase(fontName, phase string) string { _ = "STUB: not implemented"; return "" }

func missingUserfontReference(fontName, phase, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func requireFontXRef(xRefTable *model.XRefTable, fontName, phase string) error {
	_ = "STUB: not implemented"
	return nil
}

func insertFontObject(xRefTable *model.XRefTable, fontName, phase string, obj types.Object) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ttfFontName(ttf font.TTFLight, fontName string) string { _ = "STUB: not implemented"; return "" }

func validateEmbeddingMetrics(ttf font.TTFLight, fontName, phase string) error {
	_ = "STUB: not implemented"
	return nil
}

func flateEncodedStreamIndRef(xRefTable *model.XRefTable, fontName, phase string, data []byte) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ttfFontFile(xRefTable *model.XRefTable, fontName string) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type referencedFontObjectType uint8

const (
	fontStreamObject referencedFontObjectType = iota
	fontArrayObject
	fontDictObject
)

func referencedFontObject(xRefTable *model.XRefTable, fontName, phase string, indRef *types.IndirectRef, expected referencedFontObjectType) (*model.XRefTableEntry, types.Object, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.Object), nil
}

func ttfSubFontFile(xRefTable *model.XRefTable, fontName string, indRef *types.IndirectRef) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PDFDocEncoding(xRefTable *model.XRefTable) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CoreFontDict(xRefTable *model.XRefTable, coreFontName string) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CIDSet(xRefTable *model.XRefTable, ttf font.TTFLight, fontName string, indRef *types.IndirectRef) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ttfFontDescriptorFlags(ttf font.TTFLight) uint32 { _ = "STUB: not implemented"; return 0 }

func CIDFontFile(xRefTable *model.XRefTable, fontName string, subFont bool) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CIDFontDescriptor(xRefTable *model.XRefTable, ttf font.TTFLight, fontName, baseFontName, fontLang string, embed bool) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewFontDescriptor(xRefTable *model.XRefTable, ttf font.TTFLight, fontName, fontLang string) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func wArr(ttf font.TTFLight, from, thru int) types.Array {
	_ = "STUB: not implemented"
	return *new(types.Array)
}

func prepGids(xRefTable *model.XRefTable, ttf font.TTFLight, fontName string, used bool) ([]int, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func handleEqualWidths(w, w0, wl, g, g0, gl *int, a *types.Array, skip, equalWidths *bool) {
	_ = "STUB: not implemented"
	return
}

func finalizeWidths(ttf font.TTFLight, w0, g0, gl int, skip, equalWidths bool, a *types.Array) {
	_ = "STUB: not implemented"
	return
}

func calcWidthArray(xRefTable *model.XRefTable, ttf font.TTFLight, fontName string, used bool) types.Array {
	_ = "STUB: not implemented"
	return *new(types.Array)
}

func CIDWidths(xRefTable *model.XRefTable, ttf font.TTFLight, fontName string, subFont bool, indRef *types.IndirectRef) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Widths(xRefTable *model.XRefTable, ttf font.TTFLight, first, last int) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func bf(b *bytes.Buffer, ttf font.TTFLight, usedGIDs map[uint16]bool, subFont bool) {
	_ = "STUB: not implemented"
	return
}

func toUnicodeCMap(xRefTable *model.XRefTable, ttf font.TTFLight, fontName string, subFont bool, indRef *types.IndirectRef) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	errCorruptCMap = errors.New("corrupt CMap")

	ErrCorruptFontDict = errors.New("corrupt fontDict")
)

const cmapCodeSpaceEnd = "endcodespacerange"

func scanCMapLine(scanner *bufio.Scanner, phase string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func cmapBFCharCount(line string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func cmapGID(line string) (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

func cmapSingleToken(line, want string) error { _ = "STUB: not implemented"; return nil }

func usedGIDsFromCMap(cMap string) ([]uint16, error) { _ = "STUB: not implemented"; return nil, nil }

func validateUserfontUpdateReferences(fontName string, f model.FontResource) error {
	_ = "STUB: not implemented"
	return nil
}

func UpdateUserfont(xRefTable *model.XRefTable, fontName string, f model.FontResource) error {
	_ = "STUB: not implemented"
	return nil
}

func UpdateUserfonts(xRefTable *model.XRefTable, fonts map[string]types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func usedGIDsFromCMapIndRef(xRefTable *model.XRefTable, fontName string, cmapIndRef types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func subFontPrefix() string { _ = "STUB: not implemented"; return "" }

func CIDFontDict(xRefTable *model.XRefTable, ttf font.TTFLight, fontName, baseFontName, lang string, parms *cjk) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func type0FontDictData(xRefTable *model.XRefTable, ttf font.TTFLight, fontName, lang, script string) (types.Dict, bool, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), false, nil
}

func type0FontDict(xRefTable *model.XRefTable, fontName, lang, script string, indRef *types.IndirectRef) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func trueTypeFontDict(xRefTable *model.XRefTable, fontName, fontLang string) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CJK(script, lang string) bool { _ = "STUB: not implemented"; return false }

func RTL(lang string) bool { _ = "STUB: not implemented"; return false }

func EnsureFontDict(xRefTable *model.XRefTable, fontName, lang, script string, field bool, indRef *types.IndirectRef) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FontResources(xRefTable *model.XRefTable, fm model.FontMap) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func Name(xRefTable *model.XRefTable, fontDict types.Dict, objNumber int) (prefix, fontName string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func Lang(xRefTable *model.XRefTable, fontDict types.Dict) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func trivialFontDescriptor(xRefTable *model.XRefTable, fontDict types.Dict, objNr int) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func FontDescriptor(xRefTable *model.XRefTable, fontDict types.Dict, objNr int) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func Embedded(xRefTable *model.XRefTable, fontDict types.Dict, objNr int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
