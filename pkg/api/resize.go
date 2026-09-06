package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func validateResizeConfiguration(res *model.Resize) error { _ = "STUB: not implemented"; return nil }

func invalidResizeScale(res *model.Resize) bool { _ = "STUB: not implemented"; return false }

func invalidResizeDimension(v float64) bool { _ = "STUB: not implemented"; return false }

func Resize(rs io.ReadSeeker, w io.Writer, selectedPages []string, resize *model.Resize, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ResizeFile(inFile, outFile string, selectedPages []string, resize *model.Resize, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
