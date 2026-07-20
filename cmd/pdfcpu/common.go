package main

import (
	"github.com/pdfcpu/pdfcpu/pkg/cli"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/spf13/cobra"
)

const (
	selectedPagesWarn = "-selectedPages problem"
	pdfcpuErrPrefix   = "pdfcpu: "
)

func wrapHandler(handler func(*model.Configuration, []string) error) func(*cobra.Command, []string) error {
	_ = "STUB: not implemented"
	return nil
}

func commandError(err error) error { _ = "STUB: not implemented"; return nil }

type prefixStrippedError struct {
	err error
}

func (e prefixStrippedError) Error() string { _ = "STUB: not implemented"; return "" }

func (e prefixStrippedError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func addPasswordFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func addPersistentPasswordFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func addSelectedPagesFlag(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func addPersistentSelectedPagesFlag(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func addRequiredSelectedPagesFlag(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func addUnitFlag(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func modeCompletion(modePrefix string, modes []string) string { _ = "STUB: not implemented"; return "" }

func parseSelectedPages() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func hasPDFExtension(filename string) bool { _ = "STUB: not implemented"; return false }

func ensurePDFExtension(filename string) error { _ = "STUB: not implemented"; return nil }

func hasJSONExtension(filename string) bool { _ = "STUB: not implemented"; return false }

func ensureJSONExtension(filename string) error { _ = "STUB: not implemented"; return nil }

func hasCSVExtension(filename string) bool { _ = "STUB: not implemented"; return false }

func ensureJSONOrCSVExtension(filename string) error { _ = "STUB: not implemented"; return nil }

func ensureOutputFileAvailable(outFile string) error { _ = "STUB: not implemented"; return nil }

func ensureOutputDirEmpty(outDir string) error { _ = "STUB: not implemented"; return nil }

func ensureOutputDirOrFileAvailable(outDir, outFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func inputOutputPDFArgs(conf *model.Configuration, args []string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func optionalOutputPDFArgs(conf *model.Configuration, args []string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func inputPDFArg(conf *model.Configuration, inFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func stdoutForStdin(inFile string) string { _ = "STUB: not implemented"; return "" }

func validateNoEmptyArgs(args []string, name string) error { _ = "STUB: not implemented"; return nil }

func configureDisplayUnit(conf *model.Configuration) error { _ = "STUB: not implemented"; return nil }

func runCommand(cmd *cli.Command) error { _ = "STUB: not implemented"; return nil }
