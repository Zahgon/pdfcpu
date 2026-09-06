package cli

import (
	"io"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func ImportImages(cmd *Command) (result []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateImportImagesCommand(cmd *Command) error { _ = "STUB: not implemented"; return nil }

func hasStdinImage(inFiles []string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func readImportImageStdin(r io.Reader, imageIndex int, maxStreamBytes int64) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func importImageReader(fn string, imageIndex int, maxStreamBytes int64) (io.Reader, io.Closer, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), *new(io.Closer), nil
}

type importImageInputCloser struct {
	closer     io.Closer
	imageIndex int
	fileName   string
}

func importImageReaders(inFiles []string, maxStreamBytes int64) ([]io.Reader, []importImageInputCloser, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func closeImportImageInputs(closers []importImageInputCloser) error {
	_ = "STUB: not implemented"
	return nil
}

func importImagesDestination(outFile string) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createImportImagesStreamOutput(outFile string, replace bool) (*os.File, string, string, error) {
	_ = "STUB: not implemented"
	return nil, "", "", nil
}

func importImagesToFile(
	outFile string,
	readers []io.Reader,
	closers []importImageInputCloser,
	imp *pdfcpu.Import,
	conf *model.Configuration,
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func CreateCheatSheetsFonts(cmd *Command) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateFontsCommand(cmd *Command, expectedMode model.CommandMode) error {
	_ = "STUB: not implemented"
	return nil
}

func ListFonts(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func InstallFonts(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

var (
	openListImagesInput  = os.Open
	closeListImagesInput = (*os.File).Close
)

func listImagesFile(inFile string, selectedPages []string, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateListImagesInputs(inFiles []string) error { _ = "STUB: not implemented"; return nil }

func ListImagesFile(inFiles []string, selectedPages []string, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateListImagesCommand(cmd *Command) error { _ = "STUB: not implemented"; return nil }

func ListImages(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func updateImageParams(cmd *Command) (objNr, pageNr int, id string) {
	_ = "STUB: not implemented"
	return 0, 0, ""
}

func validateUpdateImagesCommand(cmd *Command) error { _ = "STUB: not implemented"; return nil }

func updateImagesInOut(cmd *Command, objNr, pageNr int, id string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UpdateImages(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func listAttachments(rs io.ReadSeeker, conf *model.Configuration, withDesc, sorted bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var closeListAttachmentsInput = (*os.File).Close

func ListAttachmentsFile(inFile string, conf *model.Configuration) (ss []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListAttachmentsCompactFile(inFile string, conf *model.Configuration) (ss []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateListAttachmentsCommand(cmd *Command) error { _ = "STUB: not implemented"; return nil }

func validateMutateAttachmentsCommand(cmd *Command, operation string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateExtractAttachmentsCommand(cmd *Command) error { _ = "STUB: not implemented"; return nil }

func ListAttachments(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func AddAttachments(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func RemoveAttachments(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ExtractAttachments(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

var closeListKeywordsInput = (*os.File).Close

func ListKeywordsFile(inFile string, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListKeywords(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func runKeywordStreamOperation(inFile, outFile, op string, fn func(io.ReadSeeker, io.Writer) error) error {
	_ = "STUB: not implemented"
	return nil
}

func validateKeywordValues(keywords []string, op string) error {
	_ = "STUB: not implemented"
	return nil
}

func AddKeywords(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func RemoveKeywords(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func renderProperties(properties map[string]string) []string { _ = "STUB: not implemented"; return nil }

func listProperties(rs io.ReadSeeker, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var closeListPropertiesInput = (*os.File).Close

func ListPropertiesFile(inFile string, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListProperties(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func validatePropertyMap(properties map[string]string) error { _ = "STUB: not implemented"; return nil }

func validatePropertyNames(properties []string) error { _ = "STUB: not implemented"; return nil }

func runPropertyStreamOperation(inFile, outFile, op string, fn func(io.ReadSeeker, io.Writer) error) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePropertyCommand(cmd *Command, operation string) error {
	_ = "STUB: not implemented"
	return nil
}

func AddProperties(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func RemoveProperties(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
