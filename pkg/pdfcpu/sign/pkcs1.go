package sign

import (
	"crypto/rsa"
	"crypto/x509"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func ValidateX509RSASHA1Signature(
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

func handleP1VerificationError(
	reason model.SignatureReason,
	err error,
	result *model.SignatureValidationResult,
) error {
	_ = "STUB: not implemented"
	return nil
}

func isFatalByteRangeRead(err error) bool { _ = "STUB: not implemented"; return false }

func p1SigningCertificate(certs []*x509.Certificate) (*x509.Certificate, *rsa.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func validateP1SigningKeyUsage(cert *x509.Certificate) error { _ = "STUB: not implemented"; return nil }

func validateP1RSAPublicKey(key *rsa.PublicKey) error { _ = "STUB: not implemented"; return nil }

func parseP1Certificates(sigDict types.Dict) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func certFromObj(obj types.Object) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func certFromStringLiteral(obj types.StringLiteral) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func certFromHexLiteral(obj types.HexLiteral) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func verifyRSASHA1Signature(ra io.ReaderAt, sigDict types.Dict, rsaPubKey *rsa.PublicKey) (model.SignatureReason, error) {
	_ = "STUB: not implemented"
	return *new(model.SignatureReason), nil
}

func buildP1CertChains(
	cert *x509.Certificate,
	certs []*x509.Certificate,
	rootCerts *x509.CertPool,
	signer *model.Signer,
	result *model.SignatureValidationResult) [][]*x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func p1IntermediatePool(certs []*x509.Certificate) (*x509.CertPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
