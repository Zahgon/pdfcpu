package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func Zoom(rs io.ReadSeeker, w io.Writer, selectedPages []string, zoom *model.Zoom, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ZoomFile(inFile, outFile string, selectedPages []string, zoom *model.Zoom, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
