package model

import (
	"errors"
)

var ErrMaxRecursionDepthExceeded = errors.New("max recursion depth exceeded")

var ErrPageTreeCycle = errors.New("circular page tree")

var ErrPageTreeDuplicate = errors.New("duplicate page tree node")

var ErrFormFieldCycle = errors.New("circular form field tree")

var ErrStructureTreeCycle = errors.New("circular structure tree")

var ErrActionCycle = errors.New("circular action chain")

var ErrBeadCycle = errors.New("circular bead chain")

func (xRefTable *XRefTable) MaxRecursionDepth() int { _ = "STUB: not implemented"; return 0 }

func (xRefTable *XRefTable) CheckRecursionDepth(name string, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func CheckRecursionDepth(name string, depth, maxDepth int) error {
	_ = "STUB: not implemented"
	return nil
}

type PageTreeVisit struct {
	ancestors map[int]bool
	seen      map[int]bool
}

func NewPageTreeVisit() *PageTreeVisit { _ = "STUB: not implemented"; return nil }

func (v *PageTreeVisit) Enter(objNr int) error { _ = "STUB: not implemented"; return nil }

func (v *PageTreeVisit) Leave(objNr int) { _ = "STUB: not implemented"; return }

type FormFieldVisit struct {
	ancestors map[int]bool
}

func NewFormFieldVisit() *FormFieldVisit { _ = "STUB: not implemented"; return nil }

func (v *FormFieldVisit) Enter(objNr int) error { _ = "STUB: not implemented"; return nil }

func (v *FormFieldVisit) Check(objNr int) error { _ = "STUB: not implemented"; return nil }

func (v *FormFieldVisit) Leave(objNr int) { _ = "STUB: not implemented"; return }

type StructureTreeVisit struct {
	ancestors map[int]bool
}

func NewStructureTreeVisit() *StructureTreeVisit { _ = "STUB: not implemented"; return nil }

func (v *StructureTreeVisit) Enter(objNr int) error { _ = "STUB: not implemented"; return nil }

func (v *StructureTreeVisit) Leave(objNr int) { _ = "STUB: not implemented"; return }

type ActionVisit struct {
	ancestors map[int]bool
}

func NewActionVisit() *ActionVisit { _ = "STUB: not implemented"; return nil }

func (v *ActionVisit) Enter(objNr int) error { _ = "STUB: not implemented"; return nil }

func (v *ActionVisit) Leave(objNr int) { _ = "STUB: not implemented"; return }

type BeadVisit struct {
	ancestors map[int]bool
}

func NewBeadVisit() *BeadVisit { _ = "STUB: not implemented"; return nil }

func (v *BeadVisit) Enter(objNr int) error { _ = "STUB: not implemented"; return nil }

func (v *BeadVisit) Leave(objNr int) { _ = "STUB: not implemented"; return }
