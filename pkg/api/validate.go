package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func Validate(rs io.ReadSeeker, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ValidateFile(inFile string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateFiles(inFiles []string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func DumpObject(rs io.ReadSeeker, mode, objNr int, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func DumpObjectFile(inFile string, mode, objNr int, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}
