package model

import (
	"crypto/x509"
	"crypto/x509/pkix"
)

var TrustedCertDir string

var UserCertPool *x509.CertPool

func IsPEM(fname string) bool { _ = "STUB: not implemented"; return false }

func IsP7C(fname string) bool { _ = "STUB: not implemented"; return false }

func strSliceString(ss []string) string { _ = "STUB: not implemented"; return "" }

func nameString(subj pkix.Name) string { _ = "STUB: not implemented"; return "" }

func CertString(cert *x509.Certificate) string { _ = "STUB: not implemented"; return "" }

func ResetCertificates() error { _ = "STUB: not implemented"; return nil }

func resetCertificatesDir() error { _ = "STUB: not implemented"; return nil }
