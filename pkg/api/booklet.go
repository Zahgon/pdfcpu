package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func BookletFromImages(conf *model.Configuration, imageFileNames []string, nup *model.NUp) (*model.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Booklet(rs io.ReadSeeker, w io.Writer, imgFiles, selectedPages []string, nup *model.NUp, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func BookletFile(inFiles []string, outFile string, selectedPages []string, nup *model.NUp, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
