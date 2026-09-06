package pdfcpu

import (
	"crypto/sha256"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func optimizeContentStreamUsage(ctx *model.Context, sd *types.StreamDict, objNr int) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func removeEmptyContentStreams(ctx *model.Context, pageDict types.Dict, obj types.Object, pageObjNumber int) error {
	_ = "STUB: not implemented"
	return nil
}

func optimizePageContent(ctx *model.Context, pageDict types.Dict, pageObjNumber int) error {
	_ = "STUB: not implemented"
	return nil
}

func resourcesDictForPageDict(xRefTable *model.XRefTable, pageDict types.Dict, pageObjNumber int) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func handleDuplicateFontObject(ctx *model.Context, fontDict types.Dict, fName, rName string, objNr, pageNr int) (*int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func pageImages(ctx *model.Context, pageNr int) types.IntSet {
	_ = "STUB: not implemented"
	return *new(types.IntSet)
}

func pageFonts(ctx *model.Context, pageNr int) types.IntSet {
	_ = "STUB: not implemented"
	return *new(types.IntSet)
}

func registerFontDictObjNr(ctx *model.Context, fName string, objNr int) {
	_ = "STUB: not implemented"
	return
}

func checkForEmbeddedFont(ctx *model.Context) bool { _ = "STUB: not implemented"; return false }

func qualifiedRName(rNamePrefix, rName string) string { _ = "STUB: not implemented"; return "" }

func optimizeFontResourcesDict(ctx *model.Context, rDict types.Dict, pageNr int, rNamePrefix string) error {
	_ = "STUB: not implemented"
	return nil
}

func imageObjectHashes(ctx *model.Context) map[[sha256.Size]byte][]int {
	_ = "STUB: not implemented"
	return nil
}

func handleDuplicateImageObject(ctx *model.Context, imageDict *types.StreamDict, resourceName string, objNr, pageNr int) (*int, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func optimizeXObjectImage(ctx *model.Context, osd *types.StreamDict, rNamePrefix, rName string, rDict types.Dict, objNr, pageNr, pageObjNumber int, pageImages types.IntSet) error {
	_ = "STUB: not implemented"
	return nil
}

func optimizeXObjectForm(ctx *model.Context, sd *types.StreamDict, objNr int) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func optimizeFormResources(ctx *model.Context, o types.Object, pageNr, pageObjNumber int, rName string, visitedRes []types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func visited(o types.Object, visited []types.Object) bool { _ = "STUB: not implemented"; return false }

func formResourcesVisited(ctx *model.Context, pageNr, objNr int) bool {
	_ = "STUB: not implemented"
	return false
}

func optimizeForm(ctx *model.Context, osd *types.StreamDict, rNamePrefix, rName string, rDict types.Dict, objNr, pageNr, pageObjNumber int, vis []types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func optimizeExtGStateResources(ctx *model.Context, rDict types.Dict, pageNr, pageObjNumber int, rNamePrefix string, vis []types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func optimizeSMaskResources(dict types.Dict, vis []types.Object, rNamePrefix string, ctx *model.Context, rDict types.Dict, pageNr int, pageImages types.IntSet, pageObjNumber int) error {
	_ = "STUB: not implemented"
	return nil
}

func optimizeExtGStateResourcesDict(ctx *model.Context, rDict types.Dict, pageNr, pageObjNumber int, rNamePrefix string, vis []types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func optimizeXObjectResource(ctx *model.Context, sd *types.StreamDict, rDict types.Dict, rNamePrefix, rName string,
	qualifiedRName string, objNr, pageNr, pageObjNumber int, pageImages types.IntSet, vis []types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func optimizeXObjectResourcesDict(ctx *model.Context, rDict types.Dict, pageNr, pageObjNumber int, rNamePrefix string, vis []types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func processFontResources(ctx *model.Context, obj types.Object, pageNr, pageObjNumber int, rNamePrefix string) error {
	_ = "STUB: not implemented"
	return nil
}

func processXObjectResources(ctx *model.Context, obj types.Object, pageNr, pageObjNumber int, rNamePrefix string, visitedRes []types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func processExtGStateResources(ctx *model.Context, obj types.Object, pageNr, pageObjNumber int, rNamePrefix string, visitedRes []types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func optimizeResources(ctx *model.Context, resourcesDict types.Dict, pageNr, pageObjNumber int, rNamePrefix string, visitedRes []types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func parseResourcesDict(ctx *model.Context, pageDict types.Dict, pageNr, pageObjNumber int) error {
	_ = "STUB: not implemented"
	return nil
}

func parsePageTreeKid(ctx *model.Context, v types.Object, kidNr, pageNr int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func parsePagesDict(ctx *model.Context, pagesDict types.Dict, pageNr int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func traverse(xRefTable *model.XRefTable, value types.Object, duplObjs types.IntSet) error {
	_ = "STUB: not implemented"
	return nil
}

func traverseObjectGraphAndMarkDuplicates(xRefTable *model.XRefTable, obj types.Object, duplObjs types.IntSet) error {
	_ = "STUB: not implemented"
	return nil
}

func calcRedundantObjects(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func fixCorruptFontResDicts(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func optimizeFontAndImages(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func streamLengthFontFile(xRefTable *model.XRefTable, indirectRef *types.IndirectRef) (*int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func calcEmbeddedFontsMemoryUsage(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func fontDescriptorFontFileIndirectObjectRef(fontDescriptorDict types.Dict) *types.IndirectRef {
	_ = "STUB: not implemented"
	return nil
}

func processFontFilesForFontDict(xRefTable *model.XRefTable, fontDict types.Dict, objectNumber int, indRefsMap map[types.IndirectRef]bool) error {
	_ = "STUB: not implemented"
	return nil
}

func calcRedundantEmbeddedFontsMemoryUsage(ctx *model.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func calcFontBinarySizes(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func calcImageBinarySizes(ctx *model.Context) { _ = "STUB: not implemented"; return }

func calcBinarySizes(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func fixDeepDict(ctx *model.Context, d types.Dict) error { _ = "STUB: not implemented"; return nil }

func fixDeepArray(ctx *model.Context, a types.Array) error { _ = "STUB: not implemented"; return nil }

func fixDirectObject(ctx *model.Context, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func fixIndirectObject(ctx *model.Context, ir *types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func fixDeepObject(ctx *model.Context, o types.Object) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fixReferencesToFreeObjects(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func CacheFormFonts(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func optimizeResourceDicts(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func resolveWidth(ctx *model.Context, sd *types.StreamDict) error {
	_ = "STUB: not implemented"
	return nil
}

func ensureDirectWidthForXObjs(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func OptimizeXRefTable(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }
