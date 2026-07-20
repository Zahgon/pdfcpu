package primitives

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type Regions struct {
	page        *PDFPage
	parent      *Content
	Name        string
	Orientation string `json:"orient"`
	horizontal  bool
	Divider     *Divider `json:"div"`
	Left, Right *Content
	Top, Bottom *Content
	mediaBox    *types.Rectangle
}

func parseRegionOrientation(s string) (types.Orientation, error) {
	_ = "STUB: not implemented"
	return *new(types.Orientation), nil
}

func (r *Regions) validate() error { _ = "STUB: not implemented"; return nil }

func (r *Regions) render(p *model.Page, pageNr int, fonts model.FontMap, images model.ImageMap) error {
	_ = "STUB: not implemented"
	return nil
}
