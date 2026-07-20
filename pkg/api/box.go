package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func PageBoundariesFromBoxList(s string) (*model.PageBoundaries, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PageBoundaries(s string, unit types.DisplayUnit) (*model.PageBoundaries, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Box(s string, u types.DisplayUnit) (*model.Box, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Boxes(rs io.ReadSeeker, selectedPages []string, conf *model.Configuration) (pb []model.PageBoundaries, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AddBoxes(rs io.ReadSeeker, w io.Writer, selectedPages []string, pb *model.PageBoundaries, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AddBoxesFile(inFile, outFile string, selectedPages []string, pb *model.PageBoundaries, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemoveBoxes(rs io.ReadSeeker, w io.Writer, selectedPages []string, pb *model.PageBoundaries, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemoveBoxesFile(inFile, outFile string, selectedPages []string, pb *model.PageBoundaries, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func Crop(rs io.ReadSeeker, w io.Writer, selectedPages []string, b *model.Box, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func CropFile(inFile, outFile string, selectedPages []string, b *model.Box, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
