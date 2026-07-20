package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func prepareForCut(rs io.ReadSeeker, selectedPages []string, conf *model.Configuration) (*model.Context, types.IntSet, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.IntSet), nil
}

func Poster(rs io.ReadSeeker, outDir, fileName string, selectedPages []string, cut *model.Cut, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func PosterFile(inFile, outDir, outFile string, selectedPages []string, cut *model.Cut, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func NDown(rs io.ReadSeeker, outDir, fileName string, selectedPages []string, n int, cut *model.Cut, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func NDownFile(inFile, outDir, outFile string, selectedPages []string, n int, cut *model.Cut, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCut(cut *model.Cut) error { _ = "STUB: not implemented"; return nil }

func Cut(rs io.ReadSeeker, outDir, fileName string, selectedPages []string, cut *model.Cut, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func CutFile(inFile, outDir, outFile string, selectedPages []string, cut *model.Cut, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}
