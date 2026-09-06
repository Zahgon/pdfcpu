package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

type UnsupportedResourceError struct {
	Err error
}

func (e *UnsupportedResourceError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *UnsupportedResourceError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func unsupportedResourceError(err error) error { _ = "STUB: not implemented"; return nil }

func skipUnsupportedResource(err error, conf *model.Configuration) bool {
	_ = "STUB: not implemented"
	return false
}

type extractionErrorWithoutSkipMarker struct {
	message string
	cause   error
}

func (e extractionErrorWithoutSkipMarker) Error() string { _ = "STUB: not implemented"; return "" }

func (e extractionErrorWithoutSkipMarker) Unwrap() error { _ = "STUB: not implemented"; return nil }

func joinExtractionCleanupError(err, cleanupErr error) error { _ = "STUB: not implemented"; return nil }

func digestImages(mm map[int]model.Image, singleImgPerPage bool, maxPageDigits int, digestImage func(model.Image, bool, int) error) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func sanitizeFilenamePart(s, fallback string) string { _ = "STUB: not implemented"; return "" }

func WriteImageToDisk(outDir, fileName string) func(model.Image, bool, int) error {
	_ = "STUB: not implemented"
	return nil
}

func WriteFontToDisk(outDir, fnBase string) func(pdfcpu.Font) error {
	_ = "STUB: not implemented"
	return nil
}

func WritePageToDisk(outDir, fnBase string) func(io.Reader, int) error {
	_ = "STUB: not implemented"
	return nil
}

func WriteContentToDisk(outDir, fnBase string) func(io.Reader, int) error {
	_ = "STUB: not implemented"
	return nil
}

func WriteMetadataToDisk(outDir, fnBase string) func(pdfcpu.Metadata) error {
	_ = "STUB: not implemented"
	return nil
}

func ExtractImagesRaw(rs io.ReadSeeker, selectedPages []string, conf *model.Configuration) (images []map[int]model.Image, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExtractImages(rs io.ReadSeeker, selectedPages []string, digestImage func(model.Image, bool, int) error, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ExtractImagesFile(inFile, outDir string, selectedPages []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func writeFonts(ff []pdfcpu.Font, digestFont func(pdfcpu.Font) error) error {
	_ = "STUB: not implemented"
	return nil
}

func ExtractFonts(rs io.ReadSeeker, selectedPages []string, digestFont func(pdfcpu.Font) error, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ExtractFontsFile(inFile, outDir string, selectedPages []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ExtractPage(ctx *model.Context, pageNr int) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func ExtractPages(rs io.ReadSeeker, selectedPages []string, digestPage func(io.Reader, int) error, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ExtractPagesFile(inFile, outDir string, selectedPages []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ExtractContent(rs io.ReadSeeker, selectedPages []string, digestContent func(io.Reader, int) error, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ExtractContentFile(inFile, outDir string, selectedPages []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ExtractMetadata(rs io.ReadSeeker, digestMetadata func(pdfcpu.Metadata) error, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ExtractMetadataFile(inFile, outDir string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
