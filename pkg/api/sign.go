package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func signatureStats(signValidResults []*model.SignatureValidationResult) model.SignatureStats {
	_ = "STUB: not implemented"
	return *new(model.SignatureStats)
}

func statsCounter(stats model.SignatureStats, ss *[]string) { _ = "STUB: not implemented"; return }

func digest(signValidResults []*model.SignatureValidationResult, full bool) []string {
	_ = "STUB: not implemented"
	return nil
}

func ValidateSignatures(inFile string, all bool, conf *model.Configuration) (svr []*model.SignatureValidationResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ValidateSignaturesFile(inFile string, all, full bool, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RemoveSignatures(rs io.ReadSeeker, w io.Writer, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemoveSignaturesFile(inFile, outFile string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
