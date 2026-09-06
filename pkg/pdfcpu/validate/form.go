package validate

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func validateAppearanceSubDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAppearanceDictEntry(xRefTable *model.XRefTable, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAppearanceEntry(xRefTable *model.XRefTable, d types.Dict, entryName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAppearanceDict(xRefTable *model.XRefTable, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDA(s string) bool { _ = "STUB: not implemented"; return false }

func validateDARelaxed(s string) bool { _ = "STUB: not implemented"; return false }

func validateFormFieldDA(xRefTable *model.XRefTable, d types.Dict, dictName string, terminalNode bool, outFieldType *types.Name, requiresDA bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func detectRectArray(xRefTable *model.XRefTable, d types.Dict, dictName string) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func cacheSig(xRefTable *model.XRefTable, d types.Dict, dictName string, form bool, objNr, incr int) error {
	_ = "STUB: not implemented"
	return nil
}

func isTextField(ft *types.Name) bool { _ = "STUB: not implemented"; return false }

func validateV(xRefTable *model.XRefTable, objNr, incr int, d types.Dict, dictName string, terminalNode, textField, oneKid bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDV(xRefTable *model.XRefTable, d types.Dict, dictName string, terminalNode, textField, oneKid bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFormFieldType(xRefTable *model.XRefTable) func(string) bool {
	_ = "STUB: not implemented"
	return nil
}

func validateFormFieldDictEntries(xRefTable *model.XRefTable, objNr, incr int, d types.Dict, terminalNode, oneKid bool, inFieldType *types.Name, requiresDA bool) (outFieldType *types.Name, hasDA bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func validateFormFieldParts(xRefTable *model.XRefTable, objNr, incr int, d types.Dict, inFieldType *types.Name, requiresDA bool) error {
	_ = "STUB: not implemented"
	return nil
}

func isWidget(d types.Dict) bool { _ = "STUB: not implemented"; return false }

func validateFormFieldKids(
	xRefTable *model.XRefTable,
	objNr,
	incr int,
	d types.Dict,
	o types.Object,
	inFieldType *types.Name,
	requiresDA bool,
	depth int,
	visit *model.FormFieldVisit,
	specViolations *[]error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFormFieldDict(xRefTable *model.XRefTable, ir types.IndirectRef, inFieldType *types.Name, requiresDA bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFormFieldDictDepth(
	xRefTable *model.XRefTable,
	ir types.IndirectRef,
	inFieldType *types.Name,
	requiresDA bool,
	depth int,
	visit *model.FormFieldVisit,
	specViolations *[]error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFormFields(xRefTable *model.XRefTable, arr types.Array, requiresDA bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFormCO(xRefTable *model.XRefTable, arr types.Array, sinceVersion model.Version, requiresDA bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFormXFA(xRefTable *model.XRefTable, d types.Dict, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateQ(i int) bool { _ = "STUB: not implemented"; return false }

func validateFormEntryCO(xRefTable *model.XRefTable, d types.Dict, sinceVersion model.Version, requiresDA bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFormEntryDR(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFormEntries(xRefTable *model.XRefTable, d types.Dict, dictName string, requiresDA bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func handleSelfReferentialAcroForm(xRefTable *model.XRefTable, rootDict types.Dict) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func validateForm(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func locateAnnForAPAndRect(d types.Dict, r *types.Rectangle, pageAnnots map[int]model.PgAnnots) *types.IndirectRef {
	_ = "STUB: not implemented"
	return nil
}

func pageAnnotIndRefForAcroField(xRefTable *model.XRefTable, indRef types.IndirectRef) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fixFormFieldsArray(xRefTable *model.XRefTable, arr types.Array) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func validateFormFieldsAgainstPageAnnotations(xRefTable *model.XRefTable) error {
	_ = "STUB: not implemented"
	return nil
}
