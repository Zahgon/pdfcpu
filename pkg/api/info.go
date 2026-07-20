package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func PDFInfo(rs io.ReadSeeker, fileName string, selectedPages []string, fonts bool, conf *model.Configuration) (info *pdfcpu.PDFInfo, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
