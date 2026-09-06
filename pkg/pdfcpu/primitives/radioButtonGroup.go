package primitives

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type RadioButtonGroup struct {
	pdf             *PDF
	content         *Content
	Label           *TextFieldLabel
	ID              string
	Tip             string
	Value           string
	Default         string
	Position        [2]float64 `json:"pos"`
	x, y            float64
	Width           float64
	boundingBox     *types.Rectangle
	Orientation     string
	hor             bool
	Dx, Dy          float64
	Margin          *Margin
	BackgroundColor string `json:"bgCol"`
	bgCol           *color.SimpleColor
	Buttons         *Buttons
	RTL             bool
	Tab             int
	Locked          bool
	Debug           bool
	Hide            bool
}

func (rbg *RadioButtonGroup) Rtl() bool { _ = "STUB: not implemented"; return false }

func (rbg *RadioButtonGroup) validateID() error { _ = "STUB: not implemented"; return nil }

func (rbg *RadioButtonGroup) validatePosition() error { _ = "STUB: not implemented"; return nil }

func parseRadioButtonOrientation(s string) (types.Orientation, error) {
	_ = "STUB: not implemented"
	return *new(types.Orientation), nil
}

func (rbg *RadioButtonGroup) validateOrientation() error { _ = "STUB: not implemented"; return nil }

func (rbg *RadioButtonGroup) validateWidth() error { _ = "STUB: not implemented"; return nil }

func (rbg *RadioButtonGroup) validateMargin() error { _ = "STUB: not implemented"; return nil }

func (rbg *RadioButtonGroup) validateLabel() error { _ = "STUB: not implemented"; return nil }

func (rbg *RadioButtonGroup) validateButtonsDefaultAndValue() error {
	_ = "STUB: not implemented"
	return nil
}

func (rbg *RadioButtonGroup) validateTab() error { _ = "STUB: not implemented"; return nil }

func (rbg *RadioButtonGroup) validate() error { _ = "STUB: not implemented"; return nil }

func (rbg *RadioButtonGroup) calcFont() error { _ = "STUB: not implemented"; return nil }

func (rbg *RadioButtonGroup) margin(name string) *Margin { _ = "STUB: not implemented"; return nil }

func (rbg *RadioButtonGroup) prepareMargin() (float64, float64, float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, nil
}

func (rbg *RadioButtonGroup) buttonLabelPosition(i int) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (rbg *RadioButtonGroup) renderButtonLabels(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (rbg *RadioButtonGroup) buttonGroupBB() *types.Rectangle {
	_ = "STUB: not implemented"
	return nil
}

func labelPos(
	relPos types.RelPosition,
	horAlign types.HAlignment,
	boundingBox *types.Rectangle,
	labelHeight, w, g float64, multiline bool) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (rbg *RadioButtonGroup) rect(i int) *types.Rectangle { _ = "STUB: not implemented"; return nil }

func (rbg *RadioButtonGroup) irDOff(asWidth float64, flip bool) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rbg *RadioButtonGroup) irDYes(asWidth float64, flip bool) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rbg *RadioButtonGroup) irNOff(asWidth float64, flip bool, bgCol *color.SimpleColor) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rbg *RadioButtonGroup) irNYes(asWidth float64, flip bool, bgCol *color.SimpleColor) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rbg *RadioButtonGroup) appearanceIndRefs(flip bool, bgCol *color.SimpleColor) (
	*types.IndirectRef, *types.IndirectRef, *types.IndirectRef, *types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil
}

func (rbg *RadioButtonGroup) prepareButtonDict(r *types.Rectangle, v string, parent types.IndirectRef, irDOff, irDYes, irNOff, irNYes *types.IndirectRef) (*types.IndirectRef, types.Dict, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.Dict), nil
}

func (rbg *RadioButtonGroup) renderRadioButtonFields(p *model.Page, parent types.IndirectRef) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func (rbg *RadioButtonGroup) bbox() *types.Rectangle { _ = "STUB: not implemented"; return nil }

func (rbg *RadioButtonGroup) prepareRectLL(mTop, mRight, mBottom, mLeft float64) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (rbg *RadioButtonGroup) prepLabel(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (rbg *RadioButtonGroup) prepForRender(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (rbg *RadioButtonGroup) prepareDict(p *model.Page, pageNr int, fonts model.FontMap) (*types.IndirectRef, types.Array, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.Array), nil
}

func (rbg *RadioButtonGroup) doRender(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (rbg *RadioButtonGroup) render(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}
