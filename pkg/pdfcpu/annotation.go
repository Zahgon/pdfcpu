package pdfcpu

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func CachedAnnotationObjNrs(ctx *model.Context) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sortedPageNrsForAnnotsFromCache(ctx *model.Context) []int {
	_ = "STUB: not implemented"
	return nil
}

func addAnnotationToCache(ctx *model.Context, ann model.AnnotationRenderer, pageNr, objNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func removeAnnotationFromCache(ctx *model.Context, pageNr, objNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func findAnnotByID(ctx *model.Context, id string, annots types.Array) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func findAnnotByObjNr(objNr int, annots types.Array) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func createAnnot(ctx *model.Context, ar model.AnnotationRenderer, pageIndRef *types.IndirectRef) (*types.IndirectRef, types.Dict, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.Dict), nil
}

func linkAnnotation(xRefTable *model.XRefTable, d types.Dict, r *types.Rectangle, apObjNr int, contents, nm string, f model.AnnotationFlags) (model.AnnotationRenderer, error) {
	_ = "STUB: not implemented"
	return *new(model.AnnotationRenderer), nil
}

func Annotation(xRefTable *model.XRefTable, d types.Dict) (model.AnnotationRenderer, error) {
	_ = "STUB: not implemented"
	return *new(model.AnnotationRenderer), nil
}

func AnnotationsForSelectedPages(ctx *model.Context, selectedPages types.IntSet) map[int]model.PgAnnots {
	_ = "STUB: not implemented"
	return nil
}

func prepareHeader(horSep *[]int, maxLen *AnnotListMaxLengths, customAnnot bool) string {
	_ = "STUB: not implemented"
	return ""
}

type AnnotListMaxLengths struct {
	ObjNr, ID, Rect, Content, Type int
}

type AnnotationList struct {
	Header      Header                        `json:"header"`
	Annotations map[int][]AnnotationListEntry `json:"annotations"`
}

type AnnotationListEntry struct {
	Type      string     `json:"type"`
	ObjNr     int        `json:"objNr"`
	ID        string     `json:"id,omitempty"`
	Rect      [4]float64 `json:"rect"`
	Content   *string    `json:"content"`
	CustomTyp string     `json:"customType,omitempty"`
}

func ListAnnotations(annots map[int]model.PgAnnots) (int, []string, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func annotationContent(ann model.AnnotationRenderer) *string { _ = "STUB: not implemented"; return nil }

func annotationListEntry(annType string, objNr int, ann model.AnnotationRenderer) AnnotationListEntry {
	_ = "STUB: not implemented"
	return *new(AnnotationListEntry)
}

func annotationJSONHeader() Header { _ = "STUB: not implemented"; return *new(Header) }

func sortedAnnotationObjNrs(annots model.Annot) []int { _ = "STUB: not implemented"; return nil }

func addAnnotationListEntries(list *AnnotationList, pageNr int, annType string, annots model.Annot) int {
	_ = "STUB: not implemented"
	return 0
}

func annotationList(annots map[int]model.PgAnnots) (AnnotationList, int) {
	_ = "STUB: not implemented"
	return *new(AnnotationList), 0
}

func sortedAnnotationPages(annots map[int]model.PgAnnots) []int {
	_ = "STUB: not implemented"
	return nil
}

func sortedAnnotationTypeNames(pageAnnots model.PgAnnots) []string {
	_ = "STUB: not implemented"
	return nil
}

func ListAnnotationsJSON(annots map[int]model.PgAnnots) (int, []string, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func addAnnotationToDirectObj(
	ctx *model.Context,
	annots types.Array,
	annotIndRef, pageDictIndRef *types.IndirectRef,
	pageDict types.Dict,
	pageNr int,
	ar model.AnnotationRenderer,
	incr bool) error {
	_ = "STUB: not implemented"
	return nil
}

func AddAnnotation(
	ctx *model.Context,
	pageDictIndRef *types.IndirectRef,
	pageDict types.Dict,
	pageNr int,
	ar model.AnnotationRenderer,
	incr bool) (*types.IndirectRef, types.Dict, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.Dict), nil
}

func AddAnnotationToPage(ctx *model.Context, pageNr int, ar model.AnnotationRenderer, incr bool) (*types.IndirectRef, types.Dict, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.Dict), nil
}

func AddAnnotations(ctx *model.Context, selectedPages types.IntSet, ar model.AnnotationRenderer, incr bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func AddAnnotationsMap(ctx *model.Context, m map[int][]model.AnnotationRenderer, incr bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func removeAllAnnotations(
	ctx *model.Context,
	pageDict types.Dict,
	pageDictObjNr,
	pageNr int,
	incr bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func removeAnnotationsByType(
	ctx *model.Context,
	annotTypes []model.AnnotationType,
	pageNr int,
	annots types.Array,
	incr bool) (types.Array, bool, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), false, nil
}

func removeAnnotationByID(
	ctx *model.Context,
	id string,
	pageNr int,
	annots types.Array,
	incr bool) (types.Array, bool, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), false, nil
}

func removeAnnotationsByID(
	ctx *model.Context,
	ids []string,
	objNrSet types.IntSet,
	pageNr int,
	annots types.Array,
	incr bool) (types.Array, bool, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), false, nil
}

func removeAnnotationsByObjNr(
	ctx *model.Context,
	objNrSet types.IntSet,
	pageNr int,
	annots types.Array,
	incr bool) (types.Array, bool, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), false, nil
}

func removeAnnotationsFromAnnots(
	ctx *model.Context,
	annotTypes []model.AnnotationType,
	ids []string,
	objNrSet types.IntSet,
	pageNr int,
	annots types.Array,
	incr bool) (types.Array, bool, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), false, nil
}

func removeAnnotationsFromIndAnnots(ctx *model.Context,
	annotTypes []model.AnnotationType,
	ids []string,
	objNrSet types.IntSet,
	pageNr int,
	annots types.Array,
	incr bool,
	pageDict types.Dict,
	pageDictObjNr int,
	indRef types.IndirectRef) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func RemoveAnnotationsFromPageDict(
	ctx *model.Context,
	annotTypes []model.AnnotationType,
	ids []string,
	objNrSet types.IntSet,
	pageDict types.Dict,
	pageDictObjNr,
	pageNr int,
	incr bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func prepForRemoveAnnotations(ctx *model.Context, idsAndTypes []string, objNrs []int, incr bool) ([]model.AnnotationType, []string, types.IntSet, bool) {
	_ = "STUB: not implemented"
	return nil, nil, *new(types.IntSet), false
}

func RemoveAnnotations(ctx *model.Context, selectedPages types.IntSet, idsAndTypes []string, objNrs []int, incr bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
