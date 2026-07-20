package primitives

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type Border struct {
	pdf   *PDF
	Name  string
	Width int
	Color string `json:"col"`
	col   *color.SimpleColor
	Style string
	style types.LineJoinStyle
}

func (b *Border) validate() error { _ = "STUB: not implemented"; return nil }

func (b *Border) mergeIn(b0 *Border) { _ = "STUB: not implemented"; return }

func (b Border) calc() (boWidth float64, boCol *color.SimpleColor) {
	_ = "STUB: not implemented"
	return 0, nil
}
