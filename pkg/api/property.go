package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func Properties(rs io.ReadSeeker, conf *model.Configuration) (m map[string]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AddProperties(rs io.ReadSeeker, w io.Writer, properties map[string]string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type propertyMutation func(io.ReadSeeker, io.Writer, *model.Configuration) error

func mutatePropertiesFile(
	inFile, outFile string,
	conf *model.Configuration,
	op string,
	mutate propertyMutation,
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AddPropertiesFile(inFile, outFile string, properties map[string]string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func RemoveProperties(rs io.ReadSeeker, w io.Writer, properties []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemovePropertiesFile(inFile, outFile string, properties []string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}
