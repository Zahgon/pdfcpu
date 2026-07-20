package sign

import (
	"crypto/x509"
	"io"
	"time"

	"github.com/hhrutter/pkcs7"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func ValidatePKCS7Signatures(
	ra io.ReaderAt,
	sigDict types.Dict,
	certified bool,
	authoritative bool,
	validateAll bool,
	perms int,
	rootCerts *x509.CertPool,
	result *model.SignatureValidationResult,
	ctx *model.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func finalizePKCS7Result(result *model.SignatureValidationResult) {
	_ = "STUB: not implemented"
	return
}

func p7(sigDict types.Dict) (*pkcs7.PKCS7, error) { _ = "STUB: not implemented"; return nil, nil }

func verifyP7Signer(
	p7Signer pkcs7.SignerInfo,
	p7Certs []*x509.Certificate,
	rootCerts *x509.CertPool,
	p7Content []byte,
	data []byte,
	detached bool,
	certified bool,
	authoritative bool,
	perms, i int,
	result *model.SignatureValidationResult,
	ctx *model.Context) {
	_ = "STUB: not implemented"
	return
}

func checkPerms(signer *model.Signer, result *model.SignatureValidationResult) {
	_ = "STUB: not implemented"
	return
}

func checkP7Digest(
	p7Signer pkcs7.SignerInfo,
	p7Content,
	data []byte, detached bool,
	signer *model.Signer,
	result *model.SignatureValidationResult) bool {
	_ = "STUB: not implemented"
	return false
}

func verifyP7Digest(p7Signer pkcs7.SignerInfo, p7Content []byte, data []byte, detached bool) (model.SignatureReason, error) {
	_ = "STUB: not implemented"
	return *new(model.SignatureReason), nil
}

func checkTimestampToken(
	detached bool,
	p7Signer pkcs7.SignerInfo,
	rootCerts *x509.CertPool,
	ctx *model.Context,
	signer *model.Signer,
	result *model.SignatureValidationResult) (signingTime *time.Time) {
	_ = "STUB: not implemented"
	return nil
}

func handleDSS(certs *[]*x509.Certificate, crls *[][]byte, ocsps *[][]byte, ctx *model.Context, signer *model.Signer, detached bool) {
	_ = "STUB: not implemented"
	return
}

func verifyP7Signature(p7Signer pkcs7.SignerInfo, cert *x509.Certificate, p7Content []byte, detached bool) error {
	_ = "STUB: not implemented"
	return nil
}

func handleClaimedSigningTime(signerInfo pkcs7.SignerInfo, signer *model.Signer, result *model.SignatureValidationResult) *time.Time {
	_ = "STUB: not implemented"
	return nil
}

func timestampToken(p7Signer pkcs7.SignerInfo, _ *x509.CertPool) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func locateTimestampToken(signerInfo pkcs7.SignerInfo) []byte {
	_ = "STUB: not implemented"
	return nil
}

func extractTimestampTokenTime(data []byte) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func handleArchivedRevocationInfo(p7Signer pkcs7.SignerInfo, signer *model.Signer) (crls [][]byte, ocsps [][]byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildP7CertChains(
	first bool,
	cert *x509.Certificate,
	certs []*x509.Certificate,
	rootCerts *x509.CertPool,
	signer *model.Signer,
	signingTime *time.Time,
	result *model.SignatureValidationResult) [][]*x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func handleTimestampToken(p7Signer pkcs7.SignerInfo, rootCerts *x509.CertPool, signer *model.Signer, result *model.SignatureValidationResult) *time.Time {
	_ = "STUB: not implemented"
	return nil
}

func revocationInfoArchival(p7Signer pkcs7.SignerInfo) (*RevocationInfoArchival, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
