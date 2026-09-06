package cli

import "github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"

func EncryptCommand(inFile, outFile string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func DecryptCommand(inFile, outFile string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ChangeUserPWCommand(inFile, outFile string, pwOld, pwNew *string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ChangeOwnerPWCommand(inFile, outFile string, pwOld, pwNew *string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ListPermissionsCommand(inFiles []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func SetPermissionsCommand(inFile, outFile string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}
