package model

import "github.com/pkg/errors"

var ErrMaxRecursionDepthExceeded = errors.New("pdfcpu: max recursion depth exceeded")

var ErrPageTreeCycle = errors.New("pdfcpu: circular page tree")

var ErrPageTreeDuplicate = errors.New("pdfcpu: duplicate page tree node")

var ErrFormFieldCycle = errors.New("pdfcpu: circular form field tree")

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
