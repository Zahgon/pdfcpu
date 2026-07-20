package cli

import "github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"

func ListCertificatesCommand(json bool, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func InspectCertificatesCommand(inFiles []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ImportCertificatesCommand(inFiles []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ValidateSignaturesCommand(inFile string, all, full bool, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func RemoveSignaturesCommand(inFile, outFile string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}
