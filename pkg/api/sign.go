package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

type ReadSeekerAt interface {
	io.ReadSeeker
	io.ReaderAt
}

type signatureValidationOperation func(
	io.ReaderAt,
	*model.Context,
	bool,
) ([]*model.SignatureValidationResult, error)

func signatureStats(signValidResults []*model.SignatureValidationResult) model.SignatureStats {
	_ = "STUB: not implemented"
	return *new(model.SignatureStats)
}

func statsCounter(stats model.SignatureStats, ss *[]string) { _ = "STUB: not implemented"; return }

func digest(signValidResults []*model.SignatureValidationResult, full bool) []string {
	_ = "STUB: not implemented"
	return nil
}

func compactSignatureReason(svr *model.SignatureValidationResult) string {
	_ = "STUB: not implemented"
	return ""
}

func ValidateSignatures(inFile string, all bool, conf *model.Configuration) (results []*model.SignatureValidationResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateSignaturesFile(
	inFile string,
	all bool,
	conf *model.Configuration,
	operation signatureValidationOperation,
) (results []*model.SignatureValidationResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ValidateSignaturesRaw(
	rs ReadSeekerAt,
	all bool,
	conf *model.Configuration,
) (results []*model.SignatureValidationResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateSignaturesRaw(
	rs ReadSeekerAt,
	all bool,
	conf *model.Configuration,
	operation signatureValidationOperation,
) (results []*model.SignatureValidationResult, err error) {
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
