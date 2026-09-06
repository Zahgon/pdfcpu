package color

import (
	"errors"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

var (
	Black     = SimpleColor{}
	White     = SimpleColor{R: 1, G: 1, B: 1}
	LightGray = SimpleColor{.9, .9, .9}
	Gray      = SimpleColor{.5, .5, .5}
	DarkGray  = SimpleColor{.3, .3, .3}
	Red       = SimpleColor{1, 0, 0}
	Green     = SimpleColor{0, 1, 0}
	Blue      = SimpleColor{0, 0, 1}
	Yellow    = SimpleColor{.5, .5, 0}
)

var ErrInvalidColor = errors.New("invalid color constant")

type SimpleColor struct {
	R, G, B float32
}

func (sc SimpleColor) String() string { _ = "STUB: not implemented"; return "" }

func (sc SimpleColor) Array() types.Array { _ = "STUB: not implemented"; return *new(types.Array) }

func NewSimpleColor(rgb uint32) SimpleColor { _ = "STUB: not implemented"; return *new(SimpleColor) }

func NewSimpleColorForArray(arr types.Array) SimpleColor {
	_ = "STUB: not implemented"
	return *new(SimpleColor)
}

func NewSimpleColorForHexCode(hexCol string) (SimpleColor, error) {
	_ = "STUB: not implemented"
	return *new(SimpleColor), nil
}

func internalSimpleColor(s string) (SimpleColor, error) {
	_ = "STUB: not implemented"
	return *new(SimpleColor), nil
}

func ParseColor(s string) (SimpleColor, error) {
	_ = "STUB: not implemented"
	return *new(SimpleColor), nil
}
