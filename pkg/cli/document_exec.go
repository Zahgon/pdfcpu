package cli

import (
	"io"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func validationInputLabel(fn string) string { _ = "STUB: not implemented"; return "" }

func reportValidationProgress(w io.Writer, conf *model.Configuration, fn string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateInput(fn string, conf *model.Configuration, progressOutput io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func reportValidationError(w io.Writer, err error) error { _ = "STUB: not implemented"; return nil }

func validateInputs(inFiles []string, conf *model.Configuration, errorOutput, progressOutput io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func Validate(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func Optimize(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func mergeStdinCount(inFiles []string) int { _ = "STUB: not implemented"; return 0 }

func mergeReader(fn string, source int) (io.ReadSeeker, *os.File, *temporaryInput, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadSeeker), nil, nil, nil
}

func closeMergeInputs(files []*os.File) error { _ = "STUB: not implemented"; return nil }

func mergeReaders(inFiles []string) ([]io.ReadSeeker, []*os.File, *temporaryInput, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func mergeCreateRaw(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func MergeCreate(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func MergeCreateZip(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func MergeAppend(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func Split(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func SplitByPageNr(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func Trim(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func Collect(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func listInfo(rs io.ReadSeeker, inFile string, selectedPages []string, fonts bool, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListInfoFile(inFile string, selectedPages []string, fonts bool, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func jsonInfo(info *pdfcpu.PDFInfo, pages types.IntSet) (map[string]model.PageBoundaries, []types.Dim) {
	_ = "STUB: not implemented"
	return nil, nil
}

func listInfoJSON(rs io.ReadSeeker, inFile string, selectedPages []string, fonts bool, conf *model.Configuration) (*pdfcpu.PDFInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type infoJSONProcessor func(io.ReadSeeker, string, []string, bool, *model.Configuration) (*pdfcpu.PDFInfo, error)

func listInfoFileJSON(
	fn string,
	selectedPages []string,
	fonts bool,
	conf *model.Configuration,
	process infoJSONProcessor,
) (*pdfcpu.PDFInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func listInfoFilesJSON(inFiles []string, selectedPages []string, fonts bool, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func jsonInfoOutput(infos []*pdfcpu.PDFInfo) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListInfoFiles(inFiles []string, selectedPages []string, fonts, json bool, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func listInfoInput(fn string, selectedPages []string, fonts, json bool, conf *model.Configuration) ([]string, *pdfcpu.PDFInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func listInfoReadSeeker(rs io.ReadSeeker, fn string, selectedPages []string, fonts, json bool, conf *model.Configuration) ([]string, *pdfcpu.PDFInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func ListInfo(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func Dump(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func dumpValidationModeHint(mode int) string { _ = "STUB: not implemented"; return "" }

func dumpValidationError(ctx *model.Context, conf *model.Configuration, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func Create(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
