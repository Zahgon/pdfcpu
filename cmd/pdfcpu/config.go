package main

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/spf13/cobra"
)

func configCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func completionCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func paperCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func selectedpagesCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func versionCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func printConfiguration(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func confirmed() bool { _ = "STUB: not implemented"; return false }

func resetConfiguration(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func printPaperSizes(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func printSelectedPages(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func printVersion(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func writeVersionInfo(w io.Writer, configPath string) { _ = "STUB: not implemented"; return }

func formatVersionDate(s string) string { _ = "STUB: not implemented"; return "" }
