package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func Encrypt(rs io.ReadSeeker, w io.Writer, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func EncryptFile(inFile, outFile string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func Decrypt(rs io.ReadSeeker, w io.Writer, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func DecryptFile(inFile, outFile string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

type securityOperation func(io.ReadSeeker, io.Writer, *model.Configuration) error

func processSecurityFile(
	inFile, outFile string,
	conf *model.Configuration,
	op string,
	process securityOperation,
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ChangeUserPassword(rs io.ReadSeeker, w io.Writer, pwOld, pwNew string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ChangeUserPasswordFile(inFile, outFile string, pwOld, pwNew string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func ChangeOwnerPassword(rs io.ReadSeeker, w io.Writer, pwOld, pwNew string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ChangeOwnerPasswordFile(inFile, outFile string, pwOld, pwNew string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}
