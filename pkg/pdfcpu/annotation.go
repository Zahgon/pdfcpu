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

func validateAnnotationOperationContext(ctx *model.Context, incr bool) error {
	_ = "STUB: not implemented"
	return nil
}

func ensureAnnotationCache(ctx *model.Context) { _ = "STUB: not implemented"; return }

func validateAnnotationRenderer(ar model.AnnotationRenderer) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationPage(ctx *model.Context, pageNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func annotsArrayFromObject(o types.Object, pageNr int) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func dereferenceAnnotsArray(ctx *model.Context, indRef types.IndirectRef, pageNr int) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func pageDictForAnnotation(ctx *model.Context, pageNr int) (*types.IndirectRef, types.Dict, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.Dict), nil
}

func addAnnotationToCache(ctx *model.Context, ann model.AnnotationRenderer, pageNr, objNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func removeAnnotationFromCache(ctx *model.Context, pageNr, objNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func stripAnnotationBackReferences(d types.Dict) { _ = "STUB: not implemented"; return }

func deleteAnnotationObject(ctx *model.Context, o types.Object, pageNr int) error {
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

func validateAddAnnotationInput(
	ctx *model.Context,
	pageDictIndRef *types.IndirectRef,
	pageDict types.Dict,
	pageNr int,
	ar model.AnnotationRenderer,
	incr bool) error {
	_ = "STUB: not implemented"
	return nil
}

func addAnnotationToPageAnnots(
	ctx *model.Context,
	annotIndRef, pageDictIndRef *types.IndirectRef,
	pageDict types.Dict,
	pageNr int,
	ar model.AnnotationRenderer,
	incr bool) error {
	_ = "STUB: not implemented"
	return nil
}

func addAnnotationToIndirectAnnots(
	ctx *model.Context,
	annotsIndRef types.IndirectRef,
	annotIndRef *types.IndirectRef,
	pageNr int,
	ar model.AnnotationRenderer,
	incr bool) error {
	_ = "STUB: not implemented"
	return nil
}

func cleanupAddedAnnotation(ctx *model.Context, pageNr, objNr int, annotIndRef *types.IndirectRef) error {
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

type annotationAddPage struct {
	pageNr         int
	pageDictIndRef *types.IndirectRef
	pageDict       types.Dict
	annots         []model.AnnotationRenderer
}

func selectedAnnotationPageNrs(selectedPages types.IntSet) []int {
	_ = "STUB: not implemented"
	return nil
}

func prepareAnnotationAddPage(ctx *model.Context, pageNr int, annots []model.AnnotationRenderer) (annotationAddPage, error) {
	_ = "STUB: not implemented"
	return *new(annotationAddPage), nil
}

func prepareAnnotationAddPages(ctx *model.Context, selectedPages types.IntSet, ar model.AnnotationRenderer) ([]annotationAddPage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func prepareAnnotationAddMap(ctx *model.Context, m map[int][]model.AnnotationRenderer) ([]annotationAddPage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func applyAnnotationAddPages(ctx *model.Context, pages []annotationAddPage, incr bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func AddAnnotations(ctx *model.Context, selectedPages types.IntSet, ar model.AnnotationRenderer, incr bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func AddAnnotationsMap(ctx *model.Context, m map[int][]model.AnnotationRenderer, incr bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func validateAnnotationObjectForRemoval(ctx *model.Context, o types.Object, pageNr, index int) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationObjectsForRemoval(ctx *model.Context, annots types.Array, pageNr int) error {
	_ = "STUB: not implemented"
	return nil
}

type annotationDeletionValidator struct {
	ctx             *model.Context
	remainingRefs   map[int]int
	deletedObjNrSet types.IntSet
	pageNr          int
}

func validateAnnotationFreeList(ctx *model.Context, pageNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *annotationDeletionValidator) remainingRefCount(objNr int, entry *model.XRefTableEntry) int {
	_ = "STUB: not implemented"
	return 0
}

func (v *annotationDeletionValidator) validateIndirectRef(indRef types.IndirectRef, stripBackRefs bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *annotationDeletionValidator) validateDict(d types.Dict, stripBackRefs bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *annotationDeletionValidator) validateObject(o types.Object, stripBackRefs bool) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDeletionTargets(ctx *model.Context, targets []types.Object, pageNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func prepareRemoveAllAnnotations(ctx *model.Context, obj types.Object, pageNr int) (types.Array, *types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil, nil
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

type annotationRemovalTarget struct {
	index  int
	indRef types.IndirectRef
}

type annotationRemovalPlan struct {
	targets         map[int]annotationRemovalTarget
	matchedObjNrSet types.IntSet
}

func newAnnotationRemovalPlan() *annotationRemovalPlan { _ = "STUB: not implemented"; return nil }

func (p *annotationRemovalPlan) sortedTargets() []annotationRemovalTarget {
	_ = "STUB: not implemented"
	return nil
}

func (p *annotationRemovalPlan) deletionTargets() []types.Object {
	_ = "STUB: not implemented"
	return nil
}

func cachedAnnotationType(ctx *model.Context, pageNr, objNr int) (model.AnnotationType, error) {
	_ = "STUB: not implemented"
	return *new(model.AnnotationType), nil
}

func resolveAnnotationRemovalTarget(
	ctx *model.Context,
	pageNr int,
	annots types.Array,
	index int,
	expectedType *model.AnnotationType) (annotationRemovalTarget, error) {
	_ = "STUB: not implemented"
	return *new(annotationRemovalTarget), nil
}

func addAnnotationRemovalTarget(
	ctx *model.Context,
	plan *annotationRemovalPlan,
	pageNr int,
	annots types.Array,
	index int,
	expectedType *model.AnnotationType) error {
	_ = "STUB: not implemented"
	return nil
}

func preflightAnnotationTypes(
	ctx *model.Context,
	plan *annotationRemovalPlan,
	annotTypes []model.AnnotationType,
	pageNr int,
	annots types.Array) error {
	_ = "STUB: not implemented"
	return nil
}

func sortedSelectedAnnotationObjNrs(objNrSet types.IntSet) []int {
	_ = "STUB: not implemented"
	return nil
}

func preflightAnnotationObjNrs(
	ctx *model.Context,
	plan *annotationRemovalPlan,
	objNrSet types.IntSet,
	pageNr int,
	annots types.Array) error {
	_ = "STUB: not implemented"
	return nil
}

func preflightAnnotationID(
	ctx *model.Context,
	plan *annotationRemovalPlan,
	id string,
	matchedObjNr,
	pageNr int,
	annots types.Array) error {
	_ = "STUB: not implemented"
	return nil
}

func preflightAnnotationIDs(
	ctx *model.Context,
	plan *annotationRemovalPlan,
	ids []string,
	objNrSet types.IntSet,
	pageNr int,
	annots types.Array) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCachedAnnotationMembership(ctx *model.Context, pageNr int, objNrSet types.IntSet) error {
	_ = "STUB: not implemented"
	return nil
}

func preflightSelectiveAnnotationArray(ctx *model.Context, pageNr int, annots types.Array) error {
	_ = "STUB: not implemented"
	return nil
}

func preflightSelectiveAnnotationRemoval(
	ctx *model.Context,
	annotTypes []model.AnnotationType,
	ids []string,
	objNrSet types.IntSet,
	pageNr int,
	annots types.Array) (*annotationRemovalPlan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func applySelectiveAnnotationRemoval(
	ctx *model.Context,
	plan *annotationRemovalPlan,
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

func annotationRemovalCatalog(ctx *model.Context, removeAll bool) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func RemoveAnnotations(ctx *model.Context, selectedPages types.IntSet, idsAndTypes []string, objNrs []int, incr bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
