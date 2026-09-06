package main

import (
	"github.com/pdfcpu/pdfcpu/pkg/cli"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"github.com/spf13/cobra"
)

type bookmarksImportOptions struct {
	replaceBookmarks bool
}

func bookmarksCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

type viewerpreferencesListOptions struct {
	all  bool
	json bool
}

func pagelayoutCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func pagemodeCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func viewerprefCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func listSinglePDFCommand(conf *model.Configuration, args []string, command func(string, *model.Configuration) *cli.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func handleListPageLayoutCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func setDocumentViewCommand(conf *model.Configuration, args []string, valid func(string) bool, invalidMsg string, command func(string, string, string, *model.Configuration) *cli.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func handleSetPageLayoutCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func resetDocumentViewCommand(conf *model.Configuration, args []string, command func(string, string, *model.Configuration) *cli.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func handleResetPageLayoutCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleListPageModeCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleSetPageModeCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleResetPageModeCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleListViewerPreferencesCommand(conf *model.Configuration, args []string, opts *viewerpreferencesListOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func viewerPreferenceInput(args []string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func handleSetViewerPreferencesCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleResetViewerPreferencesCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

type watermarkOptions struct {
	mode string
}

type stampOptions struct {
	mode string
}

type annotationListOptions struct {
	json bool
}

func annotationsCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func stampCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func watermarkCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func validateWatermarkMode(wmMode string) error { _ = "STUB: not implemented"; return nil }

func parseWatermark(args []string, onTop bool, wmMode string, unit types.DisplayUnit) (*model.Watermark, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func watermarkCommand(conf *model.Configuration, args []string, onTop bool, wmMode string, update bool) error {
	_ = "STUB: not implemented"
	return nil
}

func addWatermarks(conf *model.Configuration, args []string, onTop bool, wmMode string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleAddStampsCommand(conf *model.Configuration, args []string, opts *stampOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func handleAddWatermarksCommand(conf *model.Configuration, args []string, opts *watermarkOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func updateWatermarks(conf *model.Configuration, args []string, onTop bool, wmMode string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleUpdateStampsCommand(conf *model.Configuration, args []string, opts *stampOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func handleUpdateWatermarksCommand(conf *model.Configuration, args []string, opts *watermarkOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func removeWatermarks(conf *model.Configuration, args []string, onTop bool) error {
	_ = "STUB: not implemented"
	return nil
}

func handleRemoveStampsCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleRemoveWatermarksCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleListAnnotationsCommand(conf *model.Configuration, args []string, opts *annotationListOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func annotationRemovalArgs(conf *model.Configuration, args []string) (string, string, []string, []int, error) {
	_ = "STUB: not implemented"
	return "", "", nil, nil, nil
}

func annotationOutFile(args []string) string { _ = "STUB: not implemented"; return "" }

func handleRemoveAnnotationsCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleListBookmarksCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleExportBookmarksCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleImportBookmarksCommand(conf *model.Configuration, args []string, opts *bookmarksImportOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func handleRemoveBookmarksCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
