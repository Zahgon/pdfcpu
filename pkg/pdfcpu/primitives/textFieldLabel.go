package primitives

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type TextFieldLabel struct {
	TextField
	Width    int
	height   float64
	Gap      int
	Position string `json:"pos"`
	relPos   types.RelPosition
	td       *model.TextDescriptor
}

func (tfl *TextFieldLabel) validate() error { _ = "STUB: not implemented"; return nil }
