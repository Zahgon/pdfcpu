package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func InsertPages(rs io.ReadSeeker, w io.Writer, selectedPages []string, before bool, pageConf *pdfcpu.PageConfiguration, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func InsertPagesFile(inFile, outFile string, selectedPages []string, before bool, pageConf *pdfcpu.PageConfiguration, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemovePages(rs io.ReadSeeker, w io.Writer, selectedPages []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemovePagesFile(inFile, outFile string, selectedPages []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func PageCount(rs io.ReadSeeker, conf *model.Configuration) (count int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func PageCountFile(inFile string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func PageDims(rs io.ReadSeeker, conf *model.Configuration) (pd []types.Dim, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PageDimsFile(inFile string) ([]types.Dim, error) { _ = "STUB: not implemented"; return nil, nil }
