package cli

import (
	"io"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func ImportImages(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func hasStdinImage(inFiles []string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func importImageReader(fn string) (io.Reader, io.Closer, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), *new(io.Closer), nil
}

func importImageReaders(inFiles []string) ([]io.Reader, []io.Closer, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func closeAll(closers []io.Closer) { _ = "STUB: not implemented"; return }

func importImagesDestination(outFile string) (io.ReadSeeker, string, *os.File, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadSeeker), "", nil, nil
}

func importImagesToFile(outFile string, readers []io.Reader, imp *pdfcpu.Import, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func CreateCheatSheetsFonts(cmd *Command) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListFonts(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func InstallFonts(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func listImages(rs io.ReadSeeker, selectedPages []string, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListImagesFile(inFiles []string, selectedPages []string, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListImages(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func updateImageParams(cmd *Command) (objNr, pageNr int, id string) {
	_ = "STUB: not implemented"
	return 0, 0, ""
}

func updateImagesInOut(cmd *Command, objNr, pageNr int, id string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UpdateImages(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func listAttachments(rs io.ReadSeeker, conf *model.Configuration, withDesc, sorted bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListAttachmentsFile(inFile string, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListAttachmentsCompactFile(inFile string, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListAttachments(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func AddAttachments(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func RemoveAttachments(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ExtractAttachments(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ListKeywordsFile(inFile string, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListKeywords(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func AddKeywords(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func RemoveKeywords(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func listProperties(rs io.ReadSeeker, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListPropertiesFile(inFile string, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListProperties(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func AddProperties(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func RemoveProperties(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
