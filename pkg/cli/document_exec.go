package cli

import (
	"io"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func Validate(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func Optimize(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func mergeStdinCount(inFiles []string) int { _ = "STUB: not implemented"; return 0 }

func mergeReader(fn string) (io.ReadSeeker, *os.File, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadSeeker), nil, nil
}

func closeMergeFiles(files []*os.File) { _ = "STUB: not implemented"; return }

func mergeReaders(inFiles []string) ([]io.ReadSeeker, []*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func mergeOutput(outFile string) (io.Writer, func(), error) {
	_ = "STUB: not implemented"
	return *new(io.Writer), nil, nil
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

func ListInfo(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func Dump(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func Create(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
