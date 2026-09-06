package primitives

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type ComboBox struct {
	pdf             *PDF
	content         *Content
	Label           *TextFieldLabel
	ID              string
	Tip             string
	Default         string
	Value           string
	Options         []string
	Position        [2]float64 `json:"pos"`
	x, y            float64
	Width           float64
	Dx, Dy          float64
	BoundingBox     *types.Rectangle `json:"-"`
	Edit            bool
	Font            *FormFont
	fontID          string `json:"-"`
	Margin          *Margin
	Border          *Border
	BackgroundColor string             `json:"bgCol"`
	BgCol           *color.SimpleColor `json:"-"`
	Alignment       string             `json:"align"`
	HorAlign        types.HAlignment   `json:"-"`
	RTL             bool
	Tab             int
	Locked          bool
	Debug           bool
	Hide            bool
}

func (cb *ComboBox) SetFontID(s string) { _ = "STUB: not implemented"; return }

func (cb *ComboBox) validateID() error { _ = "STUB: not implemented"; return nil }

func (cb *ComboBox) validatePosition() error { _ = "STUB: not implemented"; return nil }

func (cb *ComboBox) validateWidth() error { _ = "STUB: not implemented"; return nil }

func (cb *ComboBox) validateOptionsValueAndDefault() error { _ = "STUB: not implemented"; return nil }

func (cb *ComboBox) validateFont() error { _ = "STUB: not implemented"; return nil }

func (cb *ComboBox) validateMargin() error { _ = "STUB: not implemented"; return nil }

func (cb *ComboBox) validateBorder() error { _ = "STUB: not implemented"; return nil }

func (cb *ComboBox) validateBackgroundColor() error { _ = "STUB: not implemented"; return nil }

func (cb *ComboBox) validateHorAlign() error { _ = "STUB: not implemented"; return nil }

func (cb *ComboBox) validateLabel() error { _ = "STUB: not implemented"; return nil }

func (cb *ComboBox) validateTab() error { _ = "STUB: not implemented"; return nil }

func (cb *ComboBox) validate() error { _ = "STUB: not implemented"; return nil }

func (cb *ComboBox) calcFontFromDA(ctx *model.Context, d types.Dict, da *string, fonts map[string]types.IndirectRef) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cb *ComboBox) calcFont() error { _ = "STUB: not implemented"; return nil }

func (cb *ComboBox) margin(name string) *Margin { _ = "STUB: not implemented"; return nil }

func (cb *ComboBox) calcMargin() (float64, float64, float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, nil
}

func (cb *ComboBox) labelPos(labelHeight, w, g float64) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (cb *ComboBox) renderN(xRefTable *model.XRefTable) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cb *ComboBox) calcBorder() (boWidth float64, boCol *color.SimpleColor) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cb *ComboBox) prepareFF() FieldFlags { _ = "STUB: not implemented"; return *new(FieldFlags) }

func (cb *ComboBox) handleBorderAndMK(d types.Dict) { _ = "STUB: not implemented"; return }

func (cb *ComboBox) prepareDict(fonts model.FontMap) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func (cb *ComboBox) bbox() *types.Rectangle { _ = "STUB: not implemented"; return nil }

func (cb *ComboBox) prepareRectLL(mTop, mRight, mBottom, mLeft float64) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (cb *ComboBox) prepLabel(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (cb *ComboBox) prepForRender(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (cb *ComboBox) doRender(p *model.Page, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (cb *ComboBox) render(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func NewComboBox(
	ctx *model.Context,
	d types.Dict,
	v string,
	da *string,
	fonts map[string]types.IndirectRef) (*ComboBox, *types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func renderComboBoxAP(ctx *model.Context, d types.Dict, v string, da *string, fonts map[string]types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func refreshComboBoxAP(ctx *model.Context, d types.Dict, v string, da *string, fonts map[string]types.IndirectRef, irN *types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func EnsureComboBoxAP(ctx *model.Context, d types.Dict, v string, da *string, fonts map[string]types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}
