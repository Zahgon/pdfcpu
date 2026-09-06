package api

import (
	"io"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func cloneCutConfiguration(cut *model.Cut) *model.Cut { _ = "STUB: not implemented"; return nil }

func invalidCutNumber(v float64) bool { _ = "STUB: not implemented"; return false }

func validateCutOptions(cut *model.Cut) error { _ = "STUB: not implemented"; return nil }

func validateCutPoints(name string, points []float64) error { _ = "STUB: not implemented"; return nil }

func validateCutConfiguration(cut *model.Cut) error { _ = "STUB: not implemented"; return nil }

func validateNDownConfiguration(n int, cut *model.Cut) error { _ = "STUB: not implemented"; return nil }

func validatePosterConfiguration(cut *model.Cut) error { _ = "STUB: not implemented"; return nil }

func selectedCutPages(pageCount int, selectedPages []string, operation string) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func prepareForCut(rs io.ReadSeeker, selectedPages []string, conf *model.Configuration, operation string) (*model.Context, []int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type cutOutputFile interface {
	io.Writer
	io.Closer
	Chmod(os.FileMode) error
	Name() string
}

type cutOutputOperations struct {
	stat          func(string) (os.FileInfo, error)
	createTemp    func(string, string) (cutOutputFile, error)
	writeAndFlush func(*model.Context, io.Writer) (error, error)
	rename        func(string, string) error
	remove        func(string) error
}

func createCutTemporaryOutput(dir, pattern string) (cutOutputFile, error) {
	_ = "STUB: not implemented"
	return *new(cutOutputFile), nil
}

func defaultCutOutputOperations() cutOutputOperations {
	_ = "STUB: not implemented"
	return *new(cutOutputOperations)
}

func cutDestinationMode(outFile, operation string, ops cutOutputOperations) (os.FileMode, bool, error) {
	_ = "STUB: not implemented"
	return *new(os.FileMode), false, nil
}

func writeAndFlushCutContext(ctx *model.Context, w io.Writer) (error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cutOutputPhaseError(err error, operation, outFile, phase string) error {
	_ = "STUB: not implemented"
	return nil
}

func removeCutTemporaryOutput(tmpFile, operation, outFile string, ops cutOutputOperations) error {
	_ = "STUB: not implemented"
	return nil
}

func writeCutOutputWith(ctx *model.Context, outFile, operation string, ops cutOutputOperations) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func writeCutOutput(ctx *model.Context, outFile, operation string) error {
	_ = "STUB: not implemented"
	return nil
}

func writePosterPage(ctxSrc *model.Context, pageNr int, outDir, fileName string, cut *model.Cut, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func Poster(rs io.ReadSeeker, outDir, fileName string, selectedPages []string, cut *model.Cut, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func PosterFile(inFile, outDir, outFile string, selectedPages []string, cut *model.Cut, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func writeNDownPage(ctxSrc *model.Context, pageNr, n int, outDir, fileName string, cut *model.Cut, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func NDown(rs io.ReadSeeker, outDir, fileName string, selectedPages []string, n int, cut *model.Cut, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func NDownFile(inFile, outDir, outFile string, selectedPages []string, n int, cut *model.Cut, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func normalizeCut(cut *model.Cut) { _ = "STUB: not implemented"; return }

func writeCutPage(ctxSrc *model.Context, pageNr int, outDir, fileName string, cut *model.Cut, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func Cut(rs io.ReadSeeker, outDir, fileName string, selectedPages []string, cut *model.Cut, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func CutFile(inFile, outDir, outFile string, selectedPages []string, cut *model.Cut, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
