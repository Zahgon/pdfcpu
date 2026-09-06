package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func validNUpDimension(v float64) bool { _ = "STUB: not implemented"; return false }

func validateNUpGrid(nup *model.NUp) error { _ = "STUB: not implemented"; return nil }

func resolveNUpPageDimension(nup *model.NUp) error { _ = "STUB: not implemented"; return nil }

func prepareNUpConfiguration(nup *model.NUp, imageInput bool) error {
	_ = "STUB: not implemented"
	return nil
}

func prepareNUpConfigurationForAPI(nup *model.NUp, imageInput bool) error {
	_ = "STUB: not implemented"
	return nil
}

func NUpValuesForBooklets() []int { _ = "STUB: not implemented"; return nil }

func NUpValues() []int { _ = "STUB: not implemented"; return nil }

func DefaultBookletConfig() *model.NUp { _ = "STUB: not implemented"; return nil }

func PDFNUpConfig(val int, desc string, conf *model.Configuration) (*model.NUp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ImageNUpConfig(val int, desc string, conf *model.Configuration) (*model.NUp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseNUpDetails(s string, nup *model.NUp) error { _ = "STUB: not implemented"; return nil }

func ParseNUpValue(n int, nup *model.NUp) error { _ = "STUB: not implemented"; return nil }

func PDFBookletConfig(val int, desc string, conf *model.Configuration) (*model.NUp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ImageBookletConfig(val int, desc string, conf *model.Configuration) (*model.NUp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NUpFromImage(conf *model.Configuration, imageFileNames []string, nup *model.NUp) (ctx *model.Context, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NUp(rs io.ReadSeeker, w io.Writer, imgFiles, selectedPages []string, nup *model.NUp, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func nUpImageOutputAliasesInput(inFile, outFile string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func rejectNUpImageOutputAlias(inFiles []string, outFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func NUpFile(inFiles []string, outFile string, selectedPages []string, nup *model.NUp, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
