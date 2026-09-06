package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func Trim(rs io.ReadSeeker, w io.Writer, selectedPages []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func TrimFile(inFile, outFile string, selectedPages []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
