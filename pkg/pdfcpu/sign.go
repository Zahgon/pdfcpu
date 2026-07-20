package pdfcpu

import (
	"crypto/x509"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func ValidateSignatures(ra io.ReaderAt, ctx *model.Context, all bool) ([]*model.SignatureValidationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkForAbortAfterFirst(first bool, svr *model.SignatureValidationResult, ctx *model.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func validateURSignature(sigDict types.Dict, ctx *model.Context, ra io.ReaderAt) (*model.SignatureValidationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateSignature(sig model.Signature, ctx *model.Context, ra io.ReaderAt, first, all bool) (*model.SignatureValidationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sigHandler(subFilter string) func(
	ra io.ReaderAt,
	sigDict types.Dict,
	certified bool,
	authoriative bool,
	validateAll bool,
	perms int,
	rootCerts *x509.CertPool,
	result *model.SignatureValidationResult,
	ctx *model.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func resultDetails(sigDict types.Dict, ctx *model.Context, resultDetails *model.SignatureDetails) error {
	_ = "STUB: not implemented"
	return nil
}

func detectPermissions(sigDict types.Dict, ctx *model.Context) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
