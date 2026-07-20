package model

import (
	"time"
)

const (
	Unknown = iota
	False
	True
)

const (
	CRL = iota
	OCSP
)

const (
	CertifiedSigPermNone = iota
	CertifiedSigPermNoChangesAllowed
	CertifiedSigPermFillingAndSigningOK
	CertifiedSigPermFillingAnnotatingAndSigningOK
)

const (
	SigTypeForm = iota
	SigTypePage
	SigTypeUR
	SigTypeDTS
)

const SignTSFormat = "2006-01-02 15:04:05 -0700"

type RevocationDetails struct {
	Status int
	Reason string
}

func (rd RevocationDetails) String() string { _ = "STUB: not implemented"; return "" }

type TrustDetails struct {
	Status                                int
	Reason                                string
	SourceObtainedFrom                    string
	AllowSignDocuments                    bool
	AllowCertifyDocuments                 bool
	AllowExecuteDynamicContent            bool
	AllowExecuteJavaScript                bool
	AllowExecutePrivilegedSystemOperation bool
}

func (td TrustDetails) String() string { _ = "STUB: not implemented"; return "" }

type CertificateDetails struct {
	Leaf              bool
	SelfSigned        bool
	Subject           string
	Issuer            string
	SerialNumber      string
	ValidFrom         time.Time
	ValidThru         time.Time
	Expired           bool
	Qualified         bool
	CA                bool
	Usage             string
	Version           int
	SignAlg           string
	KeySize           int
	Revocation        RevocationDetails
	Trust             TrustDetails
	IssuerCertificate *CertificateDetails
}

func (cd CertificateDetails) String() string { _ = "STUB: not implemented"; return "" }

type Signature struct {
	Type          int
	Certified     bool
	Authoritative bool
	Visible       bool
	Signed        bool
	ObjNr         int
	PageNr        int
}

func (sig Signature) String(status SignatureStatus) string { _ = "STUB: not implemented"; return "" }

type SignatureStats struct {
	FormSigned          int
	FormSignedVisible   int
	FormUnsigned        int
	FormUnsignedVisible int
	PageSigned          int
	PageSignedVisible   int
	PageUnsigned        int
	PageUnsignedVisible int
	URSigned            int
	URSignedVisible     int
	URUnsigned          int
	URUnsignedVisible   int
	DTSSigned           int
	DTSSignedVisible    int
	DTSUnsigned         int
	DTSUnsignedVisible  int

	Total int
}

func (sigStats SignatureStats) Counter(svr *SignatureValidationResult) (*int, *int, *int, *int) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

type SignatureStatus int

const (
	SignatureStatusUnknown SignatureStatus = 1 << iota
	SignatureStatusValid
	SignatureStatusInvalid
)

var SignatureStatusStrings = map[SignatureStatus]string{
	SignatureStatusUnknown: "validity of the signature is unknown",
	SignatureStatusValid:   "signature is valid",
	SignatureStatusInvalid: "signature is invalid",
}

func (st SignatureStatus) String() string { _ = "STUB: not implemented"; return "" }

type SignatureReason int

const (
	SignatureReasonUnknown SignatureReason = 1 << iota
	SignatureReasonDocNotModified
	SignatureReasonDocModified
	SignatureReasonSignatureForged
	SignatureReasonSigningTimeInvalid
	SignatureReasonTimestampTokenInvalid
	SignatureReasonCertInvalid
	SignatureReasonCertNotTrusted
	SignatureReasonCertExpired
	SignatureReasonCertRevoked
	SignatureReasonInternal
	SignatureReasonSelfSignedCertErr
)

var SignatureReasonStrings = map[SignatureReason]string{
	SignatureReasonUnknown:               "no reason",
	SignatureReasonDocNotModified:        "document has not been modified",
	SignatureReasonDocModified:           "document has been modified",
	SignatureReasonSignatureForged:       "signer's signature is not authentic",
	SignatureReasonTimestampTokenInvalid: "timestamp token is invalid",
	SignatureReasonCertInvalid:           "signer's certificate is invalid",
	SignatureReasonCertNotTrusted:        "signer's certificate chain is not in the configured local trusted certificate store",
	SignatureReasonCertExpired:           "signer's certificate or one of its parent certificates has expired",
	SignatureReasonCertRevoked:           "signer's certificate or one of its parent certificates has been revoked",
	SignatureReasonInternal:              "internal error",
	SignatureReasonSelfSignedCertErr:     "signer's self signed certificate is not trusted",
}

func (sr SignatureReason) String() string { _ = "STUB: not implemented"; return "" }

type Signer struct {
	Certificate           *CertificateDetails
	CertificatePathStatus int
	HasTimestamp          bool
	Timestamp             time.Time
	LTVEnabled            bool
	PAdES                 string
	Certified             bool
	Authoritative         bool
	Permissions           int
	Problems              []string
}

func (signer *Signer) AddProblem(s string) { _ = "STUB: not implemented"; return }

func permString(i int) string { _ = "STUB: not implemented"; return "" }

func (signer Signer) String(dts bool) string { _ = "STUB: not implemented"; return "" }

type SignatureDetails struct {
	SubFilter      string
	SignerIdentity string
	SignerName     string
	ContactInfo    string
	Location       string
	Reason         string
	SigningTime    time.Time
	FieldName      string
	Signers        []*Signer
}

func (sd *SignatureDetails) AddSigner(s *Signer) { _ = "STUB: not implemented"; return }

func (sd *SignatureDetails) IsETSI_CAdES_detached() bool { _ = "STUB: not implemented"; return false }

func (sd *SignatureDetails) IsETSI_RFC3161() bool { _ = "STUB: not implemented"; return false }

func (sd *SignatureDetails) Permissions() int { _ = "STUB: not implemented"; return 0 }

func (sd SignatureDetails) String() string { _ = "STUB: not implemented"; return "" }

type SignatureValidationResult struct {
	Signature
	Status      SignatureStatus
	Reason      SignatureReason
	Details     SignatureDetails
	DocModified int
	Problems    []string
}

func (svr *SignatureValidationResult) AddProblem(s string) { _ = "STUB: not implemented"; return }

func (svr *SignatureValidationResult) Certified() bool { _ = "STUB: not implemented"; return false }

func (svr *SignatureValidationResult) Permissions() int { _ = "STUB: not implemented"; return 0 }

func (svr *SignatureValidationResult) SigningTime() string { _ = "STUB: not implemented"; return "" }

func (svr SignatureValidationResult) String() string { _ = "STUB: not implemented"; return "" }

func statusString(status int) string { _ = "STUB: not implemented"; return "" }

func validString(status int) string { _ = "STUB: not implemented"; return "" }
