package pdfcpu

import (
	"bytes"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"github.com/pkg/errors"
)

var (
	errInvalidGridDims  = errors.New("pdfcpu grid: dimensions must be: m > 0, n > 0")
	errInvalidNUpConfig = errors.New("pdfcpu: invalid configuration string")
)

var (
	NUpValues = []int{2, 3, 4, 6, 8, 9, 12, 16}
	nUpDims   = map[int]types.Dim{
		2:  {Width: 2, Height: 1},
		3:  {Width: 3, Height: 1},
		4:  {Width: 2, Height: 2},
		6:  {Width: 3, Height: 2},
		8:  {Width: 4, Height: 2},
		9:  {Width: 3, Height: 3},
		12: {Width: 4, Height: 3},
		16: {Width: 4, Height: 4},
	}
)

var nupParamMap = parameterMap[model.NUp]{
	"dimensions":      parseDimensionsNUp,
	"formsize":        parsePageFormatNUp,
	"papersize":       parsePageFormatNUp,
	"orientation":     parseOrientation,
	"border":          parseElementBorder,
	"cropboxborder":   parseElementBorderOnCropbox,
	"margin":          parseElementMargin,
	"backgroundcolor": parseSheetBackgroundColor,
	"bgcolor":         parseSheetBackgroundColor,
	"guides":          parseBookletGuides,
	"multifolio":      parseBookletMultifolio,
	"foliosize":       parseBookletFolioSize,
	"btype":           parseBookletType,
	"binding":         parseBookletBinding,
	"enforce":         parseEnforce,
}

func parsePageFormatNUp(s string, nup *model.NUp) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func parseDimensionsNUp(s string, nup *model.NUp) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func parseOrientation(s string, nup *model.NUp) error { _ = "STUB: not implemented"; return nil }

func parseEnforce(s string, nup *model.NUp) error { _ = "STUB: not implemented"; return nil }

func parseElementBorder(s string, nup *model.NUp) error { _ = "STUB: not implemented"; return nil }

func parseElementBorderOnCropbox(s string, nup *model.NUp) error {
	_ = "STUB: not implemented"
	return nil
}

func parseBookletGuides(s string, nup *model.NUp) error { _ = "STUB: not implemented"; return nil }

func parseBookletMultifolio(s string, nup *model.NUp) error { _ = "STUB: not implemented"; return nil }

func parseBookletFolioSize(s string, nup *model.NUp) error { _ = "STUB: not implemented"; return nil }

func parseBookletType(s string, nup *model.NUp) error { _ = "STUB: not implemented"; return nil }

func parseBookletBinding(s string, nup *model.NUp) error { _ = "STUB: not implemented"; return nil }

func parseElementMargin(s string, nup *model.NUp) error { _ = "STUB: not implemented"; return nil }

func parseSheetBackgroundColor(s string, nup *model.NUp) error {
	_ = "STUB: not implemented"
	return nil
}

func ParseNUpDetails(s string, nup *model.NUp) error { _ = "STUB: not implemented"; return nil }

func PDFNUpConfig(val int, desc string, conf *model.Configuration) (*model.NUp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ImageNUpConfig(val int, desc string, conf *model.Configuration) (*model.NUp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PDFGridConfig(rows, cols int, desc string, conf *model.Configuration) (*model.NUp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ImageGridConfig(rows, cols int, desc string, conf *model.Configuration) (*model.NUp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseNUpValue(n int, nUp *model.NUp) error { _ = "STUB: not implemented"; return nil }

func ParseNUpGridDefinition(rows, cols int, nUp *model.NUp) error {
	_ = "STUB: not implemented"
	return nil
}

func nUpImagePDFBytes(w io.Writer, imgWidth, imgHeight int, nup *model.NUp, formResID string) {
	_ = "STUB: not implemented"
	return
}

func createNUpFormForImage(xRefTable *model.XRefTable, imgIndRef *types.IndirectRef, w, h, i int) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewNUpPageForImage(xRefTable *model.XRefTable, fileName string, parentIndRef *types.IndirectRef, nup *model.NUp) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NUpFromOneImage(ctx *model.Context, fileName string, nup *model.NUp, pagesDict types.Dict, pagesIndRef *types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func wrapUpPage(ctx *model.Context, nup *model.NUp, d types.Dict, buf bytes.Buffer, pagesDict types.Dict, pagesIndRef *types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func nupPageNumber(i int, sortedPageNumbers []int) int { _ = "STUB: not implemented"; return 0 }

func sortSelectedPages(pages types.IntSet) []int { _ = "STUB: not implemented"; return nil }

func nupPages(
	ctx *model.Context,
	selectedPages types.IntSet,
	nup *model.NUp,
	pagesDict types.Dict,
	pagesIndRef *types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func NUpFromMultipleImages(ctx *model.Context, fileNames []string, nup *model.NUp, pagesDict types.Dict, pagesIndRef *types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func NUpFromPDF(ctx *model.Context, selectedPages types.IntSet, nup *model.NUp) error {
	_ = "STUB: not implemented"
	return nil
}
