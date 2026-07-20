package api

import (
	"bytes"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func ListFonts() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func InstallFonts(fileNames []string) error { _ = "STUB: not implemented"; return nil }

func rowLabel(xRefTable *model.XRefTable, i int, td model.TextDescriptor, baseFontName, baseFontKey string, buf *bytes.Buffer, mb *types.Rectangle, left bool) {
	_ = "STUB: not implemented"
	return
}

func columnsLabel(xRefTable *model.XRefTable, td model.TextDescriptor, baseFontName, baseFontKey string, buf *bytes.Buffer, mb *types.Rectangle, top bool) {
	_ = "STUB: not implemented"
	return
}

func surrogate(r rune) bool { _ = "STUB: not implemented"; return false }

func writeUserFontDemoContent(xRefTable *model.XRefTable, p model.Page, fontName string, plane int) {
	_ = "STUB: not implemented"
	return
}

func createUserFontDemoPage(xRefTable *model.XRefTable, w, h, plane int, fontName string) model.Page {
	_ = "STUB: not implemented"
	return *new(model.Page)
}

func planeString(i int) string { _ = "STUB: not implemented"; return "" }

func CreateUserFontDemoFiles(dir, fn string) error { _ = "STUB: not implemented"; return nil }

func CreateCheatSheetsUserFonts(fontNames []string) error { _ = "STUB: not implemented"; return nil }
