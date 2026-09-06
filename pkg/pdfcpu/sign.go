package pdfcpu

import (
	"crypto/x509"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type signatureValidationHandler func(
	io.ReaderAt,
	types.Dict,
	bool,
	bool,
	bool,
	int,
	*x509.CertPool,
	*model.SignatureValidationResult,
	*model.Context,
) error

func ValidateSignatures(ra io.ReaderAt, ctx *model.Context, all bool) ([]*model.SignatureValidationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func orderedSignatures(signatures map[int]model.Signature) []model.Signature {
	_ = "STUB: not implemented"
	return nil
}

func checkForAbortAfterFirst(first bool, svr *model.SignatureValidationResult, ctx *model.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func validateURSignature(sigDict types.Dict, increment int, ctx *model.Context, ra io.ReaderAt) (*model.SignatureValidationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateSignature(sig model.Signature, ctx *model.Context, ra io.ReaderAt, first, all bool, increment int) (*model.SignatureValidationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func signatureSubFilter(sigDict types.Dict, usageRights bool, result *model.SignatureValidationResult) (string, signatureValidationHandler, bool) {
	_ = "STUB: not implemented"
	return "", *new(signatureValidationHandler), false
}

func recordSubFilterProblem(result *model.SignatureValidationResult, reason model.SignatureReason, problem string) {
	_ = "STUB: not implemented"
	return
}

func applyHistoricalRevisionReporting(increment int, signatureType int, result *model.SignatureValidationResult) {
	_ = "STUB: not implemented"
	return
}

type signedRevisionBoundaryEvidence struct {
	currentFileSize   int64
	signedRevisionEnd int64
	increment         int
	currentRevision   bool
	documentTimestamp bool
}

func recordSignedRevisionBoundaryEvidence(sigDict types.Dict, ctx *model.Context, increment int, documentTimestamp bool, result *model.SignatureValidationResult) bool {
	_ = "STUB: not implemented"
	return false
}

func collectSignedRevisionBoundaryEvidence(sigDict types.Dict, ctx *model.Context, increment int, documentTimestamp bool) (signedRevisionBoundaryEvidence, bool) {
	_ = "STUB: not implemented"
	return *new(signedRevisionBoundaryEvidence), false
}

func sigHandler(subFilter string) signatureValidationHandler {
	_ = "STUB: not implemented"
	return *new(signatureValidationHandler)
}

func validateX509RSASHA1Signature(
	ra io.ReaderAt,
	sigDict types.Dict,
	certified bool,
	authoritative bool,
	validateAll bool,
	perms int,
	rootCerts *x509.CertPool,
	result *model.SignatureValidationResult,
	ctx *model.Context,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePKCS7Signatures(
	ra io.ReaderAt,
	sigDict types.Dict,
	certified bool,
	authoritative bool,
	validateAll bool,
	perms int,
	rootCerts *x509.CertPool,
	result *model.SignatureValidationResult,
	ctx *model.Context,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDTS(
	ra io.ReaderAt,
	sigDict types.Dict,
	certified bool,
	authoritative bool,
	validateAll bool,
	perms int,
	rootCerts *x509.CertPool,
	result *model.SignatureValidationResult,
	ctx *model.Context,
) error {
	_ = "STUB: not implemented"
	return nil
}

func fieldDetails(sigField types.Dict, result *model.SignatureValidationResult) {
	_ = "STUB: not implemented"
	return
}

func signatureDetails(sigDict types.Dict, ctx *model.Context, result *model.SignatureValidationResult) {
	_ = "STUB: not implemented"
	return
}

func detectPermissions(sigDict types.Dict, ctx *model.Context, result *model.SignatureValidationResult) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func signatureReferences(sigDict types.Dict, ctx *model.Context, result *model.SignatureValidationResult) types.Array {
	_ = "STUB: not implemented"
	return *new(types.Array)
}

func signatureReferenceDict(obj types.Object, refIndex int, ctx *model.Context, result *model.SignatureValidationResult) (types.Dict, bool) {
	_ = "STUB: not implemented"
	return *new(types.Dict), false
}

func docMDPPermission(refDict types.Dict, refIndex int, ctx *model.Context, result *model.SignatureValidationResult) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func permissionEntry(params types.Dict, refIndex int, ctx *model.Context, result *model.SignatureValidationResult) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}
