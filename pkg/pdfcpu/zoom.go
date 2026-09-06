package pdfcpu

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func ParseZoomConfig(s string, u types.DisplayUnit) (*model.Zoom, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func handleZoomOutBgColAndBorder(cropBox *types.Rectangle, bb *[]byte, zoom *model.Zoom) {
	_ = "STUB: not implemented"
	return
}

func zoomPage(ctx *model.Context, pageNr int, zoom *model.Zoom) error {
	_ = "STUB: not implemented"
	return nil
}

func zoomPageNumbers(pageCount int, selectedPages types.IntSet) []int {
	_ = "STUB: not implemented"
	return nil
}

func Zoom(ctx *model.Context, selectedPages types.IntSet, zoom *model.Zoom) error {
	_ = "STUB: not implemented"
	return nil
}
