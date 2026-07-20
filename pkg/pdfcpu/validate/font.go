package validate

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"github.com/pkg/errors"
)

var ErrMissingFont = errors.New("pdfcpu: missing font dict")

func validateStandardType1Font(s string) bool { _ = "STUB: not implemented"; return false }

func validateFontFile3SubType(sd *types.StreamDict, fontType string, relaxed bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFontFile(xRefTable *model.XRefTable, d types.Dict, dictName string, entryName string, fontType string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFontDescriptorType(xRefTable *model.XRefTable, d types.Dict) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func validateFontDescriptorFontName(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFontDescriptorFontFamily(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFontDescriptorFontStretch(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFontDescriptorFontWeight(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFontDescriptorFontFlags(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFontDescriptorFontBox(xRefTable *model.XRefTable, d types.Dict, dictName, fontDictType string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFontDescriptorItalicAngle(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFontDescriptorPart1(xRefTable *model.XRefTable, d types.Dict, dictName, fontDictType string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFontDescriptorPart2(xRefTable *model.XRefTable, d types.Dict, dictName, fontDictType string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFontDescriptorFontFile(xRefTable *model.XRefTable, d types.Dict, dictName, fontDictType string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func validateFontDescriptor(xRefTable *model.XRefTable, d types.Dict, fontDictName string, fontDictType string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFontEncoding(xRefTable *model.XRefTable, d types.Dict, dictName string, required bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateTrueTypeFontDict(xRefTable *model.XRefTable, d types.Dict) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func validateCIDToGIDMap(xRefTable *model.XRefTable, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCIDFontGlyphWidths(xRefTable *model.XRefTable, d types.Dict, dictName string, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCIDFontDictEntryCIDSystemInfo(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCIDFontDictEntryCIDToGIDMap(xRefTable *model.XRefTable, d types.Dict, isCIDFontType2 bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCIDFontDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDescendantFonts(xRefTable *model.XRefTable, d types.Dict, fontDictName string, required bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateType0FontDict(xRefTable *model.XRefTable, d types.Dict) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func validateType1FontDict(xRefTable *model.XRefTable, d types.Dict) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func validateCharProcsDict(xRefTable *model.XRefTable, d types.Dict, dictName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateUseCMapEntry(xRefTable *model.XRefTable, d types.Dict, dictName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCIDSystemInfoDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCMapStreamDict(xRefTable *model.XRefTable, sd *types.StreamDict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateType0FontEncoding(xRefTable *model.XRefTable, d types.Dict, dictName string, required bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateType3FontDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func _validateFontDict(xRefTable *model.XRefTable, d types.Dict, isIndRef bool, indRef types.IndirectRef) (fontName string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func validateFontDict(xRefTable *model.XRefTable, isIndRef bool, indRef types.IndirectRef) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func fixFontObjNr(m1 map[string]string, m2 map[string]types.IndirectRef, d types.Dict) {
	_ = "STUB: not implemented"
	return
}

func validateFontResourceDict(xRefTable *model.XRefTable, o types.Object, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}
