package cli

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

type signatureValidationFileOperation func(
	string,
	bool,
	bool,
	*model.Configuration,
) ([]string, error)

func validateCertificateCommand(cmd *Command, operation string) error {
	_ = "STUB: not implemented"
	return nil
}

func ListCertificatesAll(json bool, _ *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListCertificates(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ImportCertificates(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func InspectCertificates(cmd *Command) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ValidateSignatures(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func validateSignatures(
	cmd *Command,
	operation signatureValidationFileOperation,
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RemoveSignatures(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
