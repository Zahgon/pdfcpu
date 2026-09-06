package primitives

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type Divider struct {
	pdf   *PDF
	Pos   float64 `json:"at"`
	p, q  types.Point
	Width int
	Color string `json:"col"`
	col   *color.SimpleColor
}

func (d *Divider) validate() error { _ = "STUB: not implemented"; return nil }

func (d *Divider) render(p *model.Page) error { _ = "STUB: not implemented"; return nil }
