package primitives

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type DateField struct {
	pdf             *PDF
	content         *Content
	Label           *TextFieldLabel
	ID              string
	Tip             string
	Value           string
	Default         string
	DateFormat      string `json:"format"`
	dateFormat      *DateFormat
	Position        [2]float64 `json:"pos"`
	x, y            float64
	Width           float64
	Dx, Dy          float64
	BoundingBox     *types.Rectangle `json:"-"`
	Font            *FormFont
	fontID          string
	Margin          *Margin
	Border          *Border
	BackgroundColor string             `json:"bgCol"`
	BgCol           *color.SimpleColor `json:"-"`
	Alignment       string             `json:"align"`
	HorAlign        types.HAlignment   `json:"-"`
	Tab             int
	Locked          bool
	Debug           bool
	Hide            bool
}

func (df *DateField) SetFontID(s string) { _ = "STUB: not implemented"; return }

func (df *DateField) validateID() error { _ = "STUB: not implemented"; return nil }

func (df *DateField) validatePosition() error { _ = "STUB: not implemented"; return nil }

func (df *DateField) validateWidth() error { _ = "STUB: not implemented"; return nil }

func (df *DateField) validateFont() error { _ = "STUB: not implemented"; return nil }

func (df *DateField) validateMargin() error { _ = "STUB: not implemented"; return nil }

func (df *DateField) validateBorder() error { _ = "STUB: not implemented"; return nil }

func (df *DateField) validateBackgroundColor() error { _ = "STUB: not implemented"; return nil }

func (df *DateField) validateHorAlign() error { _ = "STUB: not implemented"; return nil }

func (df *DateField) validateLabel() error { _ = "STUB: not implemented"; return nil }

func (df *DateField) validateDateFormat() error { _ = "STUB: not implemented"; return nil }

func (df *DateField) validateDefault() error { _ = "STUB: not implemented"; return nil }

func (df *DateField) validateValue() error { _ = "STUB: not implemented"; return nil }

func (df *DateField) validateTab() error { _ = "STUB: not implemented"; return nil }

func (df *DateField) validate() error { _ = "STUB: not implemented"; return nil }

func (df *DateField) calcFontFromDA(ctx *model.Context, d types.Dict, da *string, needUTF8 bool, fonts map[string]types.IndirectRef) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (df *DateField) calcFont() error { _ = "STUB: not implemented"; return nil }

func (df *DateField) margin(name string) *Margin { _ = "STUB: not implemented"; return nil }

func (df *DateField) calcMargin() (float64, float64, float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, nil
}

func (df *DateField) labelPos(labelHeight, w, g float64) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (tf *DateField) renderBackground(w io.Writer, bgCol, boCol *color.SimpleColor, boWidth, width, height float64) {
	_ = "STUB: not implemented"
	return
}

func (df *DateField) renderN(xRefTable *model.XRefTable) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (df *DateField) RefreshN(xRefTable *model.XRefTable, indRef *types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func (df *DateField) irN(fonts model.FontMap) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (df *DateField) calcBorder() (boWidth float64, boCol *color.SimpleColor) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (df *DateField) prepareFF() FieldFlags { _ = "STUB: not implemented"; return *new(FieldFlags) }

func (df *DateField) handleBorderAndMK(d types.Dict) { _ = "STUB: not implemented"; return }

func (df *DateField) prepareDict(fonts model.FontMap) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func (df *DateField) bbox() *types.Rectangle { _ = "STUB: not implemented"; return nil }

func (df *DateField) prepareRectLL(mTop, mRight, mBottom, mLeft float64) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (df *DateField) prepLabel(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (df *DateField) prepForRender(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (df *DateField) doRender(p *model.Page, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (df *DateField) render(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDateField(
	ctx *model.Context,
	d types.Dict,
	v string,
	da *string,
	fontIndRef *types.IndirectRef,
	fonts map[string]types.IndirectRef) (*DateField, *types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func renderDateFieldAP(ctx *model.Context, d types.Dict, v string, da *string, fonts map[string]types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func refreshDateFieldAP(ctx *model.Context, d types.Dict, v string, da *string, fonts map[string]types.IndirectRef, irN *types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func EnsureDateFieldAP(ctx *model.Context, d types.Dict, v string, da *string, fonts map[string]types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}
