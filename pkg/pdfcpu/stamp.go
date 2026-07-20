package pdfcpu

import (
	"bytes"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"github.com/pkg/errors"
)

const stampWithBBox = false

var (
	errNoWatermark        = errors.New("pdfcpu: no watermarks found")
	errCorruptOCGs        = errors.New("pdfcpu: OCProperties: corrupt OCGs element")
	ErrUnsupportedVersion = errors.New("pdfcpu: PDF 2.0 unsupported for this operation")
)

func textDescriptor(wm model.Watermark, timestampFormat string, pageNr, pageCount int) (model.TextDescriptor, bool) {
	_ = "STUB: not implemented"
	return *new(model.TextDescriptor), false
}

var wmParamMap = parameterMap[model.Watermark]{
	"aligntext":       parseTextHorAlignment,
	"backgroundcolor": parseBackgroundColor,
	"bgcolor":         parseBackgroundColor,
	"border":          parseBorder,
	"color":           parseFillColor,
	"diagonal":        parseDiagonal,
	"fillcolor":       parseFillColor,
	"fontname":        parseFontName,
	"scriptname":      parseScriptName,
	"margins":         parseMargins,
	"mode":            parseRenderMode,
	"offset":          parsePositionOffsetWM,
	"opacity":         parseOpacity,
	"points":          parseFontSize,
	"position":        parsePositionAnchorWM,
	"rendermode":      parseRenderMode,
	"rtl":             parseRightToLeft,
	"rotation":        parseRotation,
	"scalefactor":     parseScaleFactorWM,
	"strokecolor":     parseStrokeColor,
	"url":             parseURL,
}

func parseTextHorAlignment(s string, wm *model.Watermark) error {
	_ = "STUB: not implemented"
	return nil
}

func parsePositionAnchorWM(s string, wm *model.Watermark) error {
	_ = "STUB: not implemented"
	return nil
}

func parsePositionOffsetWM(s string, wm *model.Watermark) error {
	_ = "STUB: not implemented"
	return nil
}

func parseScaleFactorWM(s string, wm *model.Watermark) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func parseFontName(s string, wm *model.Watermark) error { _ = "STUB: not implemented"; return nil }

func parseScriptName(s string, wm *model.Watermark) error { _ = "STUB: not implemented"; return nil }

func parseURL(s string, wm *model.Watermark) error { _ = "STUB: not implemented"; return nil }

func parseFontSize(s string, wm *model.Watermark) error { _ = "STUB: not implemented"; return nil }

func parseScaleFactor(s string) (float64, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func parseRightToLeft(s string, wm *model.Watermark) error { _ = "STUB: not implemented"; return nil }

func parseStrokeColor(s string, wm *model.Watermark) error { _ = "STUB: not implemented"; return nil }

func parseFillColor(s string, wm *model.Watermark) error { _ = "STUB: not implemented"; return nil }

func parseBackgroundColor(s string, wm *model.Watermark) error {
	_ = "STUB: not implemented"
	return nil
}

func parseRotation(s string, wm *model.Watermark) error { _ = "STUB: not implemented"; return nil }

func parseDiagonal(s string, wm *model.Watermark) error { _ = "STUB: not implemented"; return nil }

func parseOpacity(s string, wm *model.Watermark) error { _ = "STUB: not implemented"; return nil }

func parseRenderMode(s string, wm *model.Watermark) error { _ = "STUB: not implemented"; return nil }

func parseMargins(s string, wm *model.Watermark) error { _ = "STUB: not implemented"; return nil }

func parseBorder(s string, wm *model.Watermark) error { _ = "STUB: not implemented"; return nil }

func watermarkModeParamName(mode int) string { _ = "STUB: not implemented"; return "" }

func ValidateWatermarkModeParam(mode int, modeParm string, onTop bool) error {
	_ = "STUB: not implemented"
	return nil
}

func parseWatermarkDetails(mode int, modeParm, s string, onTop bool, u types.DisplayUnit) (*model.Watermark, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseTextWatermarkDetails(text, desc string, onTop bool, u types.DisplayUnit) (*model.Watermark, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseImageWatermarkDetails(fileName, desc string, onTop bool, u types.DisplayUnit) (*model.Watermark, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParsePDFWatermarkDetails(fileName, desc string, onTop bool, u types.DisplayUnit) (*model.Watermark, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func onTopString(onTop bool) string { _ = "STUB: not implemented"; return "" }

func parseWatermarkError(onTop bool) error { _ = "STUB: not implemented"; return nil }

func setTextWatermark(s string, wm *model.Watermark) { _ = "STUB: not implemented"; return }

func setImageWatermark(s string, wm *model.Watermark) error { _ = "STUB: not implemented"; return nil }

func setPDFWatermark(s string, wm *model.Watermark) error { _ = "STUB: not implemented"; return nil }

func setWatermarkType(mode int, s string, wm *model.Watermark) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func createPDFRes(ctx, otherCtx *model.Context, pageNrSrc, pageNrDest int, migrated map[int]int, wm *model.Watermark) error {
	_ = "STUB: not implemented"
	return nil
}

func createPDFResForWM(ctx *model.Context, wm *model.Watermark) error {
	_ = "STUB: not implemented"
	return nil
}

func createImageResForWM(ctx *model.Context, wm *model.Watermark) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func createFontResForWM(ctx *model.Context, wm *model.Watermark, fonts map[string]types.IndirectRef) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func createResourcesForWM(ctx *model.Context, wm *model.Watermark, fonts map[string]types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func ensureOCG(ctx *model.Context, onTop bool) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func prepareOCPropertiesInRoot(ctx *model.Context, onTop bool) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createFormResDict(ctx *model.Context, pageNr int, wm *model.Watermark) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cachedForm(wm model.Watermark) bool { _ = "STUB: not implemented"; return false }

func pdfFormContent(w io.Writer, pageNr int, wm model.Watermark) error {
	_ = "STUB: not implemented"
	return nil
}

func imageFormContent(w io.Writer, wm model.Watermark) { _ = "STUB: not implemented"; return }

func formContent(w io.Writer, pageNr int, wm model.Watermark) error {
	_ = "STUB: not implemented"
	return nil
}

func setupTextDescriptor(wm model.Watermark, timestampFormat string, pageNr, pageCount int) (model.TextDescriptor, bool) {
	_ = "STUB: not implemented"
	return *new(model.TextDescriptor), false
}

func drawBoundingBox(b *bytes.Buffer, wm model.Watermark, bb *types.Rectangle) {
	_ = "STUB: not implemented"
	return
}

func calcFormBoundingBox(xRefTable *model.XRefTable, w io.Writer, timestampFormat string, pageNr, pageCount int, wm *model.Watermark) bool {
	_ = "STUB: not implemented"
	return false
}

func createForm(ctx *model.Context, pageNr, pageCount int, wm *model.Watermark, withBB bool) error {
	_ = "STUB: not implemented"
	return nil
}

func createExtGStateForStamp(ctx *model.Context, opacity float64) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func insertPageResourcesForWM(pageDict types.Dict, wm model.Watermark, gsID, xoID string) error {
	_ = "STUB: not implemented"
	return nil
}

func updatePageResourcesForWM(ctx *model.Context, resDict types.Dict, wm model.Watermark, gsID, xoID *string) error {
	_ = "STUB: not implemented"
	return nil
}

func wmContent(wm *model.Watermark, gsID, xoID string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func insertPageContentsForWM(ctx *model.Context, pageDict types.Dict, wm *model.Watermark, gsID, xoID string) error {
	_ = "STUB: not implemented"
	return nil
}

func patchFirstContentStreamForWatermark(sd *types.StreamDict, gsID, xoID string, wm *model.Watermark, isLast bool) error {
	_ = "STUB: not implemented"
	return nil
}

func newContentStreamForWatermark(ctx *model.Context, gsID, xoID string, wm *model.Watermark) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func updatePageContentsForWM(ctx *model.Context, d types.Dict, wm *model.Watermark, gsID, xoID string) error {
	_ = "STUB: not implemented"
	return nil
}

func viewPort(a *model.InheritedPageAttrs) *types.Rectangle { _ = "STUB: not implemented"; return nil }

func handleLink(ctx *model.Context, pageIndRef *types.IndirectRef, d types.Dict, pageNr int, wm model.Watermark) error {
	_ = "STUB: not implemented"
	return nil
}

func addPageWatermark(ctx *model.Context, pageNr int, wm model.Watermark) error {
	_ = "STUB: not implemented"
	return nil
}

func AddWatermarks(ctx *model.Context, selectedPages types.IntSet, wm *model.Watermark) error {
	_ = "STUB: not implemented"
	return nil
}

func AddWatermarksMap(ctx *model.Context, m map[int]*model.Watermark) error {
	_ = "STUB: not implemented"
	return nil
}

func AddWatermarksSliceMap(ctx *model.Context, m map[int][]*model.Watermark) error {
	_ = "STUB: not implemented"
	return nil
}

func removeResDictEntry(ctx *model.Context, d types.Dict, entry string, ids []string, i int) error {
	_ = "STUB: not implemented"
	return nil
}

func removeExtGStates(ctx *model.Context, d types.Dict, ids []string, i int) error {
	_ = "STUB: not implemented"
	return nil
}

func removeForms(ctx *model.Context, d types.Dict, ids []string, i int) error {
	_ = "STUB: not implemented"
	return nil
}

func removeArtifacts(sd *types.StreamDict, i int) (ok bool, extGStates []string, forms []string, err error) {
	_ = "STUB: not implemented"
	return false, nil, nil, nil
}

func removeArtifactsFromPage(ctx *model.Context, sd *types.StreamDict, resDict types.Dict, i int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func locatePageContentAndResourceDict(ctx *model.Context, pageNr int) (types.Object, *types.IndirectRef, types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil, *new(types.Dict), nil
}

func removeArtifacts1(ctx *model.Context, o types.Object, entry *model.XRefTableEntry, resDict types.Dict, pageNr int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func removePageWatermark(ctx *model.Context, pageNr int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func locateOCGs(ctx *model.Context) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func detectStampOCG(ctx *model.Context, arr types.Array) error {
	_ = "STUB: not implemented"
	return nil
}

func removePageWatermarks(ctx *model.Context, selectedPages types.IntSet) error {
	_ = "STUB: not implemented"
	return nil
}

func RemoveWatermarks(ctx *model.Context, selectedPages types.IntSet) error {
	_ = "STUB: not implemented"
	return nil
}

func detectArtifacts(sd *types.StreamDict) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func findPageWatermarks(ctx *model.Context, pageDictIndRef *types.IndirectRef) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func detectPageTreeWatermarks(ctx *model.Context, root *types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func DetectPageTreeWatermarks(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func DetectWatermarks(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }
