package validate

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func validateGoToActionDict(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateGoToRActionDict(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateTargetDictEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateGoToEActionDict(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateWinDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateLaunchActionDict(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDestinationThreadEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDestinationBeadEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateThreadActionDict(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func hasURIForChecking(xRefTable *model.XRefTable, s string) bool {
	_ = "STUB: not implemented"
	return false
}

func validateURIActionDict(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSoundDictEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSoundActionDict(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMovieStartOrDurationEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMovieActivationDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMovieActionDict(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateHideActionDictEntryT(xRefTable *model.XRefTable, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func validateHideActionDict(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNamedActionDict(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSubmitFormActionDict(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateResetFormActionDict(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateImportDataActionDict(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateJavaScript(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateJavaScriptActionDict(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSetOCGStateActionDict(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateRenditionActionDict(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateTransActionDict(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateGoTo3DViewActionDict(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateActionDictCore(xRefTable *model.XRefTable, n *types.Name, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateActionDictObject(xRefTable *model.XRefTable, d types.Dict, o types.Object, context string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateActionDictObjectDepth(
	xRefTable *model.XRefTable,
	d types.Dict,
	o types.Object,
	context string,
	depth int,
	visit *model.ActionVisit,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNextAction(xRefTable *model.XRefTable, o types.Object, depth int, visit *model.ActionVisit) error {
	_ = "STUB: not implemented"
	return nil
}

func validateActionDict(xRefTable *model.XRefTable, d types.Dict, depth int, visit *model.ActionVisit) error {
	_ = "STUB: not implemented"
	return nil
}

func validateRootAdditionalActions(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAdditionalActions(xRefTable *model.XRefTable, dict types.Dict, dictName, entryName string, required bool, sinceVersion model.Version, source string) error {
	_ = "STUB: not implemented"
	return nil
}
