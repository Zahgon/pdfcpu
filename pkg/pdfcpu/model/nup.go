package model

import (
	"bytes"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type orientation int

func (o orientation) String() string { _ = "STUB: not implemented"; return "" }

const (
	RightDown orientation = iota
	DownRight
	LeftDown
	DownLeft
)

type BorderStyling struct {
	Color     *color.SimpleColor
	LineStyle *types.LineJoinStyle
	Width     float64
}

type NUp struct {
	PageDim         *types.Dim
	PageSize        string
	UserDim         bool
	Orient          orientation
	Enforce         bool
	Grid            *types.Dim
	PageGrid        bool
	ImgInputFile    bool
	Margin          float64
	Border          bool
	BorderOnCropbox *BorderStyling
	BookletGuides   bool
	MultiFolio      bool
	FolioSize       int
	BookletType     BookletType
	BookletBinding  BookletBinding
	InpUnit         types.DisplayUnit
	BgColor         *color.SimpleColor
}

func DefaultNUpConfig() *NUp { _ = "STUB: not implemented"; return nil }

func (nup NUp) String() string { _ = "STUB: not implemented"; return "" }

func (nup NUp) operation() string { _ = "STUB: not implemented"; return "" }

func (nup NUp) N() int { _ = "STUB: not implemented"; return 0 }

func (nup NUp) IsTopFoldBinding() bool { _ = "STUB: not implemented"; return false }

func (nup NUp) IsBooklet() bool { _ = "STUB: not implemented"; return false }

func (nup NUp) RectsForGrid() []*types.Rectangle { _ = "STUB: not implemented"; return nil }

func createNUpFormForPDF(xRefTable *XRefTable, resDict *types.IndirectRef, content []byte, cropBox *types.Rectangle) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NUpTilePDFBytes(wr io.Writer, rSrc, rDest *types.Rectangle, formResID string, nup *NUp, rotate bool) {
	_ = "STUB: not implemented"
	return
}

func translationForPageRotation(pageRot int, w, h float64) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func ContentBytesForPageRotation(rot int, w, h float64) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *Context) NUpTilePDFBytesForPDF(
	pageNr int,
	formsResDict types.Dict,
	buf *bytes.Buffer,
	rDest *types.Rectangle,
	nup *NUp,
	rotate bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *Context) TilePDFBytesForImposition(
	operation string,
	pageNr int,
	formsResDict types.Dict,
	buf *bytes.Buffer,
	rDest *types.Rectangle,
	nup *NUp,
	rotate bool) error {
	_ = "STUB: not implemented"
	return nil
}

func AppendPageTree(d1 *types.IndirectRef, countd1 int, d2 types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}
