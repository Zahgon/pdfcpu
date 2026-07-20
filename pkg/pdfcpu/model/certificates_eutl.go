//go:build pdfcpu_eutl
// +build pdfcpu_eutl

package model

import (
	"embed"
)

//go:embed resources/certs/*.p7c
var certFilesEU embed.FS

func installDefaultCertificates() error { _ = "STUB: not implemented"; return nil }

func installDefaultCertificate(name, dir string) error { _ = "STUB: not implemented"; return nil }
