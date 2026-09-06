package primitives

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/matrix"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type SimpleBox struct {
	pdf       *PDF
	content   *Content
	Name      string
	Position  [2]float64 `json:"pos"`
	x, y      float64
	Dx, Dy    float64
	Anchor    string
	anchor    types.Anchor
	anchored  bool
	Width     float64
	Height    float64
	Margin    *Margin
	Border    *Border
	FillColor string `json:"fillCol"`
	fillCol   *color.SimpleColor
	Rotation  float64 `json:"rot"`
	Hide      bool
}

func (sb *SimpleBox) validate() error { _ = "STUB: not implemented"; return nil }

func (sb *SimpleBox) margin(name string) *Margin { _ = "STUB: not implemented"; return nil }

func (sb *SimpleBox) border(name string) *Border { _ = "STUB: not implemented"; return nil }

func (sb *SimpleBox) mergeIn(sb0 *SimpleBox) { _ = "STUB: not implemented"; return }

func (sb *SimpleBox) calcBorder() (float64, *color.SimpleColor, types.LineJoinStyle, error) {
	_ = "STUB: not implemented"
	return 0, nil, *new(types.LineJoinStyle), nil
}

func (sb *SimpleBox) calcMargin() (float64, float64, float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, nil
}

func (sb *SimpleBox) calcTransform(mLeft, mBottom, mRight, mTop, bWidth float64) (matrix.Matrix, *types.Rectangle) {
	_ = "STUB: not implemented"
	return *new(matrix.Matrix), nil
}

func (sb *SimpleBox) render(p *model.Page) error { _ = "STUB: not implemented"; return nil }
