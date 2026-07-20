package model

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type Box struct {
	Rect      *types.Rectangle `json:"rect"`
	Inherited bool             `json:"-"`
	RefBox    string           `json:"-"`

	MLeft, MRight float64 `json:"-"`
	MTop, MBot    float64 `json:"-"`

	Dim    *types.Dim   `json:"-"`
	Pos    types.Anchor `json:"-"`
	Dx, Dy int          `json:"-"`
}

type PageBoundaries struct {
	Media       *Box   `json:"mediaBox,omitempty"`
	Crop        *Box   `json:"cropBox,omitempty"`
	Trim        *Box   `json:"trimBox,omitempty"`
	Bleed       *Box   `json:"bleedBox,omitempty"`
	Art         *Box   `json:"artBox,omitempty"`
	Rot         int    `json:"rot"`
	Orientation string `json:"orient"`
}

func (pb *PageBoundaries) SelectAll() { _ = "STUB: not implemented"; return }

func (pb PageBoundaries) String() string { _ = "STUB: not implemented"; return "" }

func (pb PageBoundaries) MediaBox() *types.Rectangle { _ = "STUB: not implemented"; return nil }

func (pb PageBoundaries) CropBox() *types.Rectangle { _ = "STUB: not implemented"; return nil }

func (pb PageBoundaries) TrimBox() *types.Rectangle { _ = "STUB: not implemented"; return nil }

func (pb PageBoundaries) BleedBox() *types.Rectangle { _ = "STUB: not implemented"; return nil }

func (pb PageBoundaries) ArtBox() *types.Rectangle { _ = "STUB: not implemented"; return nil }

func (pb *PageBoundaries) ResolveBox(s string) error { _ = "STUB: not implemented"; return nil }

func ParseBoxList(s string) (*PageBoundaries, error) { _ = "STUB: not implemented"; return nil, nil }

func resolveBoxType(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func processBox(b **Box, boxID, paramValueStr string, unit types.DisplayUnit) error {
	_ = "STUB: not implemented"
	return nil
}

func ParsePageBoundaries(s string, unit types.DisplayUnit) (*PageBoundaries, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseBoxByRectangle(s string, u types.DisplayUnit) (*Box, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseBoxPercentage(s string) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func parseBoxBySingleMarginVal(s, s1 string, abs bool, u types.DisplayUnit) (*Box, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseBoxBy2Percentages(s, s1, s2 string) (*Box, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseBoxBy2MarginVals(s, s1, s2 string, abs bool, u types.DisplayUnit) (*Box, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseBoxBy3Percentages(s, s1, s2, s3 string) (*Box, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseBoxBy3MarginVals(s, s1, s2, s3 string, abs bool, u types.DisplayUnit) (*Box, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseBoxBy4Percentages(s, s1, s2, s3, s4 string) (*Box, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseBoxBy4MarginVals(s, s1, s2, s3, s4 string, abs bool, u types.DisplayUnit) (*Box, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseBoxOffset(s string, b *Box, u types.DisplayUnit) error {
	_ = "STUB: not implemented"
	return nil
}

func parseBoxDimByPercentage(s, s1, s2 string, b *Box) error { _ = "STUB: not implemented"; return nil }

func parseBoxDimWidthAndHeight(s1, s2 string, abs bool) (float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func parseBoxDim(s string, b *Box, u types.DisplayUnit) error {
	_ = "STUB: not implemented"
	return nil
}

func parseBoxByPosWithinParent(ss []string, u types.DisplayUnit) (*Box, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseBoxByMarginVals(ss []string, s string, abs bool, u types.DisplayUnit) (*Box, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseBox(s string, u types.DisplayUnit) (*Box, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctx *Context) addPageBoundaryString(i int, pb PageBoundaries, wantPB *PageBoundaries) []string {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *Context) ListPageBoundaries(selectedPages types.IntSet, wantPB *PageBoundaries) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctx *Context) RemovePageBoundaries(selectedPages types.IntSet, pb *PageBoundaries) error {
	_ = "STUB: not implemented"
	return nil
}

func boxLowerLeftCorner(r *types.Rectangle, w, h float64, a types.Anchor) types.Point {
	_ = "STUB: not implemented"
	return *new(types.Point)
}

func boxByDim(boxName string, b *Box, d types.Dict, parent *types.Rectangle) *types.Rectangle {
	_ = "STUB: not implemented"
	return nil
}

func ensureCropBoxWithinMediaBox(xmin, xmax, ymin, ymax float64, d types.Dict, parent *types.Rectangle) {
	_ = "STUB: not implemented"
	return
}

func ApplyBox(boxName string, b *Box, d types.Dict, parent *types.Rectangle) *types.Rectangle {
	_ = "STUB: not implemented"
	return nil
}

type boxes struct {
	mediaBox, cropBox, trimBox, bleedBox, artBox *types.Rectangle
}

func applyBoxDefinitions(d types.Dict, pb *PageBoundaries, b *boxes) {
	_ = "STUB: not implemented"
	return
}

func updateTrimBox(d types.Dict, trimBox *Box, b *boxes) { _ = "STUB: not implemented"; return }

func updateBleedBox(d types.Dict, bleedBox *Box, b *boxes) { _ = "STUB: not implemented"; return }

func updateArtBox(d types.Dict, artBox *Box, b *boxes) { _ = "STUB: not implemented"; return }

func applyBoxAssignments(d types.Dict, pb *PageBoundaries, b *boxes) {
	_ = "STUB: not implemented"
	return
}

func (ctx *Context) AddPageBoundaries(selectedPages types.IntSet, pb *PageBoundaries) error {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *Context) Crop(selectedPages types.IntSet, b *Box) error {
	_ = "STUB: not implemented"
	return nil
}
