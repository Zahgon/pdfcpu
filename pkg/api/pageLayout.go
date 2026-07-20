package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func PageLayout(rs io.ReadSeeker, conf *model.Configuration) (pl *model.PageLayout, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PageLayoutFile(inFile string, conf *model.Configuration) (*model.PageLayout, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListPageLayout(rs io.ReadSeeker, conf *model.Configuration) (ss []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListPageLayoutFile(inFile string, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetPageLayout(rs io.ReadSeeker, w io.Writer, val model.PageLayout, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func SetPageLayoutFile(inFile, outFile string, val model.PageLayout, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ResetPageLayout(rs io.ReadSeeker, w io.Writer, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ResetPageLayoutFile(inFile, outFile string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
