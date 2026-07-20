package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func Rotate(rs io.ReadSeeker, w io.Writer, rotation int, selectedPages []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RotateFile(inFile, outFile string, rotation int, selectedPages []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
