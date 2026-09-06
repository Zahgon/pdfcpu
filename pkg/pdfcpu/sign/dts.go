package sign

import (
	"crypto"
	"crypto/x509"
	"encoding/asn1"
	"io"
	"math/big"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/pkcs7"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type AlgorithmIdentifier struct {
	Algorithm  asn1.ObjectIdentifier
	Parameters asn1.RawValue `asn1:"tag:0,optional"`
}

type TSTInfo struct {
	Version        int
	Policy         asn1.ObjectIdentifier
	MessageImprint struct {
		HashAlgorithm AlgorithmIdentifier
		HashedMessage []byte
	}
	SerialNumber asn1.RawValue
	GenTime      time.Time
	Accuracy     asn1.RawValue `asn1:"optional"`
	Ordering     bool          `asn1:"optional"`
	Nonce        asn1.RawValue `asn1:"optional"`
	TSA          asn1.RawValue `asn1:"optional"`
	Extensions   asn1.RawValue `asn1:"optional"`
}

type essIssuerSerial struct {
	Issuer       []asn1.RawValue
	SerialNumber *big.Int
}

type essCertID struct {
	CertHash     []byte
	IssuerSerial asn1.RawValue `asn1:"optional"`
}

type essCertIDv2 struct {
	HashAlgorithm asn1.RawValue `asn1:"optional"`
	CertHash      []byte
	IssuerSerial  asn1.RawValue `asn1:"optional"`
}

type signingCertificate struct {
	Certs    []essCertID
	Policies asn1.RawValue `asn1:"optional"`
}

type signingCertificateV2 struct {
	Certs    []essCertIDv2
	Policies asn1.RawValue `asn1:"optional"`
}

type rawSigningCertificateV2 struct {
	Certs    []asn1.RawValue
	Policies asn1.RawValue `asn1:"optional"`
}

func ValidateDTS(
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

func validateDTS(
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

func authenticateDTSEvidence(
	p7 *pkcs7.PKCS7,
	p7Signer pkcs7.SignerInfo,
	signerCert *x509.Certificate,
	digestErr error,
	signer *model.Signer,
	result *model.SignatureValidationResult,
) bool {
	_ = "STUB: not implemented"
	return false
}

func evaluateDTSTimestamp(
	sigDict types.Dict,
	tstInfo *TSTInfo,
	p7 *pkcs7.PKCS7,
	p7Signer pkcs7.SignerInfo,
	data []byte,
	signerCert *x509.Certificate,
	certs []*x509.Certificate,
	rootCerts *x509.CertPool,
	crls, ocsps [][]byte,
	signer *model.Signer,
	result *model.SignatureValidationResult,
	ctx *model.Context,
	localAssessment localSignatureAssessment,
) error {
	_ = "STUB: not implemented"
	return nil
}

func dtsValidationMaterial(
	ctx *model.Context,
	signer *model.Signer,
	result *model.SignatureValidationResult,
	certs []*x509.Certificate,
) ([]*x509.Certificate, [][]byte, [][]byte) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func readDTSSignedData(
	ra io.ReaderAt,
	sigDict types.Dict,
	result *model.SignatureValidationResult,
) ([]byte, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func dtsSignerCertificate(
	certs []*x509.Certificate,
	p7Signer pkcs7.SignerInfo,
	signer *model.Signer,
	result *model.SignatureValidationResult,
) *x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func reportDTSSignatureError(err error, signer *model.Signer, result *model.SignatureValidationResult) {
	_ = "STUB: not implemented"
	return
}

func checkTSTInfoProfile(tstInfo *TSTInfo, signer *model.Signer, result *model.SignatureValidationResult) bool {
	_ = "STUB: not implemented"
	return false
}

func checkDTSSignedAttributes(p7Signer pkcs7.SignerInfo, signer *model.Signer, result *model.SignatureValidationResult) bool {
	_ = "STUB: not implemented"
	return false
}

func checkTimestampingEKU(cert *x509.Certificate, signer *model.Signer, result *model.SignatureValidationResult) bool {
	_ = "STUB: not implemented"
	return false
}

func validateTimestampingEKU(cert *x509.Certificate) error { _ = "STUB: not implemented"; return nil }

func checkTimestampSigningCertificate(
	p7Signer pkcs7.SignerInfo,
	cert *x509.Certificate,
	signer *model.Signer,
	result *model.SignatureValidationResult,
) bool {
	_ = "STUB: not implemented"
	return false
}

func validateESSCertificateBinding(p7Signer pkcs7.SignerInfo, cert *x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func essSigningCertificateAttribute(p7Signer pkcs7.SignerInfo) (asn1.ObjectIdentifier, asn1.RawValue, error) {
	_ = "STUB: not implemented"
	return *new(asn1.ObjectIdentifier), *new(asn1.RawValue), nil
}

func decodeESSCertificateBinding(oid asn1.ObjectIdentifier, value asn1.RawValue) (crypto.Hash, []byte, asn1.RawValue, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Hash), nil, *new(asn1.RawValue), nil
}

func parseESSCertIDv2(raw asn1.RawValue) (essCertIDv2, error) {
	_ = "STUB: not implemented"
	return *new(essCertIDv2), nil
}

func unmarshalRawValue(bb []byte) (asn1.RawValue, []byte, error) {
	_ = "STUB: not implemented"
	return *new(asn1.RawValue), nil, nil
}

func malformedESSCertIDv2(detail string) error { _ = "STUB: not implemented"; return nil }

func unmarshalESSCertificateAttribute(value asn1.RawValue, oid asn1.ObjectIdentifier, out any) error {
	_ = "STUB: not implemented"
	return nil
}

func missingESSCertificateID(oid asn1.ObjectIdentifier) error {
	_ = "STUB: not implemented"
	return nil
}

func essCertIDv2Hash(raw asn1.RawValue) (crypto.Hash, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Hash), nil
}

func matchESSCertificateHash(hash crypto.Hash, want []byte, cert *x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func matchESSCertificateIssuerSerial(raw asn1.RawValue, cert *x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func validateESSCertificateKeyUsage(cert *x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func parseTSTInfo(bb []byte) (*TSTInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func validateDTSCert(
	signerCert *x509.Certificate,
	certs []*x509.Certificate,
	rootCerts *x509.CertPool,
	crls, ocsps [][]byte,
	signer *model.Signer,
	result *model.SignatureValidationResult,
	ctx *model.Context,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func isLocallyValidatedTSAPathAssessment(assessment certificateAssessment) bool {
	_ = "STUB: not implemented"
	return false
}

func verifyDTSDigest(tstInfo *TSTInfo, data []byte) error { _ = "STUB: not implemented"; return nil }

func applyDTSDigestEvidence(err error, signer *model.Signer, result *model.SignatureValidationResult) bool {
	_ = "STUB: not implemented"
	return false
}

func collectIntermediates(signerCert *x509.Certificate, certs []*x509.Certificate) []*x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func finalizeDTSValidationResult(result *model.SignatureValidationResult, ctx *model.Context, evidence timestampEvidence, assessment localSignatureAssessment) {
	_ = "STUB: not implemented"
	return
}
