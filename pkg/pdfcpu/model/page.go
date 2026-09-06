package model

import (
	"bytes"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type Resource struct {
	ID     string
	IndRef *types.IndirectRef
}

type FontResource struct {
	Res       Resource
	Lang      string
	CIDSet    *types.IndirectRef
	FontFile  *types.IndirectRef
	ToUnicode *types.IndirectRef
	W         *types.IndirectRef
}

type FontMap map[string]FontResource

func (fm FontMap) EnsureKey(fontName string) string { _ = "STUB: not implemented"; return "" }

type ImageResource struct {
	Res    Resource
	Width  int
	Height int
}

type ImageMap map[string]ImageResource

type FieldAnnotation struct {
	Dict   types.Dict
	IndRef *types.IndirectRef
	Ind    int
	Field  bool
	Kids   types.Array
}

type Page struct {
	MediaBox   *types.Rectangle
	CropBox    *types.Rectangle
	Fm         FontMap
	Im         ImageMap
	Annots     []FieldAnnotation
	AnnotTabs  map[int]FieldAnnotation
	LinkAnnots []LinkAnnotation
	Buf        *bytes.Buffer
	Fields     types.Array
}

func NewPage(mediaBox, cropBox *types.Rectangle) Page { _ = "STUB: not implemented"; return *new(Page) }

func NewPageWithBg(mediaBox *types.Rectangle, c color.SimpleColor) Page {
	_ = "STUB: not implemented"
	return *new(Page)
}
