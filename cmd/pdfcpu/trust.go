package main

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/spf13/cobra"
)

type certificatesListOptions struct {
	json bool
}

type signaturesValidateOptions struct {
	all  bool
	full bool
}

func certificatesCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func signaturesRemoveCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func signaturesValidateCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func signaturesCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func resetCertificates(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleListCertificatesCommand(conf *model.Configuration, args []string, opts *certificatesListOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func isCertificateFile(fName string) bool { _ = "STUB: not implemented"; return false }

func certificateFiles(args []string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func certificateFilesForArg(arg string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func expandedCertificateFiles(arg string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func handleInspectCertificatesCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleImportCertificatesCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleValidateSignaturesCommand(conf *model.Configuration, args []string, opts *signaturesValidateOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func handleRemoveSignaturesCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
