package model

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type BookletType int

const (
	Booklet BookletType = iota
	BookletAdvanced
	BookletPerfectBound
)

func (b BookletType) String() string { _ = "STUB: not implemented"; return "" }

type BookletBinding int

const (
	LongEdge BookletBinding = iota
	ShortEdge
)

func (b BookletBinding) String() string { _ = "STUB: not implemented"; return "" }

type BookletPage struct {
	Number int
	Rotate bool
}

func drawGuideLineLabel(xRefTable *XRefTable, w io.Writer, x, y float64, s string, mb *types.Rectangle, fm FontMap, rot int) error {
	_ = "STUB: not implemented"
	return nil
}

func drawScissors(xRefTable *XRefTable, w io.Writer, isVerticalCut bool, horzCutYpos float64, mb *types.Rectangle, fm FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

type cutOrFold int

const (
	none cutOrFold = iota
	cut
	fold
)

func (c cutOrFold) String(nup *NUp) string { _ = "STUB: not implemented"; return "" }

func getCutFolds(nup *NUp) (horizontal cutOrFold, vertical cutOrFold) {
	_ = "STUB: not implemented"
	return *new(cutOrFold), *new(cutOrFold)
}

func drawGuideHorizontal(xRefTable *XRefTable, w io.Writer, y, width float64, cutOrFold cutOrFold, nup *NUp, mb *types.Rectangle, fm FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func drawGuideVertical(xRefTable *XRefTable, w io.Writer, x, height float64, cutOrFold cutOrFold, nup *NUp, mb *types.Rectangle, fm FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func DrawBookletGuides(xRefTable *XRefTable, nup *NUp, w io.Writer) (FontMap, error) {
	_ = "STUB: not implemented"
	return *new(FontMap), nil
}
