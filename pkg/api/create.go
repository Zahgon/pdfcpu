package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func CreatePDFFile(xRefTable *model.XRefTable, outFile string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func Create(rs io.ReadSeeker, rd io.Reader, w io.Writer, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func handleOutFilePDF(inFilePDF, outFilePDF string, tmpFile *string) {
	_ = "STUB: not implemented"
	return
}

func CreateFile(inFilePDF, inFileJSON, outFilePDF string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
