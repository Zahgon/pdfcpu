package pdfcpu

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/matrix"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func ParseResizeConfig(s string, u types.DisplayUnit) (*model.Resize, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func prepTransform(rSrc, rDest *types.Rectangle, enforce bool) (float64, float64, float64, float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, 0
}

func prepResize(res *model.Resize, cropBox *types.Rectangle) (*types.Rectangle, float64, float64, float64, float64, float64) {
	_ = "STUB: not implemented"
	return nil, 0, 0, 0, 0, 0
}

func handleBgColAndBorder(dx, dy float64, cropBox *types.Rectangle, bb *[]byte, res *model.Resize) {
	_ = "STUB: not implemented"
	return
}

func transformedRect(r *types.Rectangle, m matrix.Matrix) *types.Rectangle {
	_ = "STUB: not implemented"
	return nil
}

func resizeObjectContext(label string, obj types.Object) string {
	_ = "STUB: not implemented"
	return ""
}

func resizeAnnotationRect(ctx *model.Context, d types.Dict, m matrix.Matrix) error {
	_ = "STUB: not implemented"
	return nil
}

func resizeAnnotationQuadPoints(ctx *model.Context, d types.Dict, m matrix.Matrix) error {
	_ = "STUB: not implemented"
	return nil
}

func resizeAnnotation(ctx *model.Context, d types.Dict, m matrix.Matrix) error {
	_ = "STUB: not implemented"
	return nil
}

func resizePageAnnotations(ctx *model.Context, d types.Dict, m matrix.Matrix) error {
	_ = "STUB: not implemented"
	return nil
}

func resizePage(ctx *model.Context, pageNr int, res *model.Resize) error {
	_ = "STUB: not implemented"
	return nil
}

func resizePageNumbers(pageCount int, selectedPages types.IntSet) []int {
	_ = "STUB: not implemented"
	return nil
}

func Resize(ctx *model.Context, selectedPages types.IntSet, res *model.Resize) error {
	_ = "STUB: not implemented"
	return nil
}
