package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func PDFNUpConfig(val int, desc string, conf *model.Configuration) (*model.NUp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ImageNUpConfig(val int, desc string, conf *model.Configuration) (*model.NUp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PDFGridConfig(rows, cols int, desc string, conf *model.Configuration) (*model.NUp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ImageGridConfig(rows, cols int, desc string, conf *model.Configuration) (*model.NUp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PDFBookletConfig(val int, desc string, conf *model.Configuration) (*model.NUp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ImageBookletConfig(val int, desc string, conf *model.Configuration) (*model.NUp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NUpFromImage(conf *model.Configuration, imageFileNames []string, nup *model.NUp) (*model.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NUp(rs io.ReadSeeker, w io.Writer, imgFiles, selectedPages []string, nup *model.NUp, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func NUpFile(inFiles []string, outFile string, selectedPages []string, nup *model.NUp, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
