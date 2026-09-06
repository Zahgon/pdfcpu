package primitives

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type PDFPage struct {
	pdf             *PDF
	number          int
	Paper           string
	mediaBox        *types.Rectangle
	Crop            string
	cropBox         *types.Rectangle
	BackgroundColor string `json:"bgCol"`
	bgCol           *color.SimpleColor
	Fonts           map[string]*FormFont
	DA              types.Object
	Guides          []*Guide
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
	FileNames       map[string]string      `json:"files"`
	Tabs            types.IntSet           `json:"-"`
	Content         *Content
}

func (page *PDFPage) resolveFileName(s string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (page *PDFPage) validatePageBoundaries() error { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) validateBackgroundColor() error { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) validateFonts() error { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) validateBorders() error { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) validateMargins() error { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) validatePaddings() error { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) validateSimpleBoxPool() error { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) validateTextBoxPool() error { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) validateImageBoxPool() error { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) validateTablePool() error { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) validateFieldGroupPool() error { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) validatePools() error { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) validate() error { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) namedFont(id string) *FormFont { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) namedMargin(id string) *Margin { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) namedBorder(id string) *Border { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) namedPadding(id string) *Padding { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) namedSimpleBox(id string) *SimpleBox { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) namedImageBox(id string) *ImageBox { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) namedTextBox(id string) *TextBox { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) namedTable(id string) *Table { _ = "STUB: not implemented"; return nil }

func (page *PDFPage) namedFieldGroup(id string) *FieldGroup { _ = "STUB: not implemented"; return nil }
