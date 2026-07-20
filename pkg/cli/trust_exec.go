package cli

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
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

func listPEM(fName string, ss *[]string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func listP7C(fName string, ss *[]string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func ListCertificatesAll(json bool, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func listCertificatesAllJSON() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func certificateFilesJSON(dir string) ([]certificateFileEntry, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func skipCertificateFile(path string, d os.DirEntry) bool { _ = "STUB: not implemented"; return false }

func certificateFileJSON(dir, path string) certificateFileEntry {
	_ = "STUB: not implemented"
	return *new(certificateFileEntry)
}

func certificateListEntries(certs []*x509.Certificate) []certificateListEntry {
	_ = "STUB: not implemented"
	return nil
}

func newCertificateName(name pkix.Name) certificateName {
	_ = "STUB: not implemented"
	return *new(certificateName)
}

func ListCertificates(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ImportCertificates(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func InspectCertificates(cmd *Command) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ValidateSignatures(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func RemoveSignatures(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
