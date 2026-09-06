package cli

import "github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"

func ValidateCommand(inFiles []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func OptimizeCommand(inFile, outFile string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func InfoCommand(inFiles []string, pageSelection []string, fonts, json bool, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func DumpCommand(inFilePDF string, vals []int, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func CreateCommand(inFilePDF, inFileJSON, outFilePDF string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func MergeCreateCommand(inFiles []string, outFile string, dividerPage bool, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func MergeCreateZipCommand(inFiles []string, outFile string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func MergeAppendCommand(inFiles []string, outFile string, dividerPage bool, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func SplitCommand(inFile, dirNameOut string, span int, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func SplitByPageNrCommand(inFile, dirNameOut string, pageNrs []int, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func TrimCommand(inFile, outFile string, pageSelection []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func CollectCommand(inFile, outFile string, pageSelection []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}
