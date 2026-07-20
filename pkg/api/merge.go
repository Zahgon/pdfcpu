package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func appendTo(rs io.ReadSeeker, fName string, ctxDest *model.Context, dividerPage bool) error {
	_ = "STUB: not implemented"
	return nil
}

func appendFile(fName string, ctxDest *model.Context, dividerPage bool) error {
	_ = "STUB: not implemented"
	return nil
}

func MergeRaw(rsc []io.ReadSeeker, w io.Writer, dividerPage bool, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func prepDestContext(destFile string, rs io.ReadSeeker, conf *model.Configuration) (*model.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Merge(destFile string, inFiles []string, w io.Writer, conf *model.Configuration, dividerPage bool) error {
	_ = "STUB: not implemented"
	return nil
}

func MergeCreateFile(inFiles []string, outFile string, dividerPage bool, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func MergeAppendFile(inFiles []string, outFile string, dividerPage bool, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func MergeCreateZip(rs1, rs2 io.ReadSeeker, w io.Writer, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func MergeCreateZipFile(inFile1, inFile2, outFile string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
