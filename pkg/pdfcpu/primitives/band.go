package primitives

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type HorizontalBand struct {
	pdf             *PDF
	Left            string
	Center          string
	Right           string
	position        types.Anchor
	Height          float64
	Dx, Dy          int
	BackgroundColor string `json:"bgCol"`
	bgCol           *color.SimpleColor
	Font            *FormFont
	From            int
	Thru            int
	Border          bool
	RTL             bool
}

func (hb *HorizontalBand) validate() error { _ = "STUB: not implemented"; return nil }

func (hb *HorizontalBand) renderAnchoredImageBox(
	imageName string,
	r *types.Rectangle,
	a types.Anchor,
	p *model.Page,
	pageNr int,
	images model.ImageMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (hb *HorizontalBand) renderAnchoredTextBox(
	s string,
	r *types.Rectangle,
	a types.Anchor,
	p *model.Page,
	pageNr int,
	fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (hb *HorizontalBand) renderComponent(
	content string,
	a types.Anchor,
	r *types.Rectangle,
	p *model.Page,
	pageNr int,
	fonts model.FontMap,
	images model.ImageMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (hb *HorizontalBand) render(p *model.Page, pageNr int, fonts model.FontMap, images model.ImageMap, top bool) error {
	_ = "STUB: not implemented"
	return nil
}
