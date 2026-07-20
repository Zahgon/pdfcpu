package primitives

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type ListBox struct {
	pdf             *PDF
	content         *Content
	Label           *TextFieldLabel
	ID              string
	Tip             string
	Default         string
	Defaults        []string
	Value           string
	Values          []string
	Ind             types.Array `json:"-"`
	Options         []string
	Position        [2]float64 `json:"pos"`
	x, y            float64
	Width           float64
	Height          float64
	Dx, Dy          float64
	BoundingBox     *types.Rectangle `json:"-"`
	Multi           bool             `json:"multi"`
	Font            *FormFont
	fontID          string
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

func (lb *ListBox) SetFontID(s string) { _ = "STUB: not implemented"; return }

func (lb *ListBox) validateDefault() error { _ = "STUB: not implemented"; return nil }

func (lb *ListBox) validateValue() error { _ = "STUB: not implemented"; return nil }

func (lb *ListBox) validateID() error { _ = "STUB: not implemented"; return nil }

func (lb *ListBox) validatePosition() error { _ = "STUB: not implemented"; return nil }

func (lb *ListBox) validateWidth() error { _ = "STUB: not implemented"; return nil }

func (lb *ListBox) validateHeight() error { _ = "STUB: not implemented"; return nil }

func (lb *ListBox) validateFont() error { _ = "STUB: not implemented"; return nil }

func (lb *ListBox) validateMargin() error { _ = "STUB: not implemented"; return nil }

func (lb *ListBox) validateBorder() error { _ = "STUB: not implemented"; return nil }

func (lb *ListBox) validateBackgroundColor() error { _ = "STUB: not implemented"; return nil }

func (lb *ListBox) validateHorAlign() error { _ = "STUB: not implemented"; return nil }

func (lb *ListBox) validateLabel() error { _ = "STUB: not implemented"; return nil }

func (lb *ListBox) validateTab() error { _ = "STUB: not implemented"; return nil }

func (lb *ListBox) validate() error { _ = "STUB: not implemented"; return nil }

func (lb *ListBox) calcFontFromDA(ctx *model.Context, d types.Dict, da *string, fonts map[string]types.IndirectRef) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lb *ListBox) calcFont() error { _ = "STUB: not implemented"; return nil }

func (lb *ListBox) margin(name string) *Margin { _ = "STUB: not implemented"; return nil }

func (lb *ListBox) calcMargin() (float64, float64, float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, nil
}

func (lb *ListBox) labelPos(labelHeight, w, g float64) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func selectItem(w io.Writer, i int, width, height float64, fontName string, fontSize int, boWidth float64, col color.SimpleColor) {
	_ = "STUB: not implemented"
	return
}

func (lb *ListBox) renderN(xRefTable *model.XRefTable) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lb *ListBox) irN(fonts model.FontMap) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lb *ListBox) calcBorder() (boWidth float64, boCol *color.SimpleColor) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (lb *ListBox) prepareFF() FieldFlags { _ = "STUB: not implemented"; return *new(FieldFlags) }

func (lb *ListBox) handleBorderAndMK(d types.Dict) { _ = "STUB: not implemented"; return }

func (lb *ListBox) handleVAndDV(d types.Dict) error { _ = "STUB: not implemented"; return nil }

func (lb *ListBox) prepareDict(fonts model.FontMap) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func (lb *ListBox) bbox() *types.Rectangle { _ = "STUB: not implemented"; return nil }

func (lb *ListBox) prepareRectLL(mTop, mRight, mBottom, mLeft float64) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (lb *ListBox) prepLabel(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (lb *ListBox) prepForRender(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (lb *ListBox) doRender(p *model.Page, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (lb *ListBox) render(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func NewListBox(
	ctx *model.Context,
	d types.Dict,
	opts []string,
	ind types.Array,
	da *string,
	fonts map[string]types.IndirectRef) (*ListBox, *types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func NewForm(
	xRefTable *model.XRefTable,
	bb []byte,
	fontID string,
	fontIndRef *types.IndirectRef,
	boundingBox *types.Rectangle) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func updateForm(xRefTable *model.XRefTable, bb []byte, indRef *types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func renderListBoxAP(ctx *model.Context, d types.Dict, opts []string, ind types.Array, da *string, fonts map[string]types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func refreshListBoxAP(ctx *model.Context, d types.Dict, opts []string, ind types.Array, da *string, fonts map[string]types.IndirectRef, irN *types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func EnsureListBoxAP(ctx *model.Context, d types.Dict, opts []string, ind types.Array, da *string, fonts map[string]types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}
