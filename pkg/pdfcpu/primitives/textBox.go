package primitives

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type TextBox struct {
	pdf      *PDF
	content  *Content
	Name     string
	Value    string
	Position [2]float64 `json:"pos"`
	x, y     float64
	Dx, Dy   float64
	Anchor   string
	anchor   types.Anchor
	anchored bool
	Width    float64

	Font    *FormFont
	Margin  *Margin
	Border  *Border
	Padding *Padding

	BackgroundColor string `json:"bgCol"`
	bgCol           *color.SimpleColor
	Alignment       string `json:"align"`
	horAlign        types.HAlignment
	RTL             bool
	Rotation        float64 `json:"rot"`
	Hide            bool
}

func (tb *TextBox) validateAnchor() error { _ = "STUB: not implemented"; return nil }

func (tb *TextBox) validateFont() error { _ = "STUB: not implemented"; return nil }

func (tb *TextBox) validateMargin() error { _ = "STUB: not implemented"; return nil }

func (tb *TextBox) validateBorder() error { _ = "STUB: not implemented"; return nil }

func (tb *TextBox) validatePadding() error { _ = "STUB: not implemented"; return nil }

func (tb *TextBox) validateBackgroundColor() error { _ = "STUB: not implemented"; return nil }

func (tb *TextBox) validateHorAlign() error { _ = "STUB: not implemented"; return nil }

func (tb *TextBox) validate() error { _ = "STUB: not implemented"; return nil }

func (tb *TextBox) font(name string) *FormFont { _ = "STUB: not implemented"; return nil }

func (tb *TextBox) margin(name string) *Margin { _ = "STUB: not implemented"; return nil }

func (tb *TextBox) border(name string) *Border { _ = "STUB: not implemented"; return nil }

func (tb *TextBox) padding(name string) *Padding { _ = "STUB: not implemented"; return nil }

func (tb *TextBox) mergeInPos(tb0 *TextBox) { _ = "STUB: not implemented"; return }

func (tb *TextBox) mergeIn(tb0 *TextBox) { _ = "STUB: not implemented"; return }

func (tb *TextBox) calcFont() error { _ = "STUB: not implemented"; return nil }

func tdMargin(p *Padding, td *model.TextDescriptor) { _ = "STUB: not implemented"; return }

func (tb *TextBox) prepareTextDescriptor(p *model.Page, pageNr int, fonts model.FontMap) (*model.TextDescriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tb *TextBox) calcMargin() (float64, float64, float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, nil
}

func (tb *TextBox) render(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}
