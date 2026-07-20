package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func Resize(rs io.ReadSeeker, w io.Writer, selectedPages []string, resize *model.Resize, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ResizeFile(inFile, outFile string, selectedPages []string, resize *model.Resize, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
