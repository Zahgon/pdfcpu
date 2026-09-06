package pkcs7

import (
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"io"
	"math/big"
)

type SignedData struct {
	sd    signedData
	certs []*x509.Certificate
}

func NewSignedData() (*SignedData, error) { _ = "STUB: not implemented"; return nil, nil }

type SignerInfoConfig struct {
	ExtraSignedAttributes   []Attribute
	ExtraUnsignedAttributes []Attribute
}

type signedData struct {
	Version                    int                        `asn1:"default:1"`
	DigestAlgorithmIdentifiers []pkix.AlgorithmIdentifier `asn1:"set"`
	ContentInfo                contentInfo
	Certificates               rawCertificates `asn1:"optional,tag:0"`
	CRLs                       []asn1.RawValue `asn1:"optional,tag:1"`
	SignerInfos                []SignerInfo    `asn1:"set"`
}

type SignerInfo struct {
	Version                   int `asn1:"default:1"`
	IssuerAndSerialNumber     issuerAndSerial
	DigestAlgorithm           pkix.AlgorithmIdentifier
	AuthenticatedAttributes   []attribute `asn1:"optional,omitempty,tag:0"`
	DigestEncryptionAlgorithm pkix.AlgorithmIdentifier
	EncryptedDigest           []byte      `asn1:"octet"`
	UnauthenticatedAttributes []attribute `asn1:"optional,omitempty,tag:1"`
}

type attribute struct {
	Type  asn1.ObjectIdentifier
	Value asn1.RawValue `asn1:"set"`
}

func marshalAttributes(attrs []attribute) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type rawCertificates struct {
	Raw asn1.RawContent
}

type issuerAndSerial struct {
	IssuerName   asn1.RawValue
	SerialNumber *big.Int
}

func addDigestAlgorithmUnique(list []pkix.AlgorithmIdentifier, oid asn1.ObjectIdentifier) []pkix.AlgorithmIdentifier {
	_ = "STUB: not implemented"
	return nil
}

func (sd *SignedData) AddSigner(cert *x509.Certificate, pkey crypto.PrivateKey, messageDigest []byte, digestOid asn1.ObjectIdentifier, config SignerInfoConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (sd *SignedData) AddSignerChain(cert *x509.Certificate, pkey crypto.PrivateKey, messageDigest []byte, digestOid asn1.ObjectIdentifier, parents []*x509.Certificate, config SignerInfoConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSignerCertificates(sd *SignedData, cert *x509.Certificate, parents []*x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func validateExtraSignedAttributes(attrs []Attribute) error { _ = "STUB: not implemented"; return nil }

func validateSigningKey(cert *x509.Certificate, pkey crypto.PrivateKey) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePrivateSigningKey(pkey crypto.PrivateKey) error { _ = "STUB: not implemented"; return nil }

func signerIdentifier(cert *x509.Certificate) issuerAndSerial {
	_ = "STUB: not implemented"
	return *new(issuerAndSerial)
}

func (sd *SignedData) AddCertificate(cert *x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func marshalCertificateBytes(certs []byte) (rawCertificates, error) {
	_ = "STUB: not implemented"
	return *new(rawCertificates), nil
}

func marshalCertificates(certs []*x509.Certificate) (rawCertificates, error) {
	_ = "STUB: not implemented"
	return *new(rawCertificates), nil
}

func (sd *SignedData) Finish() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func signAttributes(attrs []attribute, pkey crypto.PrivateKey, digestAlg crypto.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func signAttributesWithRandom(
	attrs []attribute,
	pkey crypto.PrivateKey,
	digestAlg crypto.Hash,
	random io.Reader,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
