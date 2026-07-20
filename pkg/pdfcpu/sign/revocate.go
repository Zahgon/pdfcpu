package sign

import (
	"crypto/x509"
	"encoding/asn1"
	"net/http"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"golang.org/x/crypto/ocsp"
)

const (
	crlReasonUnspecified = iota
	crlReasonKeyCompromise
	crlReasonCACompromise
	crlReasonAffiliationChanged
	crlReasonSuperseded
	crlReasonCessationOfOperation
	crlReasonCertificateHold
	_
	crlReasonRemoveFromCRL
	crlReasonPrivilegeWithdrawn
	crlReasonAACompromise
)

type OtherRevInfo struct {
	Type  asn1.ObjectIdentifier
	Value []byte
}

type RevocationInfoArchival struct {
	CRLs         []asn1.RawValue `asn1:"optional,explicit,tag:0"`
	OCSPs        []asn1.RawValue `asn1:"optional,explicit,tag:1"`
	OtherRevInfo []OtherRevInfo  `asn1:"optional,explicit,tag:2"`
}

func checkRevocation(
	cert, issuer *x509.Certificate,
	rootCerts *x509.CertPool,
	signer *model.Signer,
	certDetails *model.CertificateDetails,
	signingTime *time.Time,
	crls [][]byte,
	ocsps [][]byte,
	result *model.SignatureValidationResult,
	conf *model.Configuration) {
	_ = "STUB: not implemented"
	return
}

func checkCertificateRevocation(
	cert, issuer *x509.Certificate,
	rootCerts *x509.CertPool,
	signer *model.Signer,
	signingTime *time.Time,
	crls [][]byte,
	ocsps [][]byte,
	conf *model.Configuration) (*model.RevocationDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkCertAgainstCRL(
	cert, issuer *x509.Certificate,
	rootCerts *x509.CertPool,
	signingTime *time.Time,
	crls [][]byte,
	conf *model.Configuration) (*model.RevocationDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func processArchivedCRLs(cert *x509.Certificate, signingTime time.Time, crls [][]byte) (*model.RevocationDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func processCurrentCRLs(cert *x509.Certificate, conf *model.Configuration) (*model.RevocationDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkCertViaOCSP(
	cert, issuer *x509.Certificate,
	rootCerts *x509.CertPool,
	signingTime *time.Time,
	ocsps [][]byte,
	conf *model.Configuration) (*model.RevocationDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func processArchivedOCSPResponses(
	cert, issuer *x509.Certificate,
	rootCerts *x509.CertPool,
	signingTime time.Time,
	ocsps [][]byte,
	client *http.Client) (*model.RevocationDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkArchivedOCSPResponse(resp *ocsp.Response, signingTime time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func processCurrentOCSPResponses(
	cert, issuer *x509.Certificate,
	rootCerts *x509.CertPool,
	client *http.Client) (*model.RevocationDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkCurrentOCSPResponse(resp *ocsp.Response) error { _ = "STUB: not implemented"; return nil }

func checkResponderCert(resp *ocsp.Response, rootCerts *x509.CertPool) error {
	_ = "STUB: not implemented"
	return nil
}

func findOCSPResponderCert(resp *ocsp.Response, rootCerts *x509.CertPool) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func hasNoCheckExtension(cert *x509.Certificate) bool { _ = "STUB: not implemented"; return false }

func getIssuerCertificate(cert *x509.Certificate, pool *x509.CertPool, client *http.Client) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
