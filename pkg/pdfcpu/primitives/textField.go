package primitives

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type TextField struct {
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
	Height          float64
	Dx, Dy          float64
	BoundingBox     *types.Rectangle `json:"-"`
	Multiline       bool
	Font            *FormFont
	fontID          string
	Margin          *Margin
	Border          *Border
	BackgroundColor string             `json:"bgCol"`
	BgCol           *color.SimpleColor `json:"-"`
	Alignment       string             `json:"align"`
	HorAlign        types.HAlignment   `json:"-"`
	MaxLen          int                `json:"maxlen"`
	Comb            bool               `json:"comb"`
	RTL             bool
	Tab             int
	Locked          bool
	Debug           bool
	Hide            bool
}

func (tf *TextField) SetFontID(s string) { _ = "STUB: not implemented"; return }

func (tf *TextField) validateID() error { _ = "STUB: not implemented"; return nil }

func (tf *TextField) validatePosition() error { _ = "STUB: not implemented"; return nil }

func (tf *TextField) validateWidth() error { _ = "STUB: not implemented"; return nil }

func (tf *TextField) validateHeight() error { _ = "STUB: not implemented"; return nil }

func (tf *TextField) validateFont() error { _ = "STUB: not implemented"; return nil }

func (tf *TextField) validateMargin() error { _ = "STUB: not implemented"; return nil }

func (tf *TextField) validateBorder() error { _ = "STUB: not implemented"; return nil }

func (tf *TextField) validateBackgroundColor() error { _ = "STUB: not implemented"; return nil }

func (tf *TextField) validateHorAlign() error { _ = "STUB: not implemented"; return nil }

func (tf *TextField) validateLabel() error { _ = "STUB: not implemented"; return nil }

func (tf *TextField) validateTab() error { _ = "STUB: not implemented"; return nil }

func (tf *TextField) validate() error { _ = "STUB: not implemented"; return nil }

func locateDA(ctx *model.Context, d types.Dict, inhDA *string) *string {
	_ = "STUB: not implemented"
	return nil
}

func (tf *TextField) calcFontFromDA(ctx *model.Context, d types.Dict, da *string, needUTF8 bool, fonts map[string]types.IndirectRef) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tf *TextField) calcFont() error { _ = "STUB: not implemented"; return nil }

func (tf *TextField) margin(name string) *Margin { _ = "STUB: not implemented"; return nil }

func (tf *TextField) calcMargin() (float64, float64, float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, nil
}

func (tf *TextField) labelPos(labelHeight, w, g float64) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (tf *TextField) renderBackground(w io.Writer, bgCol, boCol *color.SimpleColor, boWidth, width, height float64) {
	_ = "STUB: not implemented"
	return
}

func (tf *TextField) renderLines(xRefTable *model.XRefTable, boWidth, lh, w, y float64, lines []string, buf io.Writer) {
	_ = "STUB: not implemented"
	return
}

func (tf *TextField) renderN(xRefTable *model.XRefTable) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tf *TextField) RefreshN(xRefTable *model.XRefTable, indRef *types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func (tf *TextField) irN(fonts model.FontMap) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tf *TextField) calcBorder() (boWidth float64, boCol *color.SimpleColor) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (tf *TextField) prepareFF() FieldFlags { _ = "STUB: not implemented"; return *new(FieldFlags) }

func (tf *TextField) handleBorderAndMK(d types.Dict) { _ = "STUB: not implemented"; return }

func (tf *TextField) prepareDict(fonts model.FontMap) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func (tf *TextField) bbox() *types.Rectangle { _ = "STUB: not implemented"; return nil }

func (tf *TextField) prepareRectLL(mTop, mRight, mBottom, mLeft float64) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (tf *TextField) prepLabel(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (tf *TextField) prepForRender(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (tf *TextField) doRender(p *model.Page, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (tf *TextField) render(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func calcColsFromMK(ctx *model.Context, d types.Dict) (*color.SimpleColor, *color.SimpleColor, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func calcBorderWidth(d types.Dict) int { _ = "STUB: not implemented"; return 0 }

func hasUTF(s string) bool { _ = "STUB: not implemented"; return false }

func NewTextField(
	ctx *model.Context,
	d types.Dict,
	v string,
	multiLine bool,
	comb bool,
	maxLen int,
	da *string,
	fontIndRef *types.IndirectRef,
	fonts map[string]types.IndirectRef) (*TextField, *types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func renderTextFieldAP(ctx *model.Context, d types.Dict, v string, multiLine, comb bool, maxLen int, da *string, fonts map[string]types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func fontAttrs(ctx *model.Context, fd types.Dict, fontID, text string, fonts map[string]types.IndirectRef) (string, string, string, string, *types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return "", "", "", "", nil, nil
}

func EnsureTextFieldAP(ctx *model.Context, d types.Dict, text string, multiLine, comb bool, maxLen int, da *string, fonts map[string]types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}
