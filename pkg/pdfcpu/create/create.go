package create

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/primitives"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func ensureFontIndRef(xRefTable *model.XRefTable, fontName string, frPage model.FontResource, fonts model.FontMap) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addPageResources(xRefTable *model.XRefTable, d types.Dict, p model.Page, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func updatePageResources(xRefTable *model.XRefTable, d, resDict types.Dict, p model.Page, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func setAnnotationParentsAndFields(xRefTable *model.XRefTable, p *model.Page, pIndRef types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func addAnnotations(ff []model.FieldAnnotation, m map[int]model.FieldAnnotation) types.Array {
	_ = "STUB: not implemented"
	return *new(types.Array)
}

func mergeAnnotations(oldAnnots types.Array, ff []model.FieldAnnotation, m map[int]model.FieldAnnotation) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func CreatePage(
	xRefTable *model.XRefTable,
	parentPageIndRef types.IndirectRef,
	p *model.Page,
	fonts model.FontMap) (*types.IndirectRef, types.Dict, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.Dict), nil
}

func UpdatePage(xRefTable *model.XRefTable, dIndRef types.IndirectRef, d, res types.Dict, p *model.Page, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func cacheFormFieldIDs(ctx *model.Context, pdf *primitives.PDF) error {
	_ = "STUB: not implemented"
	return nil
}

func cacheResIDs(ctx *model.Context, pdf *primitives.PDF) error {
	_ = "STUB: not implemented"
	return nil
}

func parseFromJSON(ctx *model.Context, bb []byte) (*primitives.PDF, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func appendPage(
	ctx *model.Context,
	pagesDictIndRef types.IndirectRef,
	pagesDict types.Dict,
	p *model.Page,
	fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func updatePage(ctx *model.Context, pageNr int, p *model.Page, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func UpdatePageTree(ctx *model.Context, pages []*model.Page, fontMap model.FontMap) (types.Array, model.FontMap, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), *new(model.FontMap), nil
}

func prepareFormFontResDict(ctx *model.Context, pdf *primitives.PDF, fonts model.FontMap) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func createForm(
	ctx *model.Context,
	pdf *primitives.PDF,
	fields types.Array,
	fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func updateForm(
	ctx *model.Context,
	pdf *primitives.PDF,
	fields types.Array,
	fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func handleForm(
	ctx *model.Context,
	pdf *primitives.PDF,
	fields types.Array,
	fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func FromJSON(ctx *model.Context, rd io.Reader) error { _ = "STUB: not implemented"; return nil }
