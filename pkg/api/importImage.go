package api

import (
	"io"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func Import(s string, u types.DisplayUnit) (*pdfcpu.Import, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DefaultImportConfig() *pdfcpu.Import { _ = "STUB: not implemented"; return nil }

func validateImportConfiguration(imp *pdfcpu.Import) error { _ = "STUB: not implemented"; return nil }

func invalidImportNumber(v float64) bool { _ = "STUB: not implemented"; return false }

func invalidImportScale(scale float64, absolute bool) bool { _ = "STUB: not implemented"; return false }

func PrepareImportConfiguration(imp *pdfcpu.Import) (*pdfcpu.Import, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateImportImageReaders(imgs []io.Reader) error { _ = "STUB: not implemented"; return nil }

func validateImportImageFiles(imgFiles []string) error { _ = "STUB: not implemented"; return nil }

func validateImportImagesOutput(imgFiles []string, outFile string, skipStdinMarker bool) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateImportImagesOutput(imgFiles []string, outFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func importImagesContext(rs io.ReadSeeker, imp *pdfcpu.Import, conf *model.Configuration) (*model.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func importImagesPageTree(ctx *model.Context) (*types.IndirectRef, types.Dict, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.Dict), nil
}

func appendImportedImagePages(
	ctx *model.Context,
	indRefs []*types.IndirectRef,
	imageIndex int,
	pagesDict types.Dict,
) error {
	_ = "STUB: not implemented"
	return nil
}

func appendImportedImages(
	ctx *model.Context,
	imgs []io.Reader,
	pagesIndRef *types.IndirectRef,
	pagesDict types.Dict,
	imp *pdfcpu.Import,
) error {
	_ = "STUB: not implemented"
	return nil
}

func ImportImages(rs io.ReadSeeker, w io.Writer, imgs []io.Reader, imp *pdfcpu.Import, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type importImageFileCloser struct {
	closer     io.Closer
	imageIndex int
	fileName   string
}

func prepImgFiles(imgFiles []string) ([]importImageFileCloser, []io.Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func logImportImages(s, outFile string) { _ = "STUB: not implemented"; return }

func importImagesInputFile(outFile string) (io.ReadSeeker, *os.File, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadSeeker), nil, nil
}

func closeImportImageInputs(rc []importImageFileCloser) error {
	_ = "STUB: not implemented"
	return nil
}

func ImportImagesFile(imgFiles []string, outFile string, imp *pdfcpu.Import, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
