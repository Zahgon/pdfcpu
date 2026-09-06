package matrix

import (
	"math"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

const (
	DegToRad = math.Pi / 180
	RadToDeg = 180 / math.Pi
)

type Matrix [3][3]float64

var IdentMatrix = Matrix{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

func (m Matrix) Multiply(n Matrix) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (m Matrix) Transform(p types.Point) types.Point {
	_ = "STUB: not implemented"
	return *new(types.Point)
}

func (m Matrix) String() string { _ = "STUB: not implemented"; return "" }

func CalcTransformMatrix(sx, sy, sin, cos, dx, dy float64) Matrix {
	_ = "STUB: not implemented"
	return *new(Matrix)
}

func CalcRotateAndTranslateTransformMatrix(r, dx, dy float64) Matrix {
	_ = "STUB: not implemented"
	return *new(Matrix)
}

func CalcRotateTransformMatrix(rot float64, bb *types.Rectangle) Matrix {
	_ = "STUB: not implemented"
	return *new(Matrix)
}
