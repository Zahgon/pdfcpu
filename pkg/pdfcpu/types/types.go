package types

import (
	"fmt"
)

const (
	EolLF   = "\x0A"
	EolCR   = "\x0D"
	EolCRLF = "\x0D\x0A"
)

const FreeHeadGeneration = 65535

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
)

func (b ByteSize) String() string { _ = "STUB: not implemented"; return "" }

type IntSet map[int]bool

type StringSet map[string]bool

type Object interface {
	fmt.Stringer
	Clone() Object
	PDFString() string
}

type Boolean bool

func (boolean Boolean) Clone() Object { _ = "STUB: not implemented"; return *new(Object) }

func (boolean Boolean) String() string { _ = "STUB: not implemented"; return "" }

func (boolean Boolean) PDFString() string { _ = "STUB: not implemented"; return "" }

func (boolean Boolean) Value() bool { _ = "STUB: not implemented"; return false }

type Float float64

func (f Float) Clone() Object { _ = "STUB: not implemented"; return *new(Object) }

func (f Float) String() string { _ = "STUB: not implemented"; return "" }

func (f Float) PDFString() string { _ = "STUB: not implemented"; return "" }

func (f Float) Value() float64 { _ = "STUB: not implemented"; return 0 }

type Integer int

func (i Integer) Clone() Object { _ = "STUB: not implemented"; return *new(Object) }

func (i Integer) String() string { _ = "STUB: not implemented"; return "" }

func (i Integer) PDFString() string { _ = "STUB: not implemented"; return "" }

func (i Integer) Value() int { _ = "STUB: not implemented"; return 0 }

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func NewPoint(x, y float64) Point { _ = "STUB: not implemented"; return *new(Point) }

func (p *Point) Translate(dx, dy float64) { _ = "STUB: not implemented"; return }

func (p Point) String() string { _ = "STUB: not implemented"; return "" }

type Rectangle struct {
	LL Point `json:"ll"`
	UR Point `json:"ur"`
}

func NewRectangle(llx, lly, urx, ury float64) *Rectangle { _ = "STUB: not implemented"; return nil }

func decodeFloat(number Object) float64 { _ = "STUB: not implemented"; return 0 }

func RectForArray(arr Array) *Rectangle { _ = "STUB: not implemented"; return nil }

func RectForDim(width, height float64) *Rectangle { _ = "STUB: not implemented"; return nil }

func RectForWidthAndHeight(llx, lly, width, height float64) *Rectangle {
	_ = "STUB: not implemented"
	return nil
}

func RectForFormat(f string) *Rectangle { _ = "STUB: not implemented"; return nil }

func (r Rectangle) Width() float64 { _ = "STUB: not implemented"; return 0 }

func (r Rectangle) Height() float64 { _ = "STUB: not implemented"; return 0 }

func (r Rectangle) Equals(r2 Rectangle) bool { _ = "STUB: not implemented"; return false }

func (r Rectangle) FitsWithin(r2 *Rectangle) bool { _ = "STUB: not implemented"; return false }

func (r Rectangle) Visible() bool { _ = "STUB: not implemented"; return false }

func (r Rectangle) AspectRatio() float64 { _ = "STUB: not implemented"; return 0 }

func (r Rectangle) Landscape() bool { _ = "STUB: not implemented"; return false }

func (r Rectangle) Portrait() bool { _ = "STUB: not implemented"; return false }

func (r Rectangle) Contains(p Point) bool { _ = "STUB: not implemented"; return false }

func (r Rectangle) ScaledWidth(h float64) float64 { _ = "STUB: not implemented"; return 0 }

func (r Rectangle) ScaledHeight(w float64) float64 { _ = "STUB: not implemented"; return 0 }

func (r Rectangle) Dimensions() Dim { _ = "STUB: not implemented"; return *new(Dim) }

func (r *Rectangle) Translate(dx, dy float64) { _ = "STUB: not implemented"; return }

func (r Rectangle) Center() Point { _ = "STUB: not implemented"; return *new(Point) }

func (r Rectangle) String() string { _ = "STUB: not implemented"; return "" }

func (r Rectangle) ShortString() string { _ = "STUB: not implemented"; return "" }

func (r Rectangle) Array() Array { _ = "STUB: not implemented"; return *new(Array) }

func (r Rectangle) Clone() *Rectangle { _ = "STUB: not implemented"; return nil }

func (r Rectangle) CroppedCopy(margin float64) *Rectangle { _ = "STUB: not implemented"; return nil }

func (r Rectangle) ToInches() *Rectangle { _ = "STUB: not implemented"; return nil }

func (r Rectangle) ToCentimetres() *Rectangle { _ = "STUB: not implemented"; return nil }

func (r Rectangle) ToMillimetres() *Rectangle { _ = "STUB: not implemented"; return nil }

func (r *Rectangle) ConvertToUnit(unit DisplayUnit) *Rectangle {
	_ = "STUB: not implemented"
	return nil
}

func (r Rectangle) formatToInches() string { _ = "STUB: not implemented"; return "" }

func (r Rectangle) formatToCentimetres() string { _ = "STUB: not implemented"; return "" }

func (r Rectangle) formatToMillimetres() string { _ = "STUB: not implemented"; return "" }

func (r Rectangle) Format(unit DisplayUnit) string { _ = "STUB: not implemented"; return "" }

type QuadLiteral struct {
	P1, P2, P3, P4 Point
}

func NewQuadLiteralForRect(r *Rectangle) *QuadLiteral { _ = "STUB: not implemented"; return nil }

func (ql QuadLiteral) Array() Array { _ = "STUB: not implemented"; return *new(Array) }

func (ql QuadLiteral) EnclosingRectangle(f float64) *Rectangle {
	_ = "STUB: not implemented"
	return nil
}

type QuadPoints []QuadLiteral

func (qp *QuadPoints) AddQuadLiteral(ql QuadLiteral) { _ = "STUB: not implemented"; return }

func (qp *QuadPoints) Array() Array { _ = "STUB: not implemented"; return *new(Array) }

type Name string

func (nameObject Name) Clone() Object { _ = "STUB: not implemented"; return *new(Object) }

func (nameObject Name) String() string { _ = "STUB: not implemented"; return "" }

func (nameObject Name) PDFString() string { _ = "STUB: not implemented"; return "" }

func (nameObject Name) Value() string { _ = "STUB: not implemented"; return "" }

type StringLiteral string

func (stringliteral StringLiteral) Clone() Object { _ = "STUB: not implemented"; return *new(Object) }

func (stringliteral StringLiteral) String() string { _ = "STUB: not implemented"; return "" }

func (stringliteral StringLiteral) PDFString() string { _ = "STUB: not implemented"; return "" }

func (stringliteral StringLiteral) Value() string { _ = "STUB: not implemented"; return "" }

type HexLiteral string

func NewHexLiteral(b []byte) HexLiteral { _ = "STUB: not implemented"; return *new(HexLiteral) }

func (hexliteral HexLiteral) Clone() Object { _ = "STUB: not implemented"; return *new(Object) }

func (hexliteral HexLiteral) String() string { _ = "STUB: not implemented"; return "" }

func (hexliteral HexLiteral) PDFString() string { _ = "STUB: not implemented"; return "" }

func (hexliteral HexLiteral) Value() string { _ = "STUB: not implemented"; return "" }

func (hexliteral HexLiteral) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type IndirectRef struct {
	ObjectNumber     Integer
	GenerationNumber Integer
}

func NewIndirectRef(objectNumber, generationNumber int) *IndirectRef {
	_ = "STUB: not implemented"
	return nil
}

func (ir IndirectRef) Clone() Object { _ = "STUB: not implemented"; return *new(Object) }

func (ir IndirectRef) String() string { _ = "STUB: not implemented"; return "" }

func (ir IndirectRef) PDFString() string { _ = "STUB: not implemented"; return "" }

type DisplayUnit int

const (
	POINTS DisplayUnit = iota
	INCHES
	CENTIMETRES
	MILLIMETRES
)

const (
	userSpaceToInch = float64(1) / 72
	userSpaceToCm   = 2.54 / 72
	userSpaceToMm   = userSpaceToCm * 10

	inchToUserSpace = 1 / userSpaceToInch
	cmToUserSpace   = 1 / userSpaceToCm
	mmToUserSpace   = 1 / userSpaceToMm
)

func ToUserSpace(f float64, unit DisplayUnit) float64 { _ = "STUB: not implemented"; return 0 }

type Dim struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func (d Dim) ToInches() Dim { _ = "STUB: not implemented"; return *new(Dim) }

func (d Dim) ToCentimetres() Dim { _ = "STUB: not implemented"; return *new(Dim) }

func (d Dim) ToMillimetres() Dim { _ = "STUB: not implemented"; return *new(Dim) }

func (d Dim) ConvertToUnit(unit DisplayUnit) Dim { _ = "STUB: not implemented"; return *new(Dim) }

func (d Dim) AspectRatio() float64 { _ = "STUB: not implemented"; return 0 }

func (d Dim) Landscape() bool { _ = "STUB: not implemented"; return false }

func (d Dim) Portrait() bool { _ = "STUB: not implemented"; return false }

func (d Dim) String() string { _ = "STUB: not implemented"; return "" }
