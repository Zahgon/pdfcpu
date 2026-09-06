package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func optimize(rs io.ReadSeeker, w io.Writer, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func Optimize(rs io.ReadSeeker, w io.Writer, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func OptimizeFile(inFile, outFile string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
