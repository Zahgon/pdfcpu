package model

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/draw"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type textRenderWriter struct {
	io.Writer
	err error
}

func (w *textRenderWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

type TextDescriptor struct {
	Text           string
	FontName       string
	RTL            bool
	Embed          bool
	FontKey        string
	FontSize       float64
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

func fontVerticalMetrics(fontName string, fontSize float64) (float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func deltaAlignMiddle(fontName string, fontSize float64, lines int, mTop, mBot float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func deltaAlignTop(fontName string, fontSize, mTop float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func deltaAlignBottom(fontName string, fontSize float64, lines int, mBot float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
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

func calcBoundingBoxForLines(lines []string, x, y float64, fontName string, fontSize float64) (*types.Rectangle, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func encodeUserFontRunes(s string) string { _ = "STUB: not implemented"; return "" }

func prepareEmbeddedUserFontBytes(xRefTable *XRefTable, s, fontName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func prepareUserFontBytes(xRefTable *XRefTable, s, fontName string, embed, rtl, fillFont bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func PrepBytes(xRefTable *XRefTable, s, fontName string, embed, rtl, fillFont bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func writeStringToBuf(xRefTable *XRefTable, w io.Writer, s string, x, y float64, td TextDescriptor) error {
	_ = "STUB: not implemented"
	return nil
}

func setFont(w io.Writer, fontID string, fontSize float64) { _ = "STUB: not implemented"; return }

func CalcBoundingBox(s string, x, y float64, fontName string, fontSize int) (*types.Rectangle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CalcBoundingBoxFloat(s string, x, y float64, fontName string, fontSize float64) (*types.Rectangle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func horAdjustBoundingBoxForLines(r, box *types.Rectangle, dx, dy float64, x, y *float64) {
	_ = "STUB: not implemented"
	return
}

func prepJustifiedLine(xRefTable *XRefTable, lines *[]string, strbuf []string, strWidth, w, fontSize float64, fontName string, embed, rtl bool) error {
	_ = "STUB: not implemented"
	return nil
}

type justifiedTextPreparer struct {
	xRefTable  *XRefTable
	strbuf     []string
	strWidth   float64
	indent     bool
	blankWidth float64
}

func newJustifiedTextPreparer(xRefTable *XRefTable, fontName string, fontSize float64) (*justifiedTextPreparer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *justifiedTextPreparer) flush(lines *[]string, w float64, fontName string, fontSize *float64, lastline, parIndent, embed, rtl bool) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *justifiedTextPreparer) add(lines *[]string, s string, w float64, fontName string, fontSize *float64, parIndent, embed, rtl bool) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *justifiedTextPreparer) prepare(lines *[]string, s string, w float64, fontName string, fontSize *float64, lastline, parIndent, embed, rtl bool) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func preRenderJustifiedText(
	xRefTable *XRefTable,
	lines *[]string,
	r *types.Rectangle,
	x, y, width float64,
	td TextDescriptor,
	mLeft, mRight, borderWidth float64,
	fontSize *float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func scaleFontSize(r *types.Rectangle, lines []string, scaleAbs bool,
	scale, width, x, y, mLeft, mRight, borderWidth float64,
	fontName string, fontSize *float64) error {
	_ = "STUB: not implemented"
	return nil
}

func horizontalWrapUp(box *types.Rectangle, maxLine string, hAlign types.HAlignment,
	x *float64, width, ww, mLeft, mRight, borderWidth float64,
	fontName string, fontSize *float64) error {
	_ = "STUB: not implemented"
	return nil
}

func createBoundingBoxForColumn(xRefTable *XRefTable, r *types.Rectangle, x, y *float64,
	width float64,
	td TextDescriptor,
	dx, dy float64,
	mTop, mBot, mLeft, mRight float64,
	borderWidth float64,
	fontSize *float64, lines *[]string) (*types.Rectangle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func flushJustifiedStringToBuf(w io.Writer, s string, x, y float64, strokeCol, fillCol color.SimpleColor, rm draw.RenderMode) error {
	_ = "STUB: not implemented"
	return nil
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

func renderText(xRefTable *XRefTable, w io.Writer, lines []string, td TextDescriptor, x, y, fontSize float64) error {
	_ = "STUB: not implemented"
	return nil
}

func fieldsFunc(s string, f func(rune) bool) []string { _ = "STUB: not implemented"; return nil }

func SplitMultilineStr(s string) []string { _ = "STUB: not implemented"; return nil }

func encodeCoreFontText(s, fontName string) string { _ = "STUB: not implemented"; return "" }

func coreFontTextForWidth(s, fontName string) string { _ = "STUB: not implemented"; return "" }

func textFits(s, fontName string, fontSize, maxWidthPoints float64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func wrapLine(ss *[]string, line, space, word, fontName string, fontSize, maxWidthPoints float64) error {
	_ = "STUB: not implemented"
	return nil
}

func wrapWord(ss *[]string, line, space, word, nextSpace, fontName string, fontSize, maxWidthPoints float64) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func cjkTextWidth(text, fontName string, fontSize float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func wrapCJKRune(ss *[]string, line, space, word, fontName string, c rune, fontSize, maxWidthPoints float64) (
	string, string, string, error,
) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

type textWrapState int

const (
	wrapBeginLine textWrapState = iota
	wrapInWord
	wrapLeadingSpace
	wrapInSpace
)

func wrapTextRune(ss *[]string, line, space, word, fontName string, c rune, fontSize, maxWidthPoints float64) (
	string, string, string, textWrapState, error,
) {
	_ = "STUB: not implemented"
	return "", "", "", *new(textWrapState), nil
}

func wrap(lines []string, fontName string, fontSize, maxWidthPoints float64) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func lastRune(s string) rune { _ = "STUB: not implemented"; return 0 }

func WordWrap(s string, fontName string, fontSize int, maxWidthPoints float64) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WordWrapFloat(s string, fontName string, fontSize, maxWidthPoints float64) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isCJKTextChar(r rune) bool { _ = "STUB: not implemented"; return false }

func isCJKPunctuationChar(r rune) bool { _ = "STUB: not implemented"; return false }

func canBreakAfterChar(r rune) bool { _ = "STUB: not implemented"; return false }

func canBreakBeforeChar(r rune) bool { _ = "STUB: not implemented"; return false }

func isOpeningPunct(r rune) bool { _ = "STUB: not implemented"; return false }

func isClosingPunct(r rune) bool { _ = "STUB: not implemented"; return false }

func scaleTextColumnForRegion(
	mediaBox, region *types.Rectangle,
	dx, dy, width *float64,
	fontSize *float64,
	mTop, mBot, mLeft, mRight, borderWidth *float64,
) *types.Rectangle {
	_ = "STUB: not implemented"
	return nil
}

func positionTextColumn(r *types.Rectangle, x, y, dx, dy float64) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func writeColumn(xRefTable *XRefTable, w io.Writer, mediaBox, region *types.Rectangle, td TextDescriptor, width float64) (*types.Rectangle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WriteColumn(xRefTable *XRefTable, w io.Writer, mediaBox, region *types.Rectangle, td TextDescriptor, width float64) (*types.Rectangle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WriteMultiLine(xRefTable *XRefTable, w io.Writer, mediaBox, region *types.Rectangle, td TextDescriptor) (*types.Rectangle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AnchorPosAndAlign(a types.Anchor, r *types.Rectangle) (x, y float64, hAlign types.HAlignment, vAlign types.VAlignment) {
	_ = "STUB: not implemented"
	return 0, 0, *new(types.HAlignment), *new(types.VAlignment)
}

func WriteMultiLineAnchored(xRefTable *XRefTable, w io.Writer, mediaBox, region *types.Rectangle, td TextDescriptor, a types.Anchor) (*types.Rectangle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WriteColumnAnchored(xRefTable *XRefTable, w io.Writer, mediaBox, region *types.Rectangle, td TextDescriptor, a types.Anchor, width float64) (*types.Rectangle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
