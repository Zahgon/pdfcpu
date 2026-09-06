package api

import (
	"bytes"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/font"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type fontAPIOperations struct {
	userFontDir               string
	reloadUserFonts           func() error
	installTrueTypeFont       func(string, string) (font.InstallResult, error)
	installTrueTypeCollection func(string, string) ([]font.InstallResult, error)
	createStagingDir          func(string) (string, error)
	createInputDir            func(string, int) (string, error)
	commitStagedFonts         func(string, string) (fontInstallCommit, error)
	removeAll                 func(string) error
	rename                    func(string, string) error
	reportCleanupWarning      func(error)
}

type fontInstallCommit struct {
	rollback func() error
	finalize func() error
}

type committedFontFile struct {
	name        string
	hadOriginal bool
	committed   bool
}

type transactionFileOperations struct {
	mkdirTemp func(string, string) (string, error)
	lstat     func(string) (os.FileInfo, error)
	syncDir   func(string) error
	rename    func(string, string) error
	remove    func(string) error
	removeAll func(string) error
}

type fontInstallFileOperations = transactionFileOperations
type cheatSheetFileOperations = transactionFileOperations

func defaultFontInstallFileOperations() fontInstallFileOperations {
	_ = "STUB: not implemented"
	return *new(fontInstallFileOperations)
}

type fontDemoOperations struct {
	loadUserFonts func() error
	userFont      func(string) (font.TTFLight, bool, error)
	userFontNames func() ([]string, error)
	createXRef    func() (*model.XRefTable, error)
	createPage    func(*model.XRefTable, int, int, int, string) (model.Page, error)
	catalog       func(*model.XRefTable) (types.Dict, error)
	addPageTree   func(*model.XRefTable, types.Dict, model.Page) error
	createPDFFile func(*model.XRefTable, string, *model.Configuration) error
	files         cheatSheetFileOperations
}

type committedCheatSheet struct {
	name        string
	hadOriginal bool
	published   bool
}

func defaultCheatSheetFileOperations() cheatSheetFileOperations {
	_ = "STUB: not implemented"
	return *new(cheatSheetFileOperations)
}

func defaultFontAPIOperations() fontAPIOperations {
	_ = "STUB: not implemented"
	return *new(fontAPIOperations)
}

func defaultFontDemoOperations() fontDemoOperations {
	_ = "STUB: not implemented"
	return *new(fontDemoOperations)
}

func ListFonts() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func listFonts(userFontNames func() ([]string, error)) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateFontFiles(fileNames []string) error { _ = "STUB: not implemented"; return nil }

func InstallFonts(fileNames []string) error { _ = "STUB: not implemented"; return nil }

func stagedFontFiles(stagingDir string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func syncTransactionDirectories(fs transactionFileOperations, dirs ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func rollbackCommittedFonts(fontDir, backupDir string, files []committedFontFile, fs fontInstallFileOperations) error {
	_ = "STUB: not implemented"
	return nil
}

func commitStagedFonts(fontDir, stagingDir string) (fontInstallCommit, error) {
	_ = "STUB: not implemented"
	return *new(fontInstallCommit), nil
}

func commitStagedFontsWithOperations(fontDir, stagingDir string, fs fontInstallFileOperations) (fontInstallCommit, error) {
	_ = "STUB: not implemented"
	return *new(fontInstallCommit), nil
}

func stagedFontOrigin(input int, fileName string, result font.InstallResult) string {
	_ = "STUB: not implemented"
	return ""
}

func mergeStagedFont(inputDir, stagingDir string, result font.InstallResult, origin string, owners map[string]string, ops fontAPIOperations) error {
	_ = "STUB: not implemented"
	return nil
}

func installFontInput(stagingDir, fileName string, input int, ops fontAPIOperations) ([]font.InstallResult, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func installFontInputs(fileNames []string, stagingDir string, ops fontAPIOperations) error {
	_ = "STUB: not implemented"
	return nil
}

func installFonts(fileNames []string, ops fontAPIOperations) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func rowLabel(xRefTable *model.XRefTable, i int, td model.TextDescriptor, baseFontName, baseFontKey string, buf *bytes.Buffer, mb *types.Rectangle, left bool) error {
	_ = "STUB: not implemented"
	return nil
}

func columnsLabel(xRefTable *model.XRefTable, td model.TextDescriptor, baseFontName, baseFontKey string, buf *bytes.Buffer, mb *types.Rectangle, top bool) error {
	_ = "STUB: not implemented"
	return nil
}

func surrogate(r rune) bool { _ = "STUB: not implemented"; return false }

func writeUserFontDemoContent(xRefTable *model.XRefTable, p model.Page, fontName string, plane int) error {
	_ = "STUB: not implemented"
	return nil
}

func createUserFontDemoPage(xRefTable *model.XRefTable, w, h, plane int, fontName string) (p model.Page, err error) {
	_ = "STUB: not implemented"
	return *new(model.Page), nil
}

func planeString(i int) (string, error) { _ = "STUB: not implemented"; return "", nil }

func coveredUnicodePlanes(m map[int]bool) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func rollbackCheatSheets(dir, backupDir string, files []committedCheatSheet, fs cheatSheetFileOperations) error {
	_ = "STUB: not implemented"
	return nil
}

func publishCheatSheets(dir, stagingDir string, names []string, fs cheatSheetFileOperations) (published bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func normalizeCheatSheetDir(dir string) string { _ = "STUB: not implemented"; return "" }

func CreateUserFontDemoFiles(dir, fn string) (err error) { _ = "STUB: not implemented"; return nil }

func createUserFontDemoFiles(dir, fn string, ops fontDemoOperations) error {
	_ = "STUB: not implemented"
	return nil
}

func stageUserFontDemoFiles(dir, fn string, ttf font.TTFLight, ops fontDemoOperations) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createUserFontDemoBatch(dir string, fontNames []string, fonts map[string]font.TTFLight, ops fontDemoOperations) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func CreateCheatSheetsUserFonts(fontNames []string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func createCheatSheetsUserFonts(fontNames []string, ops fontDemoOperations) error {
	_ = "STUB: not implemented"
	return nil
}
