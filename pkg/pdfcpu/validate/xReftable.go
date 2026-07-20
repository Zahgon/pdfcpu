package validate

import (
	"net/http"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func XRefTable(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func fixInfoDict(xRefTable *model.XRefTable, rootDict types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func metaDataModifiedAfterInfoDict(xRefTable *model.XRefTable) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func setRootVersion(xRefTable *model.XRefTable, s string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateRootVersion(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateExtensions(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePageLabels(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNames(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNamedDestinations(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func pageLayoutValidator(v model.Version) func(s string) bool {
	_ = "STUB: not implemented"
	return nil
}

func validatePageLayout(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func pageModeValidator(v model.Version) func(s string) bool { _ = "STUB: not implemented"; return nil }

func validatePageMode(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOpenAction(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateURI(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMarkInfo(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateLang(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCaptureCommandDictArray(xRefTable *model.XRefTable, a types.Array) error {
	_ = "STUB: not implemented"
	return nil
}

func validateWebCaptureInfoDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSpiderInfo(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOutputIntentDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOutputIntents(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePieceDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateRootPieceInfo(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePieceInfo(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) (hasPieceInfo bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func validatePermissions(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateLegal(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateRequirementDict(xRefTable *model.XRefTable, d types.Dict, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateRequirements(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCollectionFieldDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCollectionSchemaDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCollectionSortDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateInitialView(s string) bool { _ = "STUB: not implemented"; return false }

func validateCollection(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNeedsRendering(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDSS(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAF(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDPartRoot(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func logURIError(xRefTable *model.XRefTable, pages []int) { _ = "STUB: not implemented"; return }

func checkLinks(xRefTable *model.XRefTable, client http.Client, pages []int) bool {
	_ = "STUB: not implemented"
	return false
}

func checkForBrokenLinks(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func validateRootObject(ctx *model.Context, rootDict types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}
