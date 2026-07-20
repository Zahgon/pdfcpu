package pdfcpu

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func extractAuthor(ctx *model.Context, obj types.Object) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func extractCreator(ctx *model.Context, obj types.Object) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func logKey(key string) { _ = "STUB: not implemented"; return }

func handleInfoDict(ctx *model.Context, d types.Dict) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ensureInfoDict(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func writeDocumentInfoDict(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func appendEqualMediaAndCropBoxInfo(ss *[]string, pb model.PageBoundaries, unit string, currUnit types.DisplayUnit) {
	_ = "STUB: not implemented"
	return
}

func trimBleedArtBoxString(cb, tb, bb, ab *types.Rectangle) string {
	_ = "STUB: not implemented"
	return ""
}

func appendNotEqualMediaAndCropBoxInfo(ss *[]string, pb model.PageBoundaries, unit string, currUnit types.DisplayUnit) {
	_ = "STUB: not implemented"
	return
}

func appendPageBoxesInfo(ss *[]string, pb model.PageBoundaries, unit string, currUnit types.DisplayUnit, i int) {
	_ = "STUB: not implemented"
	return
}

func pageInfo(info *PDFInfo, selectedPages types.IntSet) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PDFInfo struct {
	FileName           string                          `json:"source,omitempty"`
	Version            string                          `json:"version"`
	PageCount          int                             `json:"pageCount"`
	PageBoundaries     []model.PageBoundaries          `json:"-"`
	Boundaries         map[string]model.PageBoundaries `json:"pageBoundaries,omitempty"`
	PageDimensions     map[types.Dim]bool              `json:"-"`
	Dimensions         []types.Dim                     `json:"pageSizes,omitempty"`
	Title              string                          `json:"title"`
	Author             string                          `json:"author"`
	Subject            string                          `json:"subject"`
	Producer           string                          `json:"producer"`
	Creator            string                          `json:"creator"`
	CreationDate       string                          `json:"creationDate"`
	ModificationDate   string                          `json:"modificationDate"`
	PageMode           string                          `json:"pageMode,omitempty"`
	PageLayout         string                          `json:"pageLayout,omitempty"`
	ViewerPref         *model.ViewerPreferences        `json:"viewerPreferences,omitempty"`
	Keywords           []string                        `json:"keywords"`
	Properties         map[string]string               `json:"properties"`
	Tagged             bool                            `json:"tagged"`
	Hybrid             bool                            `json:"hybrid"`
	Linearized         bool                            `json:"linearized"`
	UsingXRefStreams   bool                            `json:"usingXRefStreams"`
	UsingObjectStreams bool                            `json:"usingObjectStreams"`
	Watermarked        bool                            `json:"watermarked"`
	Thumbnails         bool                            `json:"thumbnails"`
	Form               bool                            `json:"form"`
	Signatures         bool                            `json:"signatures"`
	AppendOnly         bool                            `json:"appendOnly"`
	Outlines           bool                            `json:"bookmarks"`
	Names              bool                            `json:"names"`
	Encrypted          bool                            `json:"encrypted"`
	Permissions        int                             `json:"permissions"`
	Attachments        []model.Attachment              `json:"attachments,omitempty"`
	Unit               types.DisplayUnit               `json:"-"`
	UnitString         string                          `json:"unit"`
	Fonts              []model.FontInfo                `json:"fonts,omitempty"`
}

func (info PDFInfo) renderKeywords(ss *[]string) error { _ = "STUB: not implemented"; return nil }

func (info PDFInfo) renderProperties(ss *[]string) error { _ = "STUB: not implemented"; return nil }

func (info PDFInfo) renderFlagsPart1(ss *[]string, separator string) {
	_ = "STUB: not implemented"
	return
}

func (info PDFInfo) renderFlagsPart2(ss *[]string, separator string) {
	_ = "STUB: not implemented"
	return
}

func (info *PDFInfo) renderFlags(ss *[]string, separator string) { _ = "STUB: not implemented"; return }

func (info *PDFInfo) renderPermissions(ss *[]string) { _ = "STUB: not implemented"; return }

func (info *PDFInfo) renderAttachments(ss *[]string) { _ = "STUB: not implemented"; return }

func (info *PDFInfo) renderFonts(ss *[]string) { _ = "STUB: not implemented"; return }

func setupFontInfos(ctx *model.Context, fontInfos *[]model.FontInfo) {
	_ = "STUB: not implemented"
	return
}

func Info(ctx *model.Context, fileName string, selectedPages types.IntSet, fonts bool) (*PDFInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListInfo(info *PDFInfo, selectedPages types.IntSet, fonts bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
