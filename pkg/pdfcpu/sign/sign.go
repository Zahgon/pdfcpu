package sign

import (
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/pkcs7"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

const (
	CertifiedSigPermsNotSupported = "Certified signature detected. Permission validation not supported."

	certImportHint = "import missing certificates into pdfcpu's local certificate store with \"pdfcpu certificates import <file>\""
)

func parseCertificate(bb []byte) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setResultReason(result *model.SignatureValidationResult, reason model.SignatureReason) {
	_ = "STUB: not implemented"
	return
}

func markDocumentUnmodified(result *model.SignatureValidationResult) {
	_ = "STUB: not implemented"
	return
}

func markInvalidEvidence(
	result *model.SignatureValidationResult,
	reason model.SignatureReason,
	docModified int,
) {
	_ = "STUB: not implemented"
	return
}

func markUnsupportedEvidence(result *model.SignatureValidationResult) {
	_ = "STUB: not implemented"
	return
}

func markMalformedEvidence(result *model.SignatureValidationResult) {
	_ = "STUB: not implemented"
	return
}

func markCertificateInvalidEvidence(result *model.SignatureValidationResult) {
	_ = "STUB: not implemented"
	return
}

func finalizeLocalSignatureResult(
	result *model.SignatureValidationResult,
	assessment localSignatureAssessment,
) bool {
	_ = "STUB: not implemented"
	return false
}

func assessCertificateEvidence(
	chains [][]*x509.Certificate,
	pathResolved bool,
	rootCerts *x509.CertPool,
	crls, ocsps [][]byte,
	reason model.SignatureReason,
	conf *model.Configuration,
) (certificateAssessment, error) {
	_ = "STUB: not implemented"
	return *new(certificateAssessment), nil
}

func applyCertificateAssessment(
	assessment certificateAssessment,
	signer *model.Signer,
	result *model.SignatureValidationResult,
) {
	_ = "STUB: not implemented"
	return
}

func validateCertChains(
	chains [][]*x509.Certificate,
	pathResolved bool,
	rootCerts *x509.CertPool,
	signer *model.Signer,
	crls [][]byte,
	ocsps [][]byte,
	result *model.SignatureValidationResult,
	conf *model.Configuration,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCertificateInChain(
	cert, issuer *x509.Certificate,
	certIndex int,
	pathResolved bool,
	rootCerts *x509.CertPool,
	signer *model.Signer,
	crls, ocsps [][]byte,
	result *model.SignatureValidationResult,
	conf *model.Configuration,
) (*model.CertificateDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func reportMissingCertificate(
	certDetails *model.CertificateDetails,
	signer *model.Signer,
	result *model.SignatureValidationResult,
	certIndex int,
) {
	_ = "STUB: not implemented"
	return
}

func reportSelfSignedCertificateError(
	cert *x509.Certificate,
	certDetails *model.CertificateDetails,
	signer *model.Signer,
	result *model.SignatureValidationResult,
	certIndex int,
	err error,
) {
	_ = "STUB: not implemented"
	return
}

func appendCertificateDetails(
	signer *model.Signer,
	previous, current *model.CertificateDetails,
) *model.CertificateDetails {
	_ = "STUB: not implemented"
	return nil
}

func certificateIssuer(chain []*x509.Certificate, certIndex int) *x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func setupCertDetails(
	cert *x509.Certificate,
	certDetails *model.CertificateDetails,
	signer *model.Signer,
	result *model.SignatureValidationResult,
	i int,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func getKeySize(
	cert *x509.Certificate,
	signer *model.Signer,
	certDetails *model.CertificateDetails,
	result *model.SignatureValidationResult,
	certIndex int,
) (int, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func getKeySizeWith(
	cert *x509.Certificate,
	signer *model.Signer,
	certDetails *model.CertificateDetails,
	result *model.SignatureValidationResult,
	certIndex int,
	inspect func(*x509.Certificate) (int, error),
) (int, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func setLocalCertificatePathStatus(certDetails *model.CertificateDetails, pathResolved bool) {
	_ = "STUB: not implemented"
	return
}

func setCertificatePathConclusion(
	certDetails *model.CertificateDetails,
	status int,
	reason string,
	method model.CertificatePathMethod,
) {
	_ = "STUB: not implemented"
	return
}

func signedData(ra io.ReaderAt, sigDict types.Dict) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateContentsGap(ra io.ReaderAt, sigDict types.Dict, values [4]int64) error {
	_ = "STUB: not implemented"
	return nil
}

func contentsGapMatches(gap []byte, contents string) bool { _ = "STUB: not implemented"; return false }

func toUpperHex(b byte) byte { _ = "STUB: not implemented"; return 0 }

func byteRangeValues(arr types.Array) ([4]int64, error) {
	_ = "STUB: not implemented"
	return [4]int64{}, nil
}

func byteRangeEnd(off, size int64) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func validateByteRange(values [4]int64) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func copyByteRange(w io.Writer, ra io.ReaderAt, off, size int64) error {
	_ = "STUB: not implemented"
	return nil
}

func bytesForByteRange(ra io.ReaderAt, arr types.Array) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isSelfSigned(cert *x509.Certificate) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func comparePKIXName(a, b pkix.Name) bool { _ = "STUB: not implemented"; return false }

func certUsage(cert *x509.Certificate) string { _ = "STUB: not implemented"; return "" }

func hasRecognizedQualifiedCertificatePolicy(cert *x509.Certificate) bool {
	_ = "STUB: not implemented"
	return false
}

func certChain(cert *x509.Certificate, certs []*x509.Certificate) []*x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func publicKeySize(cert *x509.Certificate) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func rsaPublicKeySize(pubKey *rsa.PublicKey) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func ecdsaPublicKeySize(pubKey *ecdsa.PublicKey) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ed25519PublicKeySize(pubKey ed25519.PublicKey) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func dsaPublicKeySize(pubKey *dsa.PublicKey) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func handleCertVerifyErr(err error, cert *x509.Certificate, signer *model.Signer, result *model.SignatureValidationResult) {
	_ = "STUB: not implemented"
	return
}

func handleCertParseErr(err error, result *model.SignatureValidationResult) {
	_ = "STUB: not implemented"
	return
}

func certInfo(cert *x509.Certificate) string { _ = "STUB: not implemented"; return "" }

func processDSS(ctx *model.Context, signer *model.Signer) dssEvidence {
	_ = "STUB: not implemented"
	return *new(dssEvidence)
}

func extractCertsFromDSS(ctx *model.Context) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeCerts(certLists ...[]*x509.Certificate) []*x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func extractCRLsFromDSS(ctx *model.Context) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractOCSPsFromDSS(ctx *model.Context) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateP7(sigDict types.Dict, result *model.SignatureValidationResult) *pkcs7.PKCS7 {
	_ = "STUB: not implemented"
	return nil
}
