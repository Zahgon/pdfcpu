package pdfcpu

import (
	"crypto/x509"
	"errors"
	"sync"
)

var ErrUnsupportedCertificateFile = errors.New("unsupported file type")

var ErrUnknownFileType = ErrUnsupportedCertificateFile

var ErrNoCertificates = errors.New("no certificates found")

var ErrMissingCertificate = errors.New("missing certificate")

type certificatePoolCache struct {
	sync.RWMutex
	dir           string
	loaded        bool
	pool          *x509.CertPool
	storeRevision uint64
}

var trustedCertificatePool certificatePoolCache

func addCertificateFileToPool(path string, certPool *x509.CertPool) error {
	_ = "STUB: not implemented"
	return nil
}

func buildCertificatePool(dir string) (*x509.CertPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildCurrentCertificatePool(dir string) (*x509.CertPool, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func LoadCertificates() error { _ = "STUB: not implemented"; return nil }

func InvalidateCertificatePool() { _ = "STUB: not implemented"; return }

func userCertificatePool() *x509.CertPool { _ = "STUB: not implemented"; return nil }

func loadSingleCertFile(filename string) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadCertsFromPEM(filename string) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const (
	pkcs7PEMType   = "PKCS7"
	pkcs7PEMPrefix = "-----BEGIN PKCS7-----"
	pkcs7PEMSuffix = "-----END PKCS7-----"
)

func decodePKCS7PEM(bb []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func loadCertsFromP7C(filename string) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LoadCertificatesFile(filename string) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func saveCertsAsPEM(certs []*x509.Certificate, filename string, overwrite bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func saveCertsAsP7C(certs []*x509.Certificate, filename string, overwrite bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func SaveCertificates(certs []*x509.Certificate, outFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func InspectCertificate(cert *x509.Certificate) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
