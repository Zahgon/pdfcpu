package cli

import "github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"

func AddWatermarksCommand(inFile, outFile string, pageSelection []string, wm *model.Watermark, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func RemoveWatermarksCommand(inFile, outFile string, pageSelection []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func listAnnotationsCommand(inFile string, pageSelection []string, json bool, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ListAnnotationsCommand(inFile string, pageSelection []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ListAnnotationsJSONCommand(inFile string, pageSelection []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func RemoveAnnotationsCommand(inFile, outFile string, pageSelection []string, idsAndTypes []string, objNrs []int, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ListBookmarksCommand(inFile string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ExportBookmarksCommand(inFile, outFileJSON string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ImportBookmarksCommand(inFile, inFileJSON, outFile string, replace bool, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func RemoveBookmarksCommand(inFile, outFile string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ListPageLayoutCommand(inFile string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func SetPageLayoutCommand(inFile, outFile, value string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ResetPageLayoutCommand(inFile, outFile string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ListPageModeCommand(inFile string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func SetPageModeCommand(inFile, outFile, value string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ResetPageModeCommand(inFile, outFile string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ListViewerPreferencesCommand(inFile string, all, json bool, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func SetViewerPreferencesCommand(inFilePDF, inFileJSON, outFilePDF, stringJSON string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ResetViewerPreferencesCommand(inFile, outFile string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}
