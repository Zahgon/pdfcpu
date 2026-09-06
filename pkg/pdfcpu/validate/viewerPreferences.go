package validate

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func validateDirection(xRefTable *model.XRefTable, d types.Dict, dictName string, vp *model.ViewerPreferences) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePageBoundaries(xRefTable *model.XRefTable, d types.Dict, dictName string, vp *model.ViewerPreferences) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePrintPageRange(xRefTable *model.XRefTable, d types.Dict, dictName string, vp *model.ViewerPreferences) error {
	_ = "STUB: not implemented"
	return nil
}

func validateEnforcePrintScaling(xRefTable *model.XRefTable, d types.Dict, dictName string, vp *model.ViewerPreferences) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePrinterPreferences(xRefTable *model.XRefTable, d types.Dict, dictName string, vp *model.ViewerPreferences) error {
	_ = "STUB: not implemented"
	return nil
}

func validateViewerPreferencesFlags(xRefTable *model.XRefTable, d types.Dict, dictName string, vp *model.ViewerPreferences) error {
	_ = "STUB: not implemented"
	return nil
}

func validateViewerPreferences(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}
