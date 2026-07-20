package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func Collect(rs io.ReadSeeker, w io.Writer, selectedPages []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func CollectFile(inFile, outFile string, selectedPages []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
