package types

type Anchor int

func (a Anchor) String() string { _ = "STUB: not implemented"; return "" }

const (
	TopLeft Anchor = iota
	TopCenter
	TopRight
	Left
	Center
	Right
	BottomLeft
	BottomCenter
	BottomRight
	Full
)
