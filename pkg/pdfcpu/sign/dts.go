package sign

import (
	"crypto/x509"
	"encoding/asn1"
	"io"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
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

func ValidateDTS(
	ra io.ReaderAt,
	sigDict types.Dict,
	certified bool,
	authoritative bool,
	validateAll bool,
	perms int,
	rootCerts *x509.CertPool,
	result *model.SignatureValidationResult,
	ctx *model.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDTSCert(signingTime time.Time,
	signerCert *x509.Certificate,
	certs []*x509.Certificate,
	rootCerts *x509.CertPool,
	crls, ocsps [][]byte,
	signer *model.Signer,
	result *model.SignatureValidationResult,
	ctx *model.Context) {
	_ = "STUB: not implemented"
	return
}

func checkDTSDigest(tstInfo *TSTInfo, data []byte, signer *model.Signer) bool {
	_ = "STUB: not implemented"
	return false
}

func collectIntermediates(signerCert *x509.Certificate, certs []*x509.Certificate) []*x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func finalizeDTSResult(result *model.SignatureValidationResult, ctx *model.Context, signingTime time.Time) {
	_ = "STUB: not implemented"
	return
}
