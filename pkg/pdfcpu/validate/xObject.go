package validate

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

const (
	ExcludePatternCS = true

	IncludePatternCS = false

	isAlternateImageStreamDict   = true
	isNoAlternateImageStreamDict = false
)

func validateReferenceDictPageEntry(xRefTable *model.XRefTable, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func validateReferenceDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOPIDictV13Part1(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOPIDictV13Part2(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOPIDictV13(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOPIDictInks(xRefTable *model.XRefTable, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOPIDictV20(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOPIVersionDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMaskStreamDict(xRefTable *model.XRefTable, sd *types.StreamDict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMaskEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAlternateImageStreamDicts(xRefTable *model.XRefTable, d types.Dict, dictName string, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateImageStreamDictPart1(xRefTable *model.XRefTable, sd *types.StreamDict, dictName string) (isImageMask bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func validateImageStreamDictPart2(xRefTable *model.XRefTable, sd *types.StreamDict, dictName string, isImageMask, isAlternate bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateImageStreamDict(xRefTable *model.XRefTable, sd *types.StreamDict, isAlternate bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFormStreamDictPart1(xRefTable *model.XRefTable, sd *types.StreamDict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateEntryOPI(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFormStreamDictPart2(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFormStreamDict(xRefTable *model.XRefTable, sd *types.StreamDict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateXObjectType(xRefTable *model.XRefTable, sd *types.StreamDict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateXObjectStreamDictMissingSubtype(xRefTable *model.XRefTable, sd *types.StreamDict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateXObjectStreamDictSubtype(xRefTable *model.XRefTable, sd *types.StreamDict, subtype types.Name) error {
	_ = "STUB: not implemented"
	return nil
}

func validateXObjectStreamDict(xRefTable *model.XRefTable, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func validateGroupAttributesDict(xRefTable *model.XRefTable, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func validateXObjectResourceDict(xRefTable *model.XRefTable, o types.Object, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}
