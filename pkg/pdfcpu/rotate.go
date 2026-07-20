package pdfcpu

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func rotatePage(xRefTable *model.XRefTable, i, j int) error { _ = "STUB: not implemented"; return nil }

func RotatePages(ctx *model.Context, selectedPages types.IntSet, rotation int) error {
	_ = "STUB: not implemented"
	return nil
}
