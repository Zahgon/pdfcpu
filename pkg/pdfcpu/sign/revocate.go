package sign

import (
	"crypto/x509"
	"encoding/asn1"
	"errors"
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

const defaultOCSPResponseMaxAge = 24 * time.Hour

const maxArchivedOCSPFailureCauses = 16

const maxRevocationResponseBytes int64 = 64 << 20

var errRevocationResponseTooLarge = errors.New("revocation response exceeds size limit")

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
	crls [][]byte,
	ocsps [][]byte,
	conf *model.Configuration) (*model.RevocationDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkCertAgainstCRL(
	cert, issuer *x509.Certificate,
	rootCerts *x509.CertPool,
	crls [][]byte,
	conf *model.Configuration) (*model.RevocationDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func processArchivedCRLs(
	_ *x509.Certificate,
	issuer *x509.Certificate,
	crls [][]byte,
) (*model.RevocationDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func processCurrentCRLs(
	cert, issuer *x509.Certificate,
	client *http.Client,
) (*model.RevocationDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func assessCRL(
	crl *x509.RevocationList,
	issuer *x509.Certificate,
	source model.RevocationEvidenceSource,
	applicable bool,
) *model.CRLEvidence {
	_ = "STUB: not implemented"
	return nil
}

func unsupportedCRLExtensions(crl *x509.RevocationList) []asn1.ObjectIdentifier {
	_ = "STUB: not implemented"
	return nil
}

func authenticatedApplicableCRL(evidence *model.CRLEvidence) bool {
	_ = "STUB: not implemented"
	return false
}

func crlRevocationEntries(crl *x509.RevocationList) []model.CRLRevocationEntry {
	_ = "STUB: not implemented"
	return nil
}

func failedCRLEvidence(
	source model.RevocationEvidenceSource,
	index int,
	location string,
	err error,
) *model.CRLEvidence {
	_ = "STUB: not implemented"
	return nil
}

func crlRevocationDetails(
	status int,
	reason string,
	evidence *model.CRLEvidence,
	observations []*model.CRLEvidence,
) *model.RevocationDetails {
	_ = "STUB: not implemented"
	return nil
}

func lastCRLEvidence(observations []*model.CRLEvidence) *model.CRLEvidence {
	_ = "STUB: not implemented"
	return nil
}

func revocationConcluded(details *model.RevocationDetails) bool {
	_ = "STUB: not implemented"
	return false
}

func mergeRevocationDetails(
	primary, observations *model.RevocationDetails,
) *model.RevocationDetails {
	_ = "STUB: not implemented"
	return nil
}

func boolStatus(ok bool) int { _ = "STUB: not implemented"; return 0 }

func readAndCloseResponseWithLimit(resp *http.Response, maxBytes int64) (bb []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readAndCloseResponse(resp *http.Response) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkCertViaOCSP(
	cert, issuer *x509.Certificate,
	_ *x509.CertPool,
	ocsps [][]byte,
	conf *model.Configuration) (*model.RevocationDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func processArchivedOCSPResponses(
	cert, issuer *x509.Certificate,
	ocsps [][]byte) (*model.RevocationDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ocspCandidate struct {
	details  *model.RevocationDetails
	evidence *model.OCSPEvidence
}

func processArchivedOCSPResponse(
	cert, issuer *x509.Certificate,
	bb []byte,
) (*ocspCandidate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func observeArchivedOCSPResponder(
	resp *ocsp.Response,
	issuer *x509.Certificate,
	evidence *model.OCSPEvidence,
) error {
	_ = "STUB: not implemented"
	return nil
}

func appendIndexedFailure(failures []error, location string, err error) []error {
	_ = "STUB: not implemented"
	return nil
}

func joinFailures(failures []error) error { _ = "STUB: not implemented"; return nil }

func ocspRevocationDetails(
	status int,
	evidence *model.OCSPEvidence,
) (*model.RevocationDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ocspCertificateStatus(status int) int { _ = "STUB: not implemented"; return 0 }

func authenticatedOCSPEvidence(
	source model.RevocationEvidenceSource,
	responder model.OCSPResponder,
	resp *ocsp.Response,
) *model.OCSPEvidence {
	_ = "STUB: not implemented"
	return nil
}

func failedOCSPEvidence(
	source model.RevocationEvidenceSource,
	index int,
	location string,
	err error,
) *model.OCSPEvidence {
	_ = "STUB: not implemented"
	return nil
}

func failedOCSPCandidateEvidence(
	candidate *ocspCandidate,
	source model.RevocationEvidenceSource,
	index int,
	location string,
	err error,
) *model.OCSPEvidence {
	_ = "STUB: not implemented"
	return nil
}

func observedOCSPCandidate(
	source model.RevocationEvidenceSource,
	resp *ocsp.Response,
) *ocspCandidate {
	_ = "STUB: not implemented"
	return nil
}

func inconclusiveOCSPDetails(
	reason string,
	observations []*model.OCSPEvidence,
) *model.RevocationDetails {
	_ = "STUB: not implemented"
	return nil
}

func concludeOCSPCandidates(
	candidates []*ocspCandidate,
	observations []*model.OCSPEvidence,
	defaultReason string,
) *model.RevocationDetails {
	_ = "STUB: not implemented"
	return nil
}

func inconclusiveOCSPReason(candidates []*ocspCandidate, defaultReason string) string {
	_ = "STUB: not implemented"
	return ""
}

func processCurrentOCSPResponses(
	cert, issuer *x509.Certificate,
	client *http.Client) (*model.RevocationDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func processCurrentOCSPResponse(
	cert, issuer *x509.Certificate,
	client *http.Client,
	request []byte,
	ocspURL string,
	now time.Time,
) (*ocspCandidate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkCurrentOCSPResponse(resp *ocsp.Response, now time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func checkResponderCert(
	resp *ocsp.Response,
	issuer *x509.Certificate,
	validationTime time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}

func authenticateOCSPResponder(
	resp *ocsp.Response,
	issuer *x509.Certificate,
	validationTime time.Time,
) (model.OCSPResponder, error) {
	_ = "STUB: not implemented"
	return *new(model.OCSPResponder), nil
}

func resolveOCSPResponderCertificate(
	resp *ocsp.Response,
	issuer *x509.Certificate,
) (*x509.Certificate, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func verifyResponderCertificate(
	cert, issuer *x509.Certificate,
	validationTime time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}

func hasNoCheckExtension(cert *x509.Certificate) bool { _ = "STUB: not implemented"; return false }
