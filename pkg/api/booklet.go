package api

import (
	"io"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func validBookletDimension(v float64) bool { _ = "STUB: not implemented"; return false }

func validateBookletGrid(nup *model.NUp) error { _ = "STUB: not implemented"; return nil }

func resolveBookletPageDimension(nup *model.NUp) error { _ = "STUB: not implemented"; return nil }

func validateBookletLayout(nup *model.NUp) error { _ = "STUB: not implemented"; return nil }

func prepareBookletConfiguration(nup *model.NUp) error { _ = "STUB: not implemented"; return nil }

func wrapBookletConfigurationError(err error) error { _ = "STUB: not implemented"; return nil }

func prepareBookletConfigurationForAPI(nup *model.NUp) error { _ = "STUB: not implemented"; return nil }

func BookletFromImages(conf *model.Configuration, imageFileNames []string, nup *model.NUp) (ctx *model.Context, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Booklet(rs io.ReadSeeker, w io.Writer, imgFiles, selectedPages []string, nup *model.NUp, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func bookletImageOutputAliasesInput(inFile, outFile string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func bookletImageOutputAliasesInputWith(
	inFile, outFile string,
	abs func(string) (string, error),
	stat func(string) (os.FileInfo, error),
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func rejectBookletImageOutputAlias(inFiles []string, outFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func BookletFile(inFiles []string, outFile string, selectedPages []string, nup *model.NUp, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
