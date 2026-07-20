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

func ImportImages(rs io.ReadSeeker, w io.Writer, imgs []io.Reader, imp *pdfcpu.Import, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func fileExists(filename string) bool { _ = "STUB: not implemented"; return false }

func prepImgFiles(imgFiles []string, f1 *os.File) ([]io.ReadCloser, []io.Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func logImportImages(s, outFile string) { _ = "STUB: not implemented"; return }

func importImagesInputFile(outFile string) (io.ReadSeeker, *os.File, string, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadSeeker), nil, "", nil
}

func closeImportImageFiles(rc []io.ReadCloser) error { _ = "STUB: not implemented"; return nil }

func finishImportImagesFile(ok bool, f1, f2 *os.File, rc []io.ReadCloser, tmpFile, outFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func ImportImagesFile(imgFiles []string, outFile string, imp *pdfcpu.Import, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
