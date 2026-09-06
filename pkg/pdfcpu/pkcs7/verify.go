package pkcs7

import (
	"crypto"
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"time"
)

func VerifyMessageDigestDetached(signer SignerInfo, signedData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func VerifyMessageDigestEmbedded(digest, signedData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func VerifyMessageDigestTSToken(oidHashAlg asn1.ObjectIdentifier, digest, signedData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func compareMessageDigest(hash crypto.Hash, expected, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func CheckSignature(
	cert *x509.Certificate,
	signer SignerInfo,
	content []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

func CheckSignatureWithContentType(
	cert *x509.Certificate,
	signer SignerInfo,
	content []byte,
	contentType asn1.ObjectIdentifier,
) error {
	_ = "STUB: not implemented"
	return nil
}

func checkSignature(
	cert *x509.Certificate,
	signer SignerInfo,
	content []byte,
	contentType asn1.ObjectIdentifier,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSignedContentType(
	signer SignerInfo,
	contentType asn1.ObjectIdentifier,
) error {
	_ = "STUB: not implemented"
	return nil
}

func verifyCryptographicSignature(
	cert *x509.Certificate,
	signer SignerInfo,
	sigalg x509.SignatureAlgorithm,
	signedData []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

func classifyCryptographicVerificationError(err error) error { _ = "STUB: not implemented"; return nil }

func validateCertificatePublicKey(publicKey any) error { _ = "STUB: not implemented"; return nil }

func validateRSAPublicKey(key *rsa.PublicKey) error { _ = "STUB: not implemented"; return nil }

func validateECDSAPublicKey(key *ecdsa.PublicKey) error { _ = "STUB: not implemented"; return nil }

func validateDSAPublicKey(key *dsa.PublicKey) error { _ = "STUB: not implemented"; return nil }

func parseRawCertificateSet(raw asn1.RawContent) (certs []*x509.Certificate, crls []*x509.RevocationList, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func parseCertificateSetEntries(rest []byte) (certs []*x509.Certificate, crls []*x509.RevocationList, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func parseSingleCertificateSetEntry(raw []byte) ([]*x509.Certificate, []*x509.RevocationList, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func parseSignedData(data []byte) (*PKCS7, error) { _ = "STUB: not implemented"; return nil, nil }

func validateSignerDigestAlgorithms(
	declared []pkix.AlgorithmIdentifier,
	signers []SignerInfo,
) error {
	_ = "STUB: not implemented"
	return nil
}

func parseEncapsulatedContent(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseOctetString(value asn1.RawValue) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func VerifyCertChain(
	ee *x509.Certificate,
	certs []*x509.Certificate,
	truststore *x509.CertPool,
	currentTime time.Time,
) ([][]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type MessageDigestMismatchError struct {
	ExpectedDigest []byte
	ActualDigest   []byte
}

func (err *MessageDigestMismatchError) Error() string { _ = "STUB: not implemented"; return "" }

func getSignatureAlgorithm(digestEncryption, digest pkix.AlgorithmIdentifier) (x509.SignatureAlgorithm, error) {
	_ = "STUB: not implemented"
	return *new(x509.SignatureAlgorithm), nil
}

func directSignatureAlgorithm(
	oid asn1.ObjectIdentifier,
) (x509.SignatureAlgorithm, asn1.ObjectIdentifier, bool) {
	_ = "STUB: not implemented"
	return *new(x509.SignatureAlgorithm), *new(asn1.ObjectIdentifier), false
}

func rsaSignatureAlgorithm(digest asn1.ObjectIdentifier) (x509.SignatureAlgorithm, error) {
	_ = "STUB: not implemented"
	return *new(x509.SignatureAlgorithm), nil
}

var oidMaskGenAlgorithmMGF1 = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 8}

type rsaPSSParametersASN1 struct {
	HashAlgorithm    pkix.AlgorithmIdentifier `asn1:"optional,explicit,tag:0"`
	MaskGenAlgorithm pkix.AlgorithmIdentifier `asn1:"optional,explicit,tag:1"`
	SaltLength       int                      `asn1:"optional,explicit,tag:2,default:20"`
	TrailerField     int                      `asn1:"optional,explicit,tag:3,default:1"`
}

type parsedRSAPSSParameters struct {
	HashAlgorithm    asn1.ObjectIdentifier
	MGFHashAlgorithm asn1.ObjectIdentifier
	SaltLength       int
	TrailerField     int
}

func parseRSAPSSParameters(raw asn1.RawValue) (parsedRSAPSSParameters, error) {
	_ = "STUB: not implemented"
	return *new(parsedRSAPSSParameters), nil
}

func nullOrAbsentParameters(raw asn1.RawValue) bool { _ = "STUB: not implemented"; return false }

func rsaPSSSignatureAlgorithm(
	raw asn1.RawValue,
	digest asn1.ObjectIdentifier,
) (x509.SignatureAlgorithm, error) {
	_ = "STUB: not implemented"
	return *new(x509.SignatureAlgorithm), nil
}

func dsaSignatureAlgorithm(digest asn1.ObjectIdentifier) (x509.SignatureAlgorithm, error) {
	_ = "STUB: not implemented"
	return *new(x509.SignatureAlgorithm), nil
}

func isEllipticCurveAlgorithm(oid asn1.ObjectIdentifier) bool {
	_ = "STUB: not implemented"
	return false
}

func ecdsaSignatureAlgorithm(digest asn1.ObjectIdentifier) (x509.SignatureAlgorithm, error) {
	_ = "STUB: not implemented"
	return *new(x509.SignatureAlgorithm), nil
}

func GetCertFromCertsByIssuerAndSerial(
	certs []*x509.Certificate,
	ias issuerAndSerial,
) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unmarshalAttribute(attrs []attribute, attributeType asn1.ObjectIdentifier, out interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
