package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func Permissions(rs io.ReadSeeker, conf *model.Configuration) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func SetPermissions(rs io.ReadSeeker, w io.Writer, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func SetPermissionsFile(inFile, outFile string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func GetPermissions(rs io.ReadSeeker, conf *model.Configuration) (*int16, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetPermissionsFile(inFile string, conf *model.Configuration) (*int16, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
