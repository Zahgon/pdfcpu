package api

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"os"
)

type certificateName struct {
	Organization       []string `json:"organization,omitempty"`
	OrganizationalUnit []string `json:"organizationalUnit,omitempty"`
	CommonName         string   `json:"commonName,omitempty"`
	StreetAddress      []string `json:"streetAddress,omitempty"`
	Locality           []string `json:"locality,omitempty"`
	Province           []string `json:"province,omitempty"`
	PostalCode         []string `json:"postalCode,omitempty"`
	Country            []string `json:"country,omitempty"`
}

type certificateListEntry struct {
	Subject      certificateName `json:"subject"`
	Issuer       certificateName `json:"issuer"`
	SerialNumber string          `json:"serialNumber"`
	NotBefore    string          `json:"notBefore"`
	NotAfter     string          `json:"notAfter"`
	IsCA         bool            `json:"isCA"`
}

type certificateFileEntry struct {
	Name         string                 `json:"name"`
	Certificates []certificateListEntry `json:"certificates,omitempty"`
	Error        string                 `json:"error,omitempty"`
}

type certificateImport struct {
	inFile       string
	outFile      string
	certificates []*x509.Certificate
}

type stagedCertificateImport struct {
	certificateImport
	stageFile   string
	backupFile  string
	hadOriginal bool
	published   bool
}

type certificateImportOperations struct {
	files            fileOperations
	saveCertificates func([]*x509.Certificate, string) error
}

func defaultCertificateImportOperations() certificateImportOperations {
	_ = "STUB: not implemented"
	return *new(certificateImportOperations)
}

func validateCertificateFiles(inFiles []string) error { _ = "STUB: not implemented"; return nil }

func ensureTrustedCertificateDir() error { _ = "STUB: not implemented"; return nil }

func certificateStrings(certs []*x509.Certificate) []string { _ = "STUB: not implemented"; return nil }

func appendCertificateFile(path string, ss *[]string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func listCertificatesText() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func listCertificatesJSON() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func certificateFilesJSON(dir string) ([]certificateFileEntry, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func skipCertificateFile(path string, d os.DirEntry) bool { _ = "STUB: not implemented"; return false }

func certificateFileJSON(dir, path string) (certificateFileEntry, error) {
	_ = "STUB: not implemented"
	return *new(certificateFileEntry), nil
}

func certificateListEntries(certs []*x509.Certificate) []certificateListEntry {
	_ = "STUB: not implemented"
	return nil
}

func newCertificateName(name pkix.Name) certificateName {
	_ = "STUB: not implemented"
	return *new(certificateName)
}

func ListCertificates(json bool) (ss []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ImportCertificates(inFiles []string) (ss []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func importCertificates(inFiles []string, ops certificateImportOperations) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func certificateImportDestination(inFile string) string { _ = "STUB: not implemented"; return "" }

func prepareCertificateImports(inFiles []string) ([]certificateImport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createCertificateTransactionFile(outFile, kind string, ops certificateImportOperations) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func cleanupCertificateImports(staged []stagedCertificateImport, removeBackups bool, ops certificateImportOperations) error {
	_ = "STUB: not implemented"
	return nil
}

func stageCertificateImports(imports []certificateImport, ops certificateImportOperations) ([]stagedCertificateImport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func backupCertificateDestinations(staged []stagedCertificateImport, ops certificateImportOperations) error {
	_ = "STUB: not implemented"
	return nil
}

func rollbackCertificateImports(staged []stagedCertificateImport, ops certificateImportOperations) error {
	_ = "STUB: not implemented"
	return nil
}

func publishCertificateImports(imports []certificateImport, ops certificateImportOperations) error {
	_ = "STUB: not implemented"
	return nil
}

func certificateImportSummary(imports []certificateImport) []string {
	_ = "STUB: not implemented"
	return nil
}

func inspectCertificate(inFile string, inputIndex, certificateIndex int, cert *x509.Certificate) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func InspectCertificates(inFiles []string) (ss []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
