package api

import (
	"io"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func addAttachmentsOperation(coll bool) string { _ = "STUB: not implemented"; return "" }

func addAttachmentsCommandMode(coll bool) model.CommandMode {
	_ = "STUB: not implemented"
	return *new(model.CommandMode)
}

func Attachments(rs io.ReadSeeker, conf *model.Configuration) (aa []model.Attachment, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addAttachment(ctx *model.Context, spec string, coll bool, op string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AddAttachments(rs io.ReadSeeker, w io.Writer, files []string, coll bool, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type attachmentMutation func(io.ReadSeeker, io.Writer) error

func mutateAttachmentsFile(inFile, outFile, op string, mutate attachmentMutation) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AddAttachmentsFile(inFile, outFile string, files []string, coll bool, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func RemoveAttachments(rs io.ReadSeeker, w io.Writer, files []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemoveAttachmentsFile(inFile, outFile string, files []string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func ExtractAttachmentsRaw(rs io.ReadSeeker, outDir string, fileNames []string, conf *model.Configuration) (aa []model.Attachment, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func attachmentOutputPath(outDir string, i int, a model.Attachment) string {
	_ = "STUB: not implemented"
	return ""
}

func writeAttachmentToPath(fileName string, a model.Attachment) error {
	_ = "STUB: not implemented"
	return nil
}

func writeAttachment(outDir string, i int, a model.Attachment) error {
	_ = "STUB: not implemented"
	return nil
}

func attachmentOutputPaths(outDir string, aa []model.Attachment) []string {
	_ = "STUB: not implemented"
	return nil
}

type attachmentOutputReservation struct {
	file *os.File
	path string
	id   string
}

func attachmentReservationToken() (string, error) { _ = "STUB: not implemented"; return "", nil }

func attachmentReservationPath(fileName, token string) string { _ = "STUB: not implemented"; return "" }

func attachmentReservationConflictID(path string, rr []attachmentOutputReservation) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func reserveAttachmentOutputs(
	paths []string,
	aa []model.Attachment,
) ([]attachmentOutputReservation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func releaseAttachmentOutputReservations(rr []attachmentOutputReservation) error {
	_ = "STUB: not implemented"
	return nil
}

func writeAttachments(outDir string, aa []model.Attachment) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ExtractAttachments(rs io.ReadSeeker, outDir string, fileNames []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ExtractAttachmentsFile(inFile, outDir string, files []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
