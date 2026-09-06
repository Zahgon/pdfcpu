//go:build !pdfcpu_eutl
// +build !pdfcpu_eutl

package model

const bundledDefaultCertificates = false

func installDefaultCertificates() error { _ = "STUB: not implemented"; return nil }
