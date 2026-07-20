package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func Images(rs io.ReadSeeker, selectedPages []string, conf *model.Configuration) (ii []map[int]model.Image, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UpdateImages(rs io.ReadSeeker, rd io.Reader, w io.Writer, objNr, pageNr int, id string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ensurePageNrAndId(pageNr *int, id *string, imageFile string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func UpdateImagesFile(inFile, imageFile, outFile string, objNr, pageNr int, id string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
