package cli

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func Encrypt(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func Decrypt(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ChangeUserPassword(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ChangeOwnerPassword(cmd *Command) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func listPermissions(rs io.ReadSeeker, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListPermissionsFile(inFiles []string, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListPermissions(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func SetPermissions(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
