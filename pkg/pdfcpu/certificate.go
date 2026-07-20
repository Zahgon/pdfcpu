package pdfcpu

import (
	"crypto/x509"
	"sync"

	"github.com/pkg/errors"
)

var ErrUnknownFileType = errors.New("pdfcpu: unsupported file type")

var loadCertsOnce sync.Once
var loadCertsErr error

func doAddCertificatesToCertPool(path string, certPool *x509.CertPool, n *int) error {
	_ = "STUB: not implemented"
	return nil
}

func addCertificatesToCertPool() error { _ = "STUB: not implemented"; return nil }

func LoadCertificates() error { _ = "STUB: not implemented"; return nil }

func loadSingleCertFile(filename string) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadCertsFromPEM(filename string) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const PKCS7_PREFIX = "-----BEGIN PKCS7-----"
const PKCS7_SUFFIX = "-----END PKCS7-----"

func isPEMEncoded(s string) bool { _ = "STUB: not implemented"; return false }

func decodePKCS7Block(s string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

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

func ImportCertificate(inFile string, overwrite bool) (int, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func InspectCertificate(cert *x509.Certificate) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
