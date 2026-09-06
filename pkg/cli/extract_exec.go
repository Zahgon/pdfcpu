package cli

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

const extractPagesOperation = "extract pages"

func validateExtractionCommand(cmd *Command, operation string) error {
	_ = "STUB: not implemented"
	return nil
}

func reportUnsupportedResourceSkips(err error) error { _ = "STUB: not implemented"; return nil }

func writeExtractedPageToStdout(ctx *model.Context, pageNr int, w io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func extractSelectedPageToStdout(rs io.ReadSeeker, w io.Writer, cmd *Command) error {
	_ = "STUB: not implemented"
	return nil
}

func extractPageToStdout(cmd *Command) error { _ = "STUB: not implemented"; return nil }

func ExtractImages(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ExtractFonts(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ExtractPages(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ExtractContent(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ExtractMetadata(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
