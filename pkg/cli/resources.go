package cli

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func ImportImagesCommand(imageFiles []string, outFile string, imp *pdfcpu.Import, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ListFontsCommand(conf *model.Configuration) *Command { _ = "STUB: not implemented"; return nil }

func InstallFontsCommand(fontFiles []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func CreateCheatSheetsFontsCommand(fontFiles []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ListImagesCommand(inFiles []string, pageSelection []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func UpdateImagesCommand(inFile, imageFile, outFile string, objNrOrPageNr int, id string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ListAttachmentsCommand(inFile string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func AddAttachmentsCommand(inFile, outFile string, fileNames []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func AddAttachmentsPortfolioCommand(inFile, outFile string, fileNames []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func RemoveAttachmentsCommand(inFile, outFile string, fileNames []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ExtractAttachmentsCommand(inFile string, outDir string, fileNames []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ListKeywordsCommand(inFile string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func AddKeywordsCommand(inFile, outFile string, keywords []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func RemoveKeywordsCommand(inFile, outFile string, keywords []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ListPropertiesCommand(inFile string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func AddPropertiesCommand(inFile, outFile string, properties map[string]string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func RemovePropertiesCommand(inFile, outFile string, propKeys []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}
