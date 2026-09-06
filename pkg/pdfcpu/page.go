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

func validateAddPagesInputs(
	ctxSrc, ctxDest *model.Context,
	pagesDict types.Dict,
	fieldsSrc, fieldsDest *types.Array,
	migrated map[int]int) error {
	_ = "STUB: not implemented"
	return nil
}

func migratedPageDict(ctxSrc, ctxDest *model.Context, pageNr int, migrated map[int]int) (types.Dict, *types.IndirectRef, *model.InheritedPageAttrs, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil, nil, nil
}

func addPage(
	ctxSrc, ctxDest *model.Context,
	pageNr int,
	pagesIndRef types.IndirectRef,
	pagesDict types.Dict,
	fieldsSrc, fieldsDest *types.Array,
	migrated map[int]int) (*types.IndirectRef, error) {
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

func destinationPageMigrated(arr types.Array, migrated map[int]int) bool {
	_ = "STUB: not implemented"
	return false
}

func migrateNamedDestArray(arr types.Array, migrated map[int]int) bool {
	_ = "STUB: not implemented"
	return false
}

func migrateNamedDestValue(xRefTable *model.XRefTable, v *types.Object, migrated map[int]int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func migrateNamedDests(ctxSrc *model.Context, n *model.Node, migrated map[int]int) error {
	_ = "STUB: not implemented"
	return nil
}

func AddPages(ctxSrc, ctxDest *model.Context, pageNrs []int, usePgCache bool) error {
	_ = "STUB: not implemented"
	return nil
}
