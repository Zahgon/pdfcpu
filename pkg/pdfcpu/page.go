package pdfcpu

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

var pParamMap = parameterMap[PageConfiguration]{
	"dimensions": parseDimensions,
	"formsize":   parsePageFormat,
	"papersize":  parsePageFormat,
}

type PageConfiguration struct {
	PageDim  *types.Dim
	PageSize string
	UserDim  bool
	InpUnit  types.DisplayUnit
}

func DefaultPageConfiguration() *PageConfiguration { _ = "STUB: not implemented"; return nil }

func (p PageConfiguration) String() string { _ = "STUB: not implemented"; return "" }

func parsePageFormat(s string, p *PageConfiguration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func parseDimensions(s string, p *PageConfiguration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ParsePageConfiguration(s string, u types.DisplayUnit) (*PageConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addPages(
	ctxSrc, ctxDest *model.Context,
	pageNrs []int,
	usePgCache bool,
	pagesIndRef types.IndirectRef,
	pagesDict types.Dict,
	fieldsSrc, fieldsDest *types.Array,
	migrated map[int]int) error {
	_ = "STUB: not implemented"
	return nil
}

func migrateNamedDests(ctxSrc *model.Context, n *model.Node, migrated map[int]int) error {
	_ = "STUB: not implemented"
	return nil
}

func AddPages(ctxSrc, ctxDest *model.Context, pageNrs []int, usePgCache bool) error {
	_ = "STUB: not implemented"
	return nil
}
