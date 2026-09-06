package sign

import (
	"crypto/x509"
	"encoding/asn1"
	"io"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/pkcs7"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type timestampApplication struct {
	signer               *model.Signer
	result               *model.SignatureValidationResult
	setResultSigningTime bool
	problemPrefix        string
}

const embeddedTimestampNotAuthenticated = "embedded timestamp token observed but not fully authenticated"

func applyTimestampEvidence(evidence timestampEvidence, application timestampApplication) {
	_ = "STUB: not implemented"
	return
}

func embeddedSignatureTimestampEvidence(p7Signer pkcs7.SignerInfo) timestampEvidence {
	_ = "STUB: not implemented"
	return *new(timestampEvidence)
}

func populateEmbeddedTimestampInfo(evidence *timestampEvidence) error {
	_ = "STUB: not implemented"
	return nil
}

func documentTimestampEvidence(signingTime time.Time) timestampEvidence {
	_ = "STUB: not implemented"
	return *new(timestampEvidence)
}

func preparedDocumentTimestampEvidence(
	tstInfo *TSTInfo,
	rawToken []byte,
	cms *pkcs7.PKCS7,
	sourceSigner pkcs7.SignerInfo,
	signedData []byte,
) timestampEvidence {
	_ = "STUB: not implemented"
	return *new(timestampEvidence)
}

func structuredTimestampTokenInfo(tstInfo *TSTInfo) *timestampTokenInfo {
	_ = "STUB: not implemented"
	return nil
}

func ValidatePKCS7Signatures(
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

func finalizePKCS7Result(
	result *model.SignatureValidationResult,
	assessment localSignatureAssessment,
) {
	_ = "STUB: not implemented"
	return
}

func p7(sigDict types.Dict) (*pkcs7.PKCS7, error) { _ = "STUB: not implemented"; return nil, nil }

func signatureContents(sigDict types.Dict) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

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
	ctx *model.Context,
) error {
	_ = "STUB: not implemented"
	return nil
}

func verifyP7SignerWithContentType(
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
	ctx *model.Context,
	contentType asn1.ObjectIdentifier,
	localAssessment *localSignatureAssessment,
) error {
	_ = "STUB: not implemented"
	return nil
}

func applyP7ProfileAssessment(
	subFilter string,
	detached bool,
	contentType asn1.ObjectIdentifier,
	p7Signer pkcs7.SignerInfo,
	signerCert *x509.Certificate,
	signer *model.Signer,
	result *model.SignatureValidationResult,
	assessment *localSignatureAssessment,
) {
	_ = "STUB: not implemented"
	return
}

func validateAdobePKCS7Profile(
	subFilter string,
	detached bool,
	contentType asn1.ObjectIdentifier,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCAdESBaselineBProfile(
	detached bool,
	contentType asn1.ObjectIdentifier,
	p7Signer pkcs7.SignerInfo,
	signerCert *x509.Certificate,
) error {
	_ = "STUB: not implemented"
	return nil
}

func classifyCAdESBaselineBProfileError(err error) error { _ = "STUB: not implemented"; return nil }

func reportCAdESBaselineBProfileError(
	err error,
	signer *model.Signer,
	result *model.SignatureValidationResult,
) {
	_ = "STUB: not implemented"
	return
}

func checkPerms(signer *model.Signer, result *model.SignatureValidationResult) {
	_ = "STUB: not implemented"
	return
}

func applyP7DigestEvidence(
	reason model.SignatureReason,
	err error,
	signer *model.Signer,
	result *model.SignatureValidationResult,
) bool {
	_ = "STUB: not implemented"
	return false
}

func verifyP7Digest(p7Signer pkcs7.SignerInfo, p7Content []byte, data []byte, detached bool) (model.SignatureReason, error) {
	_ = "STUB: not implemented"
	return *new(model.SignatureReason), nil
}

func checkTimestampToken(
	p7Signer pkcs7.SignerInfo,
	_ *model.Context,
	signer *model.Signer,
	result *model.SignatureValidationResult,
) {
	_ = "STUB: not implemented"
	return
}

func handleDSS(
	certs *[]*x509.Certificate,
	crls *[][]byte,
	ocsps *[][]byte,
	ctx *model.Context,
	signer *model.Signer,
	result *model.SignatureValidationResult,
	_ bool,
) {
	_ = "STUB: not implemented"
	return
}

func verifyP7Signature(
	p7Signer pkcs7.SignerInfo,
	cert *x509.Certificate,
	p7Content []byte,
	contentType ...asn1.ObjectIdentifier,
) error {
	_ = "STUB: not implemented"
	return nil
}

func reportP7SignatureError(
	err error,
	signer *model.Signer,
	result *model.SignatureValidationResult,
) {
	_ = "STUB: not implemented"
	return
}

func reportSignatureVerificationError(
	phase string,
	err error,
	contentMismatchReason model.SignatureReason,
	signer *model.Signer,
	result *model.SignatureValidationResult,
) {
	_ = "STUB: not implemented"
	return
}

func isUnsupportedP7SignatureError(err error) bool { _ = "STUB: not implemented"; return false }

func isMalformedP7SignatureError(err error) bool { _ = "STUB: not implemented"; return false }

func handleClaimedSigningTime(signerInfo pkcs7.SignerInfo, signer *model.Signer, result *model.SignatureValidationResult) {
	_ = "STUB: not implemented"
	return
}

func parseClaimedSigningTime(signerInfo pkcs7.SignerInfo) (*time.Time, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func locateTimestampToken(signerInfo pkcs7.SignerInfo) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractTimestampTokenTime(data []byte) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func timestampTokenSigningTime(signer pkcs7.SignerInfo) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func parseTimestampTokenSigningTime(bb []byte) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func parseTimestampSigningTime(rawValue asn1.RawValue) (time.Time, error) {
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
	result *model.SignatureValidationResult) [][]*x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func handleTimestampToken(
	p7Signer pkcs7.SignerInfo,
	signer *model.Signer,
	result *model.SignatureValidationResult,
) {
	_ = "STUB: not implemented"
	return
}

func revocationInfoArchival(p7Signer pkcs7.SignerInfo) (*RevocationInfoArchival, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
