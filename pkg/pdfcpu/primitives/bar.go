package primitives

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type Bar struct {
	pdf     *PDF
	content *Content
	X, Y    float64
	Width   int
	Color   string `json:"col"`
	col     *color.SimpleColor
	Style   string
	style   types.LineJoinStyle
	Hide    bool
}

func (b *Bar) validate() error { _ = "STUB: not implemented"; return nil }

func (b *Bar) render(p *model.Page) error { _ = "STUB: not implemented"; return nil }
