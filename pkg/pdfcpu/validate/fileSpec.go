package validate

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func validateFileSpecString(s string) bool { _ = "STUB: not implemented"; return false }

func validateURLString(s string) bool { _ = "STUB: not implemented"; return false }

func validateEmbeddedFileStreamMacParameterDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateEmbeddedFileStreamParameterDict(xRefTable *model.XRefTable, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func validateEmbeddedFileStreamDict(xRefTable *model.XRefTable, sd *types.StreamDict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFileSpecDictEntriesEFAndRFKeys(k string) bool { _ = "STUB: not implemented"; return false }

func validateFileSpecDictEntryEFDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateRFDictFilesArray(xRefTable *model.XRefTable, a types.Array) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFileSpecDictEntriesEFAndRF(xRefTable *model.XRefTable, efDict, rfDict types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func requiredF(dosFound, macFound, unixFound bool) bool { _ = "STUB: not implemented"; return false }

func validateFileSpecDictEFAndRF(xRefTable *model.XRefTable, d types.Dict, dictName string, hasEP bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFileSpecDictPart1(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFileSpecDictPart2(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFileSpecDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFileSpecification(xRefTable *model.XRefTable, o types.Object) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func validateURLSpecification(xRefTable *model.XRefTable, o types.Object) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func validateFileSpecEntry(xRefTable *model.XRefTable, d types.Dict, dictName string, entryName string, required bool, sinceVersion model.Version) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func validateURLSpecEntry(xRefTable *model.XRefTable, d types.Dict, dictName string, entryName string, required bool, sinceVersion model.Version) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func validateFileSpecificationOrFormObject(xRefTable *model.XRefTable, obj types.Object) error {
	_ = "STUB: not implemented"
	return nil
}
