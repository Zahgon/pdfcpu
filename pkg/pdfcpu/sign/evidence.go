package sign

import (
	"crypto/x509"
	"encoding/asn1"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/pkcs7"
)

type timestampKind uint8

const (
	timestampKindSignature timestampKind = iota
	timestampKindDocument
)

type timestampTokenInfo struct {
	Version                 int
	Policy                  asn1.ObjectIdentifier
	MessageImprintAlgorithm asn1.ObjectIdentifier
	MessageImprint          []byte
	SerialNumber            asn1.RawValue
	GeneratedAt             time.Time
	Accuracy                asn1.RawValue
	Ordering                bool
	Nonce                   asn1.RawValue
	TSA                     asn1.RawValue
	Extensions              asn1.RawValue
}

type timestampEvidence struct {
	Kind                  timestampKind
	SigningTime           time.Time
	Present               bool
	AssessmentScope       model.AssessmentScope
	RawToken              []byte
	CMS                   *pkcs7.PKCS7
	SourceSigner          pkcs7.SignerInfo
	TokenInfo             *timestampTokenInfo
	TokenInfoErr          error
	SignedData            []byte
	DigestVerified        bool
	SignatureVerified     bool
	CorrectProfile        bool
	LocalTSAPathValidated bool
	Err                   error
}

func isCryptographicallyAuthenticatedDocumentTimestampEvidence(evidence timestampEvidence) bool {
	_ = "STUB: not implemented"
	return false
}

func isLocallyValidatedDocumentTimestampEvidence(evidence timestampEvidence) bool {
	_ = "STUB: not implemented"
	return false
}

type certificateAssessment struct {
	Certificate           *model.CertificateDetails
	CertificatePathStatus int
	Problems              []string
	Reason                model.SignatureReason
}

type localSignatureAssessment struct {
	SignersProcessed       int
	SignatureAuthenticated bool
	DigestVerified         bool
	ProfileValidated       bool
	CertificateIdentified  bool
	PathValidated          bool
	RevocationGood         bool
}

func (a localSignatureAssessment) complete() bool { _ = "STUB: not implemented"; return false }

func (a *localSignatureAssessment) applyCertificateAssessment(assessment certificateAssessment) {
	_ = "STUB: not implemented"
	return
}

func (a *localSignatureAssessment) merge(other localSignatureAssessment) {
	_ = "STUB: not implemented"
	return
}

func completedLocalSignatureAssessment() localSignatureAssessment {
	_ = "STUB: not implemented"
	return *new(localSignatureAssessment)
}

type dssEvidence struct {
	Certificates  []*x509.Certificate
	CRLs          [][]byte
	OCSPResponses [][]byte
	Supported     bool
}
