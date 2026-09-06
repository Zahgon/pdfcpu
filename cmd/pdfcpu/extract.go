package main

import (
	"github.com/pdfcpu/pdfcpu/pkg/cli"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/spf13/cobra"
)

type extractOptions struct {
	mode string
}

func extractCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func extractMode(opts *extractOptions) error { _ = "STUB: not implemented"; return nil }

func extractInputOutput(conf *model.Configuration, args []string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func extractCommandForMode(mode, inFile, outDir string, pages []string, conf *model.Configuration) (*cli.Command, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func handleExtractCommand(conf *model.Configuration, args []string, opts *extractOptions) error {
	_ = "STUB: not implemented"
	return nil
}
