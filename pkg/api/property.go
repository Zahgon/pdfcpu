package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func Properties(rs io.ReadSeeker, conf *model.Configuration) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AddProperties(rs io.ReadSeeker, w io.Writer, properties map[string]string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func AddPropertiesFile(inFile, outFile string, properties map[string]string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemoveProperties(rs io.ReadSeeker, w io.Writer, properties []string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func RemovePropertiesFile(inFile, outFile string, properties []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
