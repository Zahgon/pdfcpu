package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func validationModeHint(mode int) string { _ = "STUB: not implemented"; return "" }

func validationError(ctx *model.Context, conf *model.Configuration, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func Validate(rs io.ReadSeeker, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ValidateFile(inFile string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ValidateFiles(inFiles []string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}
