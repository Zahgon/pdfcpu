package validate

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func validateDeviceColorSpaceName(s string) bool { _ = "STUB: not implemented"; return false }

func validateAllColorSpaceNamesExceptPattern(s string) bool {
	_ = "STUB: not implemented"
	return false
}

func validateCalGrayColorSpace(xRefTable *model.XRefTable, a types.Array, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCalRGBColorSpace(xRefTable *model.XRefTable, a types.Array, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateLabColorSpace(xRefTable *model.XRefTable, a types.Array, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAlternateColorSpaceEntryForICC(xRefTable *model.XRefTable, d types.Dict, dictName string, entryName string, required bool, excludePatternCS bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateICCBasedColorSpace(xRefTable *model.XRefTable, a types.Array, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateIndexedColorSpaceLookuptable(xRefTable *model.XRefTable, o types.Object, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateIndexedColorSpace(xRefTable *model.XRefTable, a types.Array, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePatternColorSpace(xRefTable *model.XRefTable, a types.Array, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSeparationColorSpace(xRefTable *model.XRefTable, a types.Array, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDeviceNColorSpaceColorantsDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDeviceNColorSpaceProcessDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDeviceNColorSpaceSoliditiesDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDeviceNColorSpaceDotGainDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDeviceNColorSpaceMixingHintsDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDeviceNColorSpaceAttributesDict(xRefTable *model.XRefTable, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDeviceNColorSpace(xRefTable *model.XRefTable, a types.Array, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCSArray(xRefTable *model.XRefTable, a types.Array, csName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateColorSpaceArraySubset(xRefTable *model.XRefTable, a types.Array, cs []string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateColorSpaceArray(xRefTable *model.XRefTable, a types.Array, excludePatternCS bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func validateColorSpace(xRefTable *model.XRefTable, o types.Object, excludePatternCS bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateColorSpaceEntry(xRefTable *model.XRefTable, d types.Dict, dictName string, entryName string, required bool, excludePatternCS bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateColorSpaceResourceDict(xRefTable *model.XRefTable, o types.Object, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}
