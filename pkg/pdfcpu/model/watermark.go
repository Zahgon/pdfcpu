package model

import (
	"io"
	"math"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/draw"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/matrix"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

const (
	DegToRad = math.Pi / 180
	RadToDeg = 180 / math.Pi
)

const (
	NoDiagonal = iota
	DiagonalLLToUR
	DiagonalULToLR
)

const (
	WMText = iota
	WMImage
	WMPDF
)

type formCache map[types.Rectangle]*types.IndirectRef

type PdfResources struct {
	Content []byte
	ResDict *types.IndirectRef
	Bb      *types.Rectangle
}

type Watermark struct {
	OnTop                     bool
	Mode                      int
	FileName                  string
	Image                     io.Reader
	PDF                       io.ReadSeeker
	TextString                string
	TextLines                 []string
	URL                       string
	InpUnit                   types.DisplayUnit
	Pos                       types.Anchor
	Dx, Dy                    float64
	HAlign                    *types.HAlignment
	FontName                  string
	FontSize                  int
	ScaledFontSize            int
	ScriptName                string
	RTL                       bool
	Color                     color.SimpleColor
	FillColor                 color.SimpleColor
	StrokeColor               color.SimpleColor
	BgColor                   *color.SimpleColor
	MLeft, MRight             float64
	MTop, MBot                float64
	BorderWidth               float64
	BorderStyle               types.LineJoinStyle
	BorderColor               *color.SimpleColor
	Rotation                  float64
	Diagonal                  int
	UserRotOrDiagonal         bool
	Opacity                   float64
	RenderMode                draw.RenderMode
	Scale                     float64
	ScaleEff                  float64
	ScaleAbs                  bool
	Update                    bool
	Ocg, ExtGState, Font, Img *types.IndirectRef
	Width, Height             int

	bbPDF                   *types.Rectangle
	PdfRes                  map[int]PdfResources
	PdfPageNrSrc            int
	PdfMultiStartPageNrSrc  int
	PdfMultiStartPageNrDest int

	Bb      *types.Rectangle
	BbTrans types.QuadLiteral
	Vp      *types.Rectangle
	PageRot int
	Form    *types.IndirectRef

	Objs   types.IntSet
	FCache formCache
}

func DefaultWatermarkConfig() *Watermark { _ = "STUB: not implemented"; return nil }

func (wm *Watermark) Recycle() { _ = "STUB: not implemented"; return }

func (wm Watermark) IsText() bool { _ = "STUB: not implemented"; return false }

func (wm Watermark) IsPDF() bool { _ = "STUB: not implemented"; return false }

func (wm Watermark) IsImage() bool { _ = "STUB: not implemented"; return false }

func (wm Watermark) Typ() string { _ = "STUB: not implemented"; return "" }

func (wm Watermark) String() string { _ = "STUB: not implemented"; return "" }

func (wm Watermark) OnTopString() string { _ = "STUB: not implemented"; return "" }

func (wm Watermark) MultiStamp() bool { _ = "STUB: not implemented"; return false }

func (wm *Watermark) CalcBoundingBox(pageNr int) { _ = "STUB: not implemented"; return }

func LowerLeftCorner(vp *types.Rectangle, bbw, bbh float64, a types.Anchor) types.Point {
	_ = "STUB: not implemented"
	return *new(types.Point)
}

func (wm *Watermark) alignWithPageBoundariesForNegRot() (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (wm *Watermark) alignWithPageBoundariesForPosRot() (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (wm *Watermark) alignWithPageBoundaries() (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (wm *Watermark) CalcTransformMatrix() matrix.Matrix {
	_ = "STUB: not implemented"
	return *new(matrix.Matrix)
}

func (wm *Watermark) PdfResIndex(pageNr int) int { _ = "STUB: not implemented"; return 0 }
