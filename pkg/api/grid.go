package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func PDFGridConfig(rows, cols int, desc string, conf *model.Configuration) (*model.NUp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ImageGridConfig(rows, cols int, desc string, conf *model.Configuration) (*model.NUp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseGridDefinition(rows, cols int, nup *model.NUp) error {
	_ = "STUB: not implemented"
	return nil
}

func ParseNUpGridDefinition(rows, cols int, nup *model.NUp) error {
	_ = "STUB: not implemented"
	return nil
}

func prepareGridConfiguration(nup *model.NUp, imageInput bool) error {
	_ = "STUB: not implemented"
	return nil
}

func prepareGridConfigurationForAPI(nup *model.NUp, imageInput bool) error {
	_ = "STUB: not implemented"
	return nil
}

func GridFromImage(conf *model.Configuration, imageFileNames []string, nup *model.NUp) (ctx *model.Context, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Grid(rs io.ReadSeeker, w io.Writer, imgFiles, selectedPages []string, nup *model.NUp, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func rejectGridImageOutputAlias(inFiles []string, outFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func GridFile(inFiles []string, outFile string, selectedPages []string, nup *model.NUp, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
