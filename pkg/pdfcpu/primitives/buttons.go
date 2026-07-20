package primitives

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type Buttons struct {
	pdf         *PDF
	Values      []string
	Label       *TextFieldLabel
	Gap         int
	widths      []float64
	maxWidth    float64
	boundingBox *types.Rectangle
}

func (b *Buttons) Rtl() bool { _ = "STUB: not implemented"; return false }

func (b *Buttons) validate(defValue, value string) error { _ = "STUB: not implemented"; return nil }

func (b *Buttons) calcLeftAlignedHorLabelWidths(td model.TextDescriptor) {
	_ = "STUB: not implemented"
	return
}

func (b *Buttons) calcRightAlignedHorLabelWidths(td model.TextDescriptor) {
	_ = "STUB: not implemented"
	return
}

func (b *Buttons) calcHorLabelWidths(td model.TextDescriptor) { _ = "STUB: not implemented"; return }

func (b *Buttons) calcVerLabelWidths(td model.TextDescriptor) { _ = "STUB: not implemented"; return }

func (b *Buttons) calcLabelWidths(hor bool) { _ = "STUB: not implemented"; return }
