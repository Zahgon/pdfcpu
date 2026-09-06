package validate

import (
	"errors"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

var errUnsupportedPDFObject = errors.New("unsupported PDF object")

type invalidStructElementKError struct {
	err error
}

func (e *invalidStructElementKError) Error() string { _ = "STUB: not implemented"; return "" }

func handleInvalidStructElementK(
	xRefTable *model.XRefTable,
	err error,
	specViolations *[]error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func showDigestedSpecViolations(xRefTable *model.XRefTable, specViolations []error) {
	_ = "STUB: not implemented"
	return
}

func validateMarkedContentReferenceDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateObjectReferenceDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func enterStructureTreeObject(visit *model.StructureTreeVisit, o types.Object) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func validateStructElementKArrayElement(
	xRefTable *model.XRefTable,
	o types.Object,
	useIDs bool,
	depth int,
	visit *model.StructureTreeVisit,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructElementDictEntryKArrayElement(
	xRefTable *model.XRefTable,
	rawObject types.Object,
	index int,
	useIDs bool,
	depth int,
	visit *model.StructureTreeVisit,
	specViolations *[]error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructElementDictEntryKArrayDepth(
	xRefTable *model.XRefTable,
	a types.Array,
	useIDs bool,
	depth int,
	visit *model.StructureTreeVisit,
	specViolations *[]error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructElementDictEntryKArray(xRefTable *model.XRefTable, a types.Array, useIDs bool, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructElementDictEntryKDepth(
	xRefTable *model.XRefTable,
	rawObject types.Object,
	useIDs bool,
	depth int,
	visit *model.StructureTreeVisit,
	specViolations *[]error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructElementDictEntryK(xRefTable *model.XRefTable, o types.Object, useIDs bool, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func processStructElementDictPgEntry(xRefTable *model.XRefTable, ir types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructElementDictEntryA(xRefTable *model.XRefTable, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructElementDictEntryC(xRefTable *model.XRefTable, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructElementDictEntryP(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructElementDictEntryPg(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructElementDictEntryS(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructElementDictPart1(
	xRefTable *model.XRefTable,
	d types.Dict,
	dictName string,
	useIDs bool,
	depth int,
	visit *model.StructureTreeVisit,
	specViolations *[]error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructElementLang(xRefTable *model.XRefTable, d types.Dict, dictName string, sinceVersion model.Version) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func validateStructElementDictPart2(xRefTable *model.XRefTable, d types.Dict, dictName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func validateStructElementDict(xRefTable *model.XRefTable, d types.Dict, useIDs bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructElementDictDepth(
	xRefTable *model.XRefTable,
	d types.Dict,
	useIDs bool,
	depth int,
	visit *model.StructureTreeVisit,
) error {
	_ = "STUB: not implemented"
	return nil
}

func handleInvalidStructTreeObject(xRefTable *model.XRefTable, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructTreeRootDictEntryKArrayElement(
	xRefTable *model.XRefTable,
	rawObject types.Object,
	index int,
	useIDs bool,
	visit *model.StructureTreeVisit,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructTreeRootDictEntryKArrayDepth(
	xRefTable *model.XRefTable,
	a types.Array,
	useIDs bool,
	visit *model.StructureTreeVisit,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructTreeRootDictEntryKArray(xRefTable *model.XRefTable, a types.Array, useIDs bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructTreeRootDictEntryKDepth(
	xRefTable *model.XRefTable,
	rawObject types.Object,
	useIDs bool,
	visit *model.StructureTreeVisit,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructTreeRootDictEntryK(xRefTable *model.XRefTable, o types.Object, useIDs bool) error {
	_ = "STUB: not implemented"
	return nil
}

func processStructTreeClassMapDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructTreeRootDictEntryParentTree(xRefTable *model.XRefTable, ir *types.IndirectRef, useIDs bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructTreeRootDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructTree(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}
