package pdfcpu

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

var impParamMap = parameterMap[Import]{
	"dimensions":      parseDimensionsImp,
	"dpi":             parseDPI,
	"formsize":        parsePageFormatImp,
	"papersize":       parsePageFormatImp,
	"position":        parsePositionAnchorImp,
	"offset":          parsePositionOffsetImp,
	"scalefactor":     parseScaleFactorImp,
	"gray":            parseGray,
	"sepia":           parseSepia,
	"backgroundcolor": parseImportBackgroundColor,
	"bgcolor":         parseImportBackgroundColor,
}

type Import struct {
	PageDim  *types.Dim
	PageSize string
	UserDim  bool
	DPI      int
	Pos      types.Anchor
	Dx, Dy   float64
	Scale    float64
	ScaleAbs bool
	InpUnit  types.DisplayUnit
	Gray     bool
	Sepia    bool
	BgColor  *color.SimpleColor
}

func DefaultImportConfig() *Import { _ = "STUB: not implemented"; return nil }

func (imp Import) String() string { _ = "STUB: not implemented"; return "" }

func parsePageFormatImp(s string, imp *Import) (err error) { _ = "STUB: not implemented"; return nil }

func ParsePageDim(v string, u types.DisplayUnit) (*types.Dim, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func parseDimensionsImp(s string, imp *Import) (err error) { _ = "STUB: not implemented"; return nil }

func parsePositionAnchorImp(s string, imp *Import) error { _ = "STUB: not implemented"; return nil }

func parsePositionOffsetImp(s string, imp *Import) error { _ = "STUB: not implemented"; return nil }

func parseScaleFactorImp(s string, imp *Import) (err error) { _ = "STUB: not implemented"; return nil }

func parseDPI(s string, imp *Import) (err error) { _ = "STUB: not implemented"; return nil }

func parseGray(s string, imp *Import) error { _ = "STUB: not implemented"; return nil }

func parseSepia(s string, imp *Import) error { _ = "STUB: not implemented"; return nil }

func parseImportBackgroundColor(s string, imp *Import) error { _ = "STUB: not implemented"; return nil }

func ParseImportDetails(s string, u types.DisplayUnit) (*Import, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func importImagePDFBytes(wr io.Writer, pageDim *types.Dim, imgWidth, imgHeight float64, imp *Import) {
	_ = "STUB: not implemented"
	return
}

func NewPagesForImage(xRefTable *model.XRefTable, r io.Reader, parentIndRef *types.IndirectRef, imp *Import) ([]*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
