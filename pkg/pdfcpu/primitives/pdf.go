package primitives

import (
	"io"
	"net/http"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type FieldFlags int

const (
	FieldReadOnly FieldFlags = 1 << iota
	FieldRequired
	FieldNoExport
	UnusedFlag4
	UnusedFlag5
	UnusedFlag6
	UnusedFlag7
	UnusedFlag8
	UnusedFlag9
	UnusedFlag10
	UnusedFlag11
	UnusedFlag12
	FieldMultiline
	FieldPassword
	FieldNoToggleToOff
	FieldRadio
	FieldPushbutton
	FieldCombo
	FieldEdit
	FieldSort
	FieldFileSelect
	FieldMultiselect
	FieldDoNotSpellCheck
	FieldDoNotScroll
	FieldComb
	FieldRichTextAndRadiosInUnison
	FieldCommitOnSelChange
)

type PDF struct {
	Paper           string
	mediaBox        *types.Rectangle
	Crop            string
	cropBox         *types.Rectangle
	Origin          string
	origin          types.Corner
	Guides          bool
	ContentBox      bool
	Debug           bool
	BackgroundColor string `json:"bgCol"`
	bgCol           *color.SimpleColor
	Fonts           map[string]*FormFont
	FormFonts       map[string]*FormFont
	FieldIDs        types.StringSet
	Fields          types.Array
	InheritedDA     string
	Header          *HorizontalBand
	Footer          *HorizontalBand
	Pages           map[string]*PDFPage
	pages           []*PDFPage
	Margin          *Margin
	Border          *Border
	Padding         *Padding
	Margins         map[string]*Margin
	Borders         map[string]*Border
	Paddings        map[string]*Padding
	SimpleBoxPool   map[string]*SimpleBox  `json:"boxes"`
	TextBoxPool     map[string]*TextBox    `json:"texts"`
	ImageBoxPool    map[string]*ImageBox   `json:"images"`
	TablePool       map[string]*Table      `json:"tables"`
	FieldGroupPool  map[string]*FieldGroup `json:"fieldgroups"`
	Colors          map[string]string
	colors          map[string]color.SimpleColor
	DirNames        map[string]string          `json:"dirs"`
	FileNames       map[string]string          `json:"files"`
	TimestampFormat string                     `json:"timestamp"`
	DateFormat      string                     `json:"dateFormat"`
	Conf            *model.Configuration       `json:"-"`
	XRefTable       *model.XRefTable           `json:"-"`
	Optimize        *model.OptimizationContext `json:"-"`
	FontResIDs      map[int]types.Dict         `json:"-"`
	XObjectResIDs   map[int]types.Dict         `json:"-"`
	CheckBoxAPs     map[float64]*AP            `json:"-"`
	RadioBtnAPs     map[float64]*AP            `json:"-"`
	HasForm         bool                       `json:"-"`
	OldFieldIDs     types.StringSet            `json:"-"`
	Offline         bool                       `json:"-"`
	Timeout         int                        `json:"-"`
	httpClient      *http.Client
}

func (pdf *PDF) Update() bool { _ = "STUB: not implemented"; return false }

func (pdf *PDF) pageCount() int { _ = "STUB: not implemented"; return 0 }

func (pdf *PDF) color(s string) *color.SimpleColor { _ = "STUB: not implemented"; return nil }

func (pdf *PDF) parseColor(s string) (*color.SimpleColor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pdf *PDF) resolveFileName(s string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (pdf *PDF) validatePageBoundaries() error { _ = "STUB: not implemented"; return nil }

func (pdf *PDF) validateOrigin() error { _ = "STUB: not implemented"; return nil }

func (pdf *PDF) validateColors() error { _ = "STUB: not implemented"; return nil }

func (pdf *PDF) validateFonts() error { _ = "STUB: not implemented"; return nil }

func (pdf *PDF) validateHeader() error { _ = "STUB: not implemented"; return nil }

func (pdf *PDF) validateFooter() error { _ = "STUB: not implemented"; return nil }

func (pdf *PDF) validateBorders() error { _ = "STUB: not implemented"; return nil }

func (pdf *PDF) validateMargins() error { _ = "STUB: not implemented"; return nil }

func (pdf *PDF) validatePaddings() error { _ = "STUB: not implemented"; return nil }

func (pdf *PDF) validateSimpleBoxPool() error { _ = "STUB: not implemented"; return nil }

func (pdf *PDF) validateTextBoxPool() error { _ = "STUB: not implemented"; return nil }

func (pdf *PDF) validateImageBoxPool() error { _ = "STUB: not implemented"; return nil }

func (pdf *PDF) validateTablePool() error { _ = "STUB: not implemented"; return nil }

func (pdf *PDF) validateFieldGroupPool() error { _ = "STUB: not implemented"; return nil }

func (pdf *PDF) validatePools() error { _ = "STUB: not implemented"; return nil }

func (pdf *PDF) validateBordersMarginsPaddings() error { _ = "STUB: not implemented"; return nil }

func (pdf *PDF) validatePages() ([]int, error) { _ = "STUB: not implemented"; return nil, nil }

func (pdf *PDF) Validate() error { _ = "STUB: not implemented"; return nil }

func (pdf *PDF) DuplicateField(ID string) bool { _ = "STUB: not implemented"; return false }

func (pdf *PDF) calcFont(f *FormFont) error { _ = "STUB: not implemented"; return nil }

func (pdf *PDF) newPageFontID(indRef *types.IndirectRef, nextInd, pageNr int) string {
	_ = "STUB: not implemented"
	return ""
}

func (pdf *PDF) idForFontName(fontName, fontLang string, pageFonts, globalFonts model.FontMap, pageNr int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func fontIndRef(xRefTable *model.XRefTable, fontName, fontLang string) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pdf *PDF) ensureFont(fontID, fontName, fontLang string, fonts model.FontMap) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pdf *PDF) ensureFormFont(font *FormFont) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (pdf *PDF) calcTopLevelFonts() { _ = "STUB: not implemented"; return }

func (pdf *PDF) calcInheritedPageFonts() { _ = "STUB: not implemented"; return }

func (pdf *PDF) calcInheritedContentFonts() { _ = "STUB: not implemented"; return }

func (pdf *PDF) calcInheritedFonts() { _ = "STUB: not implemented"; return }

func (pdf *PDF) calcInheritedMargins() { _ = "STUB: not implemented"; return }

func (pdf *PDF) calcInheritedBorders() { _ = "STUB: not implemented"; return }

func (pdf *PDF) calcInheritedPaddings() { _ = "STUB: not implemented"; return }

func (pdf *PDF) calcInheritedSimpleBoxes() { _ = "STUB: not implemented"; return }

func (pdf *PDF) calcInheritedTextBoxes() { _ = "STUB: not implemented"; return }

func (pdf *PDF) calcInheritedImageBoxes() { _ = "STUB: not implemented"; return }

func (pdf *PDF) calcInheritedTables() { _ = "STUB: not implemented"; return }

func (pdf *PDF) calcInheritedFieldGroups() { _ = "STUB: not implemented"; return }

func (pdf *PDF) calcInheritedAttrs() { _ = "STUB: not implemented"; return }

func (pdf *PDF) highlightPos(w io.Writer, x, y float64, cBox *types.Rectangle) {
	_ = "STUB: not implemented"
	return
}

func (pdf *PDF) renderPageBackground(page *PDFPage, w io.Writer) { _ = "STUB: not implemented"; return }

func (pdf *PDF) newModelPageforPDFPage(page *PDFPage) model.Page {
	_ = "STUB: not implemented"
	return *new(model.Page)
}

func (pdf *PDF) renderPageHeader(p *model.Page, pageNr int, fonts model.FontMap, images model.ImageMap) (float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (pdf *PDF) renderPageFooter(p *model.Page, pageNr int, fonts model.FontMap, images model.ImageMap) (float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (pdf *PDF) renderBlankPage(p *model.Page, pageNr int, fonts model.FontMap, images model.ImageMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (pdf *PDF) renderContentPage(page *PDFPage, p *model.Page, pageNr int, fonts model.FontMap, images model.ImageMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (pdf *PDF) RenderPages() ([]*model.Page, model.FontMap, error) {
	_ = "STUB: not implemented"
	return nil, *new(model.FontMap), nil
}
