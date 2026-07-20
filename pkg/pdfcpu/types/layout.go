package types

type Corner int

const (
	LowerLeft Corner = iota
	LowerRight
	UpperLeft
	UpperRight
)

type HAlignment int

const (
	AlignLeft HAlignment = iota
	AlignCenter
	AlignRight
	AlignJustify
)

type VAlignment int

const (
	AlignBaseline VAlignment = iota
	AlignTop
	AlignMiddle
	AlignBottom
)

type LineJoinStyle int

const (
	LJMiter LineJoinStyle = iota
	LJRound
	LJBevel
)

func ParseHorAlignment(s string) (HAlignment, error) {
	_ = "STUB: not implemented"
	return *new(HAlignment), nil
}

func ParseOrigin(s string) (Corner, error) { _ = "STUB: not implemented"; return *new(Corner), nil }

func ParseAnchor(s string) (Anchor, error) { _ = "STUB: not implemented"; return *new(Anchor), nil }

func ParsePositionAnchor(s string) (Anchor, error) {
	_ = "STUB: not implemented"
	return *new(Anchor), nil
}

func AnchorPosition(a Anchor, r *Rectangle, w, h float64) (x float64, y float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

type Orientation int

const (
	Horizontal Orientation = iota
	Vertical
)

type RelPosition int

const (
	RelPosLeft RelPosition = iota
	RelPosRight
	RelPosTop
	RelPosBottom
)

func ParseRelPosition(s string) (RelPosition, error) {
	_ = "STUB: not implemented"
	return *new(RelPosition), nil
}

func NormalizeCoord(x, y float64, r *Rectangle, origin Corner, absolute bool) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func NormalizeOffset(x, y float64, origin Corner) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func BestFitRectIntoRect(rSrc, rDest *Rectangle, enforceOrient, scaleUp bool) (w, h, dx, dy, rot float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, 0
}

func ParsePageFormat(v string) (*Dim, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}
