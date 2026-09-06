package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func prepareImagesContext(rs io.ReadSeeker, selectedPages []string, conf *model.Configuration) (*model.Context, types.IntSet, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.IntSet), nil
}

func Images(rs io.ReadSeeker, selectedPages []string, conf *model.Configuration) (ii []map[int]model.Image, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListImages(rs io.ReadSeeker, selectedPages []string, conf *model.Configuration) (ss []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateImageSelection(objNr, pageNr int, id string) error {
	_ = "STUB: not implemented"
	return nil
}

func UpdateImages(rs io.ReadSeeker, rd io.Reader, w io.Writer, objNr, pageNr int, id string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func imageSelectionFromFile(imageFile string) (int, string, error) {
	_ = "STUB: not implemented"
	return 0, "", nil
}

func resolveImageSelection(imageFile string, objNr, pageNr int, id string) (int, string, error) {
	_ = "STUB: not implemented"
	return 0, "", nil
}

func ValidateUpdateImagesOutput(imageFile, outFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func UpdateImagesFile(inFile, imageFile, outFile string, objNr, pageNr int, id string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
