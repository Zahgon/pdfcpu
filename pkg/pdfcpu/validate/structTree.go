package validate

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func validateMarkedContentReferenceDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateObjectReferenceDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructElementKArrayElement(xRefTable *model.XRefTable, o types.Object, useIDs bool, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructElementDictEntryKArray(xRefTable *model.XRefTable, a types.Array, useIDs bool, depth int) error {
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

func validateStructElementDictPart1(xRefTable *model.XRefTable, d types.Dict, dictName string, useIDs bool, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructElementDictPart2(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructElementDict(xRefTable *model.XRefTable, d types.Dict, useIDs bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructElementDictDepth(xRefTable *model.XRefTable, d types.Dict, useIDs bool, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStructTreeRootDictEntryKArray(xRefTable *model.XRefTable, a types.Array, useIDs bool) error {
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
