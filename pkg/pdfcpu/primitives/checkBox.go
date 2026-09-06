package primitives

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type CheckBox struct {
	pdf             *PDF
	content         *Content
	Label           *TextFieldLabel
	ID              string
	Tip             string
	Value           bool
	Default         bool
	Position        [2]float64 `json:"pos"`
	x, y            float64
	Width           float64
	Dx, Dy          float64
	boundingBox     *types.Rectangle
	Margin          *Margin
	BackgroundColor string `json:"bgCol"`
	bgCol           *color.SimpleColor
	Tab             int
	Locked          bool
	Debug           bool
	Hide            bool
}

type AP struct {
	irDOffL, irDYesL *types.IndirectRef
	irNOffL, irNYesL *types.IndirectRef
	irDOffR, irDYesR *types.IndirectRef
	irNOffR, irNYesR *types.IndirectRef
}

func (cb *CheckBox) validateID() error { _ = "STUB: not implemented"; return nil }

func (cb *CheckBox) validatePosition() error { _ = "STUB: not implemented"; return nil }

func (cb *CheckBox) validateMargin() error { _ = "STUB: not implemented"; return nil }

func (cb *CheckBox) validateWidth() error { _ = "STUB: not implemented"; return nil }

func (cb *CheckBox) validateLabel() error { _ = "STUB: not implemented"; return nil }

func (cb *CheckBox) validateTab() error { _ = "STUB: not implemented"; return nil }

func (cb *CheckBox) validate() error { _ = "STUB: not implemented"; return nil }

func (cb *CheckBox) margin(name string) *Margin { _ = "STUB: not implemented"; return nil }

func (cb *CheckBox) calcMargin() (float64, float64, float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, nil
}

func (cb *CheckBox) labelPos(labelHeight, w, g float64) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (cb *CheckBox) ensureZapfDingbats(fonts model.FontMap) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cb *CheckBox) calcFont() error { _ = "STUB: not implemented"; return nil }

func (cb *CheckBox) irNOff(bgCol *color.SimpleColor) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cb *CheckBox) irNYes(fonts model.FontMap, bgCol *color.SimpleColor) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cb *CheckBox) irDOff(bgCol *color.SimpleColor) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cb *CheckBox) irDYes(fonts model.FontMap, bgCol *color.SimpleColor) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cb *CheckBox) appearanceIndRefs(fonts model.FontMap, bgCol *color.SimpleColor) (
	*types.IndirectRef, *types.IndirectRef, *types.IndirectRef, *types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil
}

func (cb *CheckBox) prepareDict(fonts model.FontMap) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func (cb *CheckBox) bbox() *types.Rectangle { _ = "STUB: not implemented"; return nil }

func (cb *CheckBox) prepareRectLL(mTop, mRight, mBottom, mLeft float64) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (cb *CheckBox) prepLabel(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (cb *CheckBox) prepForRender(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (cb *CheckBox) doRender(p *model.Page, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (cb *CheckBox) render(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func CalcCheckBoxASNames(ctx *model.Context, d types.Dict) (types.Name, types.Name, error) {
	_ = "STUB: not implemented"
	return *new(types.Name), *new(types.Name), nil
}
