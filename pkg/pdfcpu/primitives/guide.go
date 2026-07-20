package primitives

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type Guide struct {
	Position [2]float64 `json:"pos"`
	x, y     float64
}

func (g *Guide) validate() { _ = "STUB: not implemented"; return }

func (g *Guide) render(w io.Writer, r *types.Rectangle, pdf *PDF) {
	_ = "STUB: not implemented"
	return
}
