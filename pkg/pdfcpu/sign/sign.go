package sign

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"io"
	"time"

	"github.com/hhrutter/pkcs7"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

const CertifiedSigPermsNotSupported = "Certified signature detected. Permission validation not supported."

func validateCertChains(
	chains [][]*x509.Certificate,
	rootCerts *x509.CertPool,
	signer *model.Signer,
	signingTime *time.Time,
	crls [][]byte,
	ocsps [][]byte,
	result *model.SignatureValidationResult,
	conf *model.Configuration) {
	_ = "STUB: not implemented"
	return
}

func setupCertDetails(
	cert *x509.Certificate,
	certDetails *model.CertificateDetails,
	signer *model.Signer,
	signingTime *time.Time,
	result *model.SignatureValidationResult,
	i int) bool {
	_ = "STUB: not implemented"
	return false
}

func getKeySize(cert *x509.Certificate, signer *model.Signer, certDetails *model.CertificateDetails, result *model.SignatureValidationResult) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func setTrustStatus(certDetails *model.CertificateDetails, result *model.SignatureValidationResult) {
	_ = "STUB: not implemented"
	return
}

func signedData(ra io.ReaderAt, sigDict types.Dict) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
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

func qualifiedCertificate(cert *x509.Certificate) bool { _ = "STUB: not implemented"; return false }

func certChain(cert *x509.Certificate, certs []*x509.Certificate) []*x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func publicKeySize(cert *x509.Certificate) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func handleCertVerifyErr(err error, cert *x509.Certificate, signer *model.Signer, result *model.SignatureValidationResult) {
	_ = "STUB: not implemented"
	return
}

func certInfo(cert *x509.Certificate) string { _ = "STUB: not implemented"; return "" }

func processDSS(ctx *model.Context, signer *model.Signer) ([]*x509.Certificate, [][]byte, [][]byte, bool) {
	_ = "STUB: not implemented"
	return nil, nil, nil, false
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
