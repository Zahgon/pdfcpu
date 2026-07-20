package cli

import "github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"

func listFormFieldsCommand(inFiles []string, json bool, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ListFormFieldsCommand(inFiles []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ListFormFieldsJSONCommand(inFiles []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func RemoveFormFieldsCommand(inFile, outFile string, fieldIDs []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func LockFormCommand(inFile, outFile string, fieldIDs []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func UnlockFormCommand(inFile, outFile string, fieldIDs []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ResetFormCommand(inFile, outFile string, fieldIDs []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ExportFormCommand(inFilePDF, outFileJSON string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func FillFormCommand(inFilePDF, inFileJSON, outFilePDF string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func MultiFillFormCommand(inFilePDF, inFileData, outDir, outFilePDF string, merge bool, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}
