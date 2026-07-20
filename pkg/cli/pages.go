package cli

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func NUpCommand(inFiles []string, outFile string, pageSelection []string, nUp *model.NUp, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func BookletCommand(inFiles []string, outFile string, pageSelection []string, nup *model.NUp, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ResizeCommand(inFile, outFile string, pageSelection []string, resize *model.Resize, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func PosterCommand(inFile, outDir, outFile string, pageSelection []string, cut *model.Cut, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func NDownCommand(inFile, outDir, outFile string, pageSelection []string, n int, cut *model.Cut, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func CutCommand(inFile, outDir, outFile string, pageSelection []string, cut *model.Cut, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ZoomCommand(inFile, outFile string, pageSelection []string, zoom *model.Zoom, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func InsertPagesCommand(inFile, outFile string, pageSelection []string, conf *model.Configuration, mode string, pageConf *pdfcpu.PageConfiguration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func RemovePagesCommand(inFile, outFile string, pageSelection []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func RotateCommand(inFile, outFile string, rotation int, pageSelection []string, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func CropCommand(inFile, outFile string, pageSelection []string, box *model.Box, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func ListBoxesCommand(inFile string, pageSelection []string, pb *model.PageBoundaries, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func AddBoxesCommand(inFile, outFile string, pageSelection []string, pb *model.PageBoundaries, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}

func RemoveBoxesCommand(inFile, outFile string, pageSelection []string, pb *model.PageBoundaries, conf *model.Configuration) *Command {
	_ = "STUB: not implemented"
	return nil
}
