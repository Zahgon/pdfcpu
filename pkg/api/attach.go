package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func Attachments(rs io.ReadSeeker, conf *model.Configuration) (aa []model.Attachment, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AddAttachments(rs io.ReadSeeker, w io.Writer, files []string, coll bool, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AddAttachmentsFile(inFile, outFile string, files []string, coll bool, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemoveAttachments(rs io.ReadSeeker, w io.Writer, files []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemoveAttachmentsFile(inFile, outFile string, files []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ExtractAttachmentsRaw(rs io.ReadSeeker, outDir string, fileNames []string, conf *model.Configuration) (aa []model.Attachment, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExtractAttachments(rs io.ReadSeeker, outDir string, fileNames []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ExtractAttachmentsFile(inFile, outDir string, files []string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}
