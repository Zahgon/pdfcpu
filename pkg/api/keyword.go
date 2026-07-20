package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func Keywords(rs io.ReadSeeker, conf *model.Configuration) (ss []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AddKeywords(rs io.ReadSeeker, w io.Writer, files []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AddKeywordsFile(inFile, outFile string, files []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemoveKeywords(rs io.ReadSeeker, w io.Writer, keywords []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemoveKeywordsFile(inFile, outFile string, keywords []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
