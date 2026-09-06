package api

import (
	"io"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func validPageMode(pm model.PageMode) bool { _ = "STUB: not implemented"; return false }

func invalidPageModeError(pm model.PageMode) error { _ = "STUB: not implemented"; return nil }

func closePageModeInput(err error, f *os.File, context string) error {
	_ = "STUB: not implemented"
	return nil
}

func PageMode(rs io.ReadSeeker, conf *model.Configuration) (pm *model.PageMode, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PageModeFile(inFile string, conf *model.Configuration) (pm *model.PageMode, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListPageMode(rs io.ReadSeeker, conf *model.Configuration) (ss []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListPageModeFile(inFile string, conf *model.Configuration) (ss []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetPageMode(rs io.ReadSeeker, w io.Writer, val model.PageMode, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func SetPageModeFile(inFile, outFile string, val model.PageMode, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ResetPageMode(rs io.ReadSeeker, w io.Writer, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ResetPageModeFile(inFile, outFile string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
