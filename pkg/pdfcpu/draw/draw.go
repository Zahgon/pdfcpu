package draw

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type RenderMode int

const (
	RMFill RenderMode = iota
	RMStroke
	RMFillAndStroke
)

func SetLineJoinStyle(w io.Writer, s types.LineJoinStyle) { _ = "STUB: not implemented"; return }

func SetLineWidth(w io.Writer, width float64) { _ = "STUB: not implemented"; return }

func SetStrokeColor(w io.Writer, c color.SimpleColor) { _ = "STUB: not implemented"; return }

func SetFillColor(w io.Writer, c color.SimpleColor) { _ = "STUB: not implemented"; return }

func DrawLineSimple(w io.Writer, xp, yp, xq, yq float64) { _ = "STUB: not implemented"; return }

func DrawLine(w io.Writer, xp, yp, xq, yq float64, lineWidth float64, strokeColor *color.SimpleColor, style *types.LineJoinStyle) {
	_ = "STUB: not implemented"
	return
}

func DrawRectSimple(w io.Writer, r *types.Rectangle) { _ = "STUB: not implemented"; return }

func DrawRect(w io.Writer, r *types.Rectangle, lineWidth float64, strokeColor *color.SimpleColor, style *types.LineJoinStyle) {
	_ = "STUB: not implemented"
	return
}

func FillRect(w io.Writer, r *types.Rectangle, lineWidth float64, strokeCol *color.SimpleColor, fillCol color.SimpleColor, style *types.LineJoinStyle) {
	_ = "STUB: not implemented"
	return
}

func DrawCircle(w io.Writer, x, y, r float64, strokeCol color.SimpleColor, fillCol *color.SimpleColor) {
	_ = "STUB: not implemented"
	return
}

func FillRectNoBorder(w io.Writer, r *types.Rectangle, fillCol color.SimpleColor) {
	_ = "STUB: not implemented"
	return
}

func DrawGrid(w io.Writer, x, y int, r *types.Rectangle, strokeCol color.SimpleColor, fillCol *color.SimpleColor) {
	_ = "STUB: not implemented"
	return
}

func DrawHairCross(w io.Writer, x, y float64, r *types.Rectangle) {
	_ = "STUB: not implemented"
	return
}

const (
	HBar     = "\u2501"
	VBar     = "\u2502"
	CrossBar = "\u253f"
)

func HorSepLine(ii []int) string { _ = "STUB: not implemented"; return "" }
