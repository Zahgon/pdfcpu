package cli

import (
	"io"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func runContentStreamOperation(inFile, outFile, op string, fn func(io.ReadSeeker, io.Writer) error) error {
	_ = "STUB: not implemented"
	return nil
}

func validateWatermarkCommand(cmd *Command, operation string, requireWatermark bool) error {
	_ = "STUB: not implemented"
	return nil
}

func AddWatermarks(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func RemoveWatermarks(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func listAnnotations(rs io.ReadSeeker, selectedPages []string, json bool, conf *model.Configuration) (int, []string, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func closeListAnnotationsInput(f *os.File, err error) error { _ = "STUB: not implemented"; return nil }

func listAnnotationsFile(inFile string, selectedPages []string, json bool, conf *model.Configuration) (count int, ss []string, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func ListAnnotationsFile(inFile string, selectedPages []string, conf *model.Configuration) (int, []string, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func ListAnnotationsJSONFile(inFile string, selectedPages []string, conf *model.Configuration) (int, []string, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func ListAnnotations(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func RemoveAnnotations(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ListBookmarksFile(inFile string, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListBookmarks(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ExportBookmarks(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ImportBookmarks(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func RemoveBookmarks(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ListPageLayout(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func SetPageLayout(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ResetPageLayout(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ListPageMode(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func SetPageMode(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ResetPageMode(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ListViewerPreferences(cmd *Command) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetViewerPreferences(cmd *Command) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ResetViewerPreferences(cmd *Command) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
