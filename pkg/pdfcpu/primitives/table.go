package primitives

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/matrix"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type TableHeader struct {
	Values          []string
	ColAnchors      []string
	colAnchors      []types.Anchor
	ColPaddings     []*Padding
	LineHeight      int    `json:"lheight"`
	BackgroundColor string `json:"bgCol"`
	bgCol           *color.SimpleColor
	Font            *FormFont
	RTL             bool
}

func (th *TableHeader) validateColumnPaddings(cols int) error {
	_ = "STUB: not implemented"
	return nil
}

func (th *TableHeader) validate(pdf *PDF, cols int) error { _ = "STUB: not implemented"; return nil }

func applyTextDescriptorPadding(td *model.TextDescriptor, p *Padding) {
	_ = "STUB: not implemented"
	return
}

func (th *TableHeader) calcColumnPadding(td *model.TextDescriptor, col int) {
	_ = "STUB: not implemented"
	return
}

type Table struct {
	pdf             *PDF
	content         *Content
	Name            string
	Values          [][]string
	Position        [2]float64 `json:"pos"`
	x, y            float64
	Dx, Dy          float64
	Anchor          string
	anchor          types.Anchor
	anchored        bool
	Width           float64
	Rows, Cols      int
	ColWidths       []int
	ColAnchors      []string
	colAnchors      []types.Anchor
	ColPaddings     []*Padding
	LineHeight      int `json:"lheight"`
	Font            *FormFont
	Margin          *Margin
	Border          *Border
	Padding         *Padding
	BackgroundColor string `json:"bgCol"`
	OddColor        string `json:"oddCol"`
	EvenColor       string `json:"evenCol"`
	bgCol           *color.SimpleColor
	oddCol          *color.SimpleColor
	evenCol         *color.SimpleColor
	RTL             bool
	Rotation        float64 `json:"rot"`
	Grid            bool
	Hide            bool
	Header          *TableHeader
}

func (t *Table) Height() float64 { _ = "STUB: not implemented"; return 0 }

func (t *Table) validateAnchor() error { _ = "STUB: not implemented"; return nil }

func (t *Table) validateColWidths() error { _ = "STUB: not implemented"; return nil }

func (t *Table) validateColAnchors() error { _ = "STUB: not implemented"; return nil }

func (t *Table) validateColPaddings() error { _ = "STUB: not implemented"; return nil }

func (t *Table) validateColumns() error { _ = "STUB: not implemented"; return nil }

func (t *Table) validateValues() error { _ = "STUB: not implemented"; return nil }

func (t *Table) validateFont() error { _ = "STUB: not implemented"; return nil }

func (t *Table) validateMargin() error { _ = "STUB: not implemented"; return nil }

func (t *Table) validateBorder() error { _ = "STUB: not implemented"; return nil }

func (t *Table) validatePadding() error { _ = "STUB: not implemented"; return nil }

func (t *Table) validateBackgroundColor() error { _ = "STUB: not implemented"; return nil }

func (t *Table) validateOddColor() error { _ = "STUB: not implemented"; return nil }

func (t *Table) validateEvenColor() error { _ = "STUB: not implemented"; return nil }

func (t *Table) validateColors() error { _ = "STUB: not implemented"; return nil }

func (t *Table) validate() error { _ = "STUB: not implemented"; return nil }

func (t *Table) font(name string) *FormFont { _ = "STUB: not implemented"; return nil }

func (t *Table) margin(name string) *Margin { _ = "STUB: not implemented"; return nil }

func (t *Table) border(name string) *Border { _ = "STUB: not implemented"; return nil }

func (t *Table) padding(name string) *Padding { _ = "STUB: not implemented"; return nil }

func (t *Table) mergeInAnchor(t0 *Table) { _ = "STUB: not implemented"; return }

func (t *Table) mergeIn(t0 *Table) { _ = "STUB: not implemented"; return }

func (t *Table) calcFont() error { _ = "STUB: not implemented"; return nil }

func (t *Table) calcBorder() (float64, *color.SimpleColor, types.LineJoinStyle, error) {
	_ = "STUB: not implemented"
	return 0, nil, *new(types.LineJoinStyle), nil
}

func (t *Table) calcMargin() (float64, float64, float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, nil
}

func (t *Table) calcTransform(mLeft, mBottom, mRight, mTop, bWidth float64) (matrix.Matrix, *types.Rectangle) {
	_ = "STUB: not implemented"
	return *new(matrix.Matrix), nil
}

func (t *Table) renderBackground(p *model.Page, bWidth float64, r *types.Rectangle) {
	_ = "STUB: not implemented"
	return
}

func (t *Table) prepareColWidths(bWidth float64) []float64 { _ = "STUB: not implemented"; return nil }

func (t *Table) renderGrid(p *model.Page, colWidths []float64, bWidth float64, bCol *color.SimpleColor, r *types.Rectangle) {
	_ = "STUB: not implemented"
	return
}

func (t *Table) cellLowerLeft(r *types.Rectangle, colWidths []float64, bWidth float64, row, col int) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (t *Table) valueCellLowerLeft(r *types.Rectangle, colWidths []float64, bWidth float64, row, col int) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (t *Table) headerCellLowerLeft(r *types.Rectangle, colWidths []float64, bWidth float64, col int) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (t *Table) prepareTextDescriptor() (model.TextDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(model.TextDescriptor), nil
}

func (t *Table) calcTextDescriptorPadding(td *model.TextDescriptor, p *Padding) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Table) renderValues(p *model.Page, pageNr int, fonts model.FontMap, colWidths []float64, td model.TextDescriptor, ll func(row, col int) (float64, float64)) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Table) renderHeader(p *model.Page, pageNr int, fonts model.FontMap, colWidths []float64, td model.TextDescriptor, ll func(col int) (float64, float64)) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Table) rectForFill(r *types.Rectangle, bWidth float64) *types.Rectangle {
	_ = "STUB: not implemented"
	return nil
}

func (t *Table) render(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}
