package pkcs7

import (
	"crypto"
	"crypto/x509"
	"encoding/asn1"
	"errors"

	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
)

type PKCS7 struct {
	Content      []byte
	ContentType  asn1.ObjectIdentifier
	Certificates []*x509.Certificate
	CRLs         []*x509.RevocationList
	Signers      []SignerInfo
}

type contentInfo struct {
	ContentType asn1.ObjectIdentifier
	Content     asn1.RawValue `asn1:"explicit,optional,tag:0"`
}

var ErrUnsupportedContentType = errors.New("pkcs7: cannot parse data: unimplemented content type")

var ErrEmptyInput = errors.New("pkcs7: input data is empty")

var ErrUnsupportedAlgorithm = errors.New("pkcs7: unsupported algorithm")

var ErrAlgorithmMismatch = errors.New("pkcs7: algorithm mismatch")

var ErrInvalidPSSParameters = errors.New("pkcs7: invalid RSASSA-PSS parameters")

var ErrSignatureMismatch = errors.New("pkcs7: cryptographic signature mismatch")

var ErrCertificateParse = errors.New("pkcs7: certificate parse error")

var ErrMissingSignerIdentifier = errors.New("pkcs7: missing signer identifier")

var ErrMalformedAttribute = errors.New("pkcs7: malformed attribute")

type CertificateParseError struct {
	cause error
}

func (e *CertificateParseError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *CertificateParseError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (e *CertificateParseError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func certificateParseError(err error) error { _ = "STUB: not implemented"; return nil }

var (
	OIDData                   = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 7, 1}
	OIDSignedData             = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 7, 2}
	oidAttributeContentType   = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 3}
	oidAttributeMessageDigest = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 4}

	OIDDigestAlgorithmSHA1   = asn1.ObjectIdentifier{1, 3, 14, 3, 2, 26}
	OIDDigestAlgorithmSHA256 = asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 2, 1}
	OIDDigestAlgorithmSHA384 = asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 2, 2}
	OIDDigestAlgorithmSHA512 = asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 2, 3}

	OIDDigestAlgorithmDSA     = asn1.ObjectIdentifier{1, 2, 840, 10040, 4, 1}
	OIDDigestAlgorithmDSASHA1 = asn1.ObjectIdentifier{1, 2, 840, 10040, 4, 3}

	OIDDigestAlgorithmECDSASHA1   = asn1.ObjectIdentifier{1, 2, 840, 10045, 4, 1}
	OIDDigestAlgorithmECDSASHA256 = asn1.ObjectIdentifier{1, 2, 840, 10045, 4, 3, 2}
	OIDDigestAlgorithmECDSASHA384 = asn1.ObjectIdentifier{1, 2, 840, 10045, 4, 3, 3}
	OIDDigestAlgorithmECDSASHA512 = asn1.ObjectIdentifier{1, 2, 840, 10045, 4, 3, 4}

	OIDEncryptionAlgorithmRSA       = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}
	OIDEncryptionAlgorithmRSASHA1   = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 5}
	OIDEncryptionAlgorithmRSAPSS    = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 10}
	OIDEncryptionAlgorithmRSASHA256 = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 11}
	OIDEncryptionAlgorithmRSASHA384 = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 12}
	OIDEncryptionAlgorithmRSASHA512 = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 13}

	OIDEncryptionAlgorithmECDSAP256   = asn1.ObjectIdentifier{1, 2, 840, 10045, 3, 1, 7}
	OIDEncryptionAlgorithmECDSAP384   = asn1.ObjectIdentifier{1, 3, 132, 0, 34}
	OIDEncryptionAlgorithmECDSAP521   = asn1.ObjectIdentifier{1, 3, 132, 0, 35}
	OIDEncryptionAlgorithmECPUBLICKEY = asn1.ObjectIdentifier{1, 2, 840, 10045, 2, 1}

	OIDEncryptionAlgorithmEd25519 = asn1.ObjectIdentifier{1, 3, 101, 112}
)

func HashForOID(oid asn1.ObjectIdentifier) (crypto.Hash, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Hash), nil
}

func OIDForEncryptionAlgorithm(pkey crypto.PrivateKey, OIDDigestAlg asn1.ObjectIdentifier) (asn1.ObjectIdentifier, error) {
	_ = "STUB: not implemented"
	return *new(asn1.ObjectIdentifier), nil
}

func Parse(data []byte) (*PKCS7, error) { _ = "STUB: not implemented"; return nil, nil }

func isCertMatchForIssuerAndSerial(cert *x509.Certificate, ias issuerAndSerial) bool {
	_ = "STUB: not implemented"
	return false
}

func validateSignerIdentifier(ias issuerAndSerial) error { _ = "STUB: not implemented"; return nil }

type Attribute struct {
	Type  asn1.ObjectIdentifier
	Value interface{}
}

type attributes struct {
	types  []asn1.ObjectIdentifier
	values []interface{}
}

func (attrs *attributes) Add(attrType asn1.ObjectIdentifier, value interface{}) {
	_ = "STUB: not implemented"
	return
}

type sortableAttribute struct {
	SortKey   []byte
	Attribute attribute
}

type attributeSet []sortableAttribute

func (sa attributeSet) Len() int { _ = "STUB: not implemented"; return 0 }

func (sa attributeSet) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (sa attributeSet) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (sa attributeSet) Attributes() []attribute { _ = "STUB: not implemented"; return nil }

func (attrs *attributes) ForMarshalling() ([]attribute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
