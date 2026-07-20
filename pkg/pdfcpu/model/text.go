package model

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/draw"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type TextDescriptor struct {
	Text           string
	FontName       string
	RTL            bool
	Embed          bool
	FontKey        string
	FontSize       int
	X, Y           float64
	Dx, Dy         float64
	MTop, MBot     float64
	MLeft, MRight  float64
	MinHeight      float64
	Rotation       float64
	ScaleAbs       bool
	Scale          float64
	HAlign         types.HAlignment
	VAlign         types.VAlignment
	RMode          draw.RenderMode
	StrokeCol      color.SimpleColor
	FillCol        color.SimpleColor
	ShowTextBB     bool
	ShowBackground bool
	BackgroundCol  color.SimpleColor
	ShowBorder     bool
	BorderWidth    float64
	BorderStyle    types.LineJoinStyle
	BorderCol      color.SimpleColor
	ParIndent      bool
	ShowLineBB     bool
	ShowMargins    bool
	ShowPosition   bool
	HairCross      bool
}

func deltaAlignMiddle(fontName string, fontSize, lines int, mTop, mBot float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func deltaAlignTop(fontName string, fontSize int, mTop float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func deltaAlignBottom(fontName string, fontSize, lines int, mBot float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

var unicodeToCP1252 = map[rune]byte{
	0x20AC: 128,
	0x201A: 130,
	0x0192: 131,
	0x201E: 132,
	0x2026: 133,
	0x2020: 134,
	0x2021: 135,
	0x02C6: 136,
	0x2030: 137,
	0x0160: 138,
	0x2039: 139,
	0x0152: 140,
	0x017D: 142,
	0x2018: 145,
	0x2019: 146,
	0x201C: 147,
	0x201D: 148,
	0x2022: 149,
	0x2013: 150,
	0x2014: 151,
	0x02DC: 152,
	0x2122: 153,
	0x0161: 154,
	0x203A: 155,
	0x0153: 156,
	0x017E: 158,
	0x0178: 159,
}

func DecodeUTF8ToByte(s string) string { _ = "STUB: not implemented"; return "" }

func calcBoundingBoxForRectAndPoint(r *types.Rectangle, p types.Point) *types.Rectangle {
	_ = "STUB: not implemented"
	return nil
}

func CalcBoundingBoxForRects(r1, r2 *types.Rectangle) *types.Rectangle {
	_ = "STUB: not implemented"
	return nil
}

func calcBoundingBoxForLines(lines []string, x, y float64, fontName string, fontSize int) (*types.Rectangle, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

func PrepBytes(xRefTable *XRefTable, s, fontName string, embed, rtl, fillFont bool) string {
	_ = "STUB: not implemented"
	return ""
}

func writeStringToBuf(xRefTable *XRefTable, w io.Writer, s string, x, y float64, td TextDescriptor) {
	_ = "STUB: not implemented"
	return
}

func setFont(w io.Writer, fontID string, fontSize float32) { _ = "STUB: not implemented"; return }

func CalcBoundingBox(s string, x, y float64, fontName string, fontSize int) *types.Rectangle {
	_ = "STUB: not implemented"
	return nil
}

func horAdjustBoundingBoxForLines(r, box *types.Rectangle, dx, dy float64, x, y *float64) {
	_ = "STUB: not implemented"
	return
}

func prepJustifiedLine(xRefTable *XRefTable, lines *[]string, strbuf []string, strWidth, w float64, fontSize int, fontName string, embed, rtl bool) {
	_ = "STUB: not implemented"
	return
}

func newPrepJustifiedString(
	xRefTable *XRefTable,
	fontName string,
	fontSize int) func(lines *[]string, s string, w float64, fontName string, fontSize *int, lastline, parIndent, cjk, rtl bool) int {
	_ = "STUB: not implemented"
	return nil
}

func preRenderJustifiedText(
	xRefTable *XRefTable,
	lines *[]string,
	r *types.Rectangle,
	x, y, width float64,
	td TextDescriptor,
	mLeft, mRight, borderWidth float64,
	fontSize *int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func scaleFontSize(r *types.Rectangle, lines []string, scaleAbs bool,
	scale, width, x, y, mLeft, mRight, borderWidth float64,
	fontName string, fontSize *int) {
	_ = "STUB: not implemented"
	return
}

func horizontalWrapUp(box *types.Rectangle, maxLine string, hAlign types.HAlignment,
	x *float64, width, ww, mLeft, mRight, borderWidth float64,
	fontName string, fontSize *int) {
	_ = "STUB: not implemented"
	return
}

func createBoundingBoxForColumn(xRefTable *XRefTable, r *types.Rectangle, x, y *float64,
	width float64,
	td TextDescriptor,
	dx, dy float64,
	mTop, mBot, mLeft, mRight float64,
	borderWidth float64,
	fontSize *int, lines *[]string) *types.Rectangle {
	_ = "STUB: not implemented"
	return nil
}

func flushJustifiedStringToBuf(w io.Writer, s string, x, y float64, strokeCol, fillCol color.SimpleColor, rm draw.RenderMode) {
	_ = "STUB: not implemented"
	return
}

func scaleXForRegion(x float64, mediaBox, region *types.Rectangle) float64 {
	_ = "STUB: not implemented"
	return 0
}

func scaleYForRegion(y float64, mediaBox, region *types.Rectangle) float64 {
	_ = "STUB: not implemented"
	return 0
}

func DrawMargins(w io.Writer, c color.SimpleColor, colBB *types.Rectangle, borderWidth, mLeft, mRight, mTop, mBot float64) {
	_ = "STUB: not implemented"
	return
}

func renderBackgroundAndBorder(w io.Writer, td TextDescriptor, borderWidth float64, colBB *types.Rectangle) {
	_ = "STUB: not implemented"
	return
}

func renderText(xRefTable *XRefTable, w io.Writer, lines []string, td TextDescriptor, x, y float64, fontSize int) {
	_ = "STUB: not implemented"
	return
}

func fieldsFunc(s string, f func(rune) bool) []string { _ = "STUB: not implemented"; return nil }

func SplitMultilineStr(s string) []string { _ = "STUB: not implemented"; return nil }

func wrapLine(ss *[]string, line, space, word, fontName string, fontSize int, maxWidthPoints float64) {
	_ = "STUB: not implemented"
	return
}

func wrap(lines []string, fontName string, fontSize int, maxWidthPoints float64) []string {
	_ = "STUB: not implemented"
	return nil
}

func WordWrap(s string, fontName string, fontSize int, maxWidthPoints float64) []string {
	_ = "STUB: not implemented"
	return nil
}

func WriteColumn(xRefTable *XRefTable, w io.Writer, mediaBox, region *types.Rectangle, td TextDescriptor, width float64) *types.Rectangle {
	_ = "STUB: not implemented"
	return nil
}

func WriteMultiLine(xRefTable *XRefTable, w io.Writer, mediaBox, region *types.Rectangle, td TextDescriptor) *types.Rectangle {
	_ = "STUB: not implemented"
	return nil
}

func AnchorPosAndAlign(a types.Anchor, r *types.Rectangle) (x, y float64, hAlign types.HAlignment, vAlign types.VAlignment) {
	_ = "STUB: not implemented"
	return 0, 0, *new(types.HAlignment), *new(types.VAlignment)
}

func WriteMultiLineAnchored(xRefTable *XRefTable, w io.Writer, mediaBox, region *types.Rectangle, td TextDescriptor, a types.Anchor) *types.Rectangle {
	_ = "STUB: not implemented"
	return nil
}

func WriteColumnAnchored(xRefTable *XRefTable, w io.Writer, mediaBox, region *types.Rectangle, td TextDescriptor, a types.Anchor, width float64) *types.Rectangle {
	_ = "STUB: not implemented"
	return nil
}
