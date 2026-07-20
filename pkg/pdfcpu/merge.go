package pdfcpu

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func EnsureOutlines(ctx *model.Context, fName string, append bool) error {
	_ = "STUB: not implemented"
	return nil
}

func mergeOutlinesWrapped(fName string, p int, ctxSrc, ctxDest *model.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func destOutlines(ctxDest *model.Context) (*types.IndirectRef, types.Dict, *types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.Dict), nil, nil
}

func outlineTopCount(outlinesDict types.Dict) int { _ = "STUB: not implemented"; return 0 }

func appendOutlineWrapper(ctxDest *model.Context, outlinesDict types.Dict, indRef, oldLast *types.IndirectRef, fName string, p int) (*types.IndirectRef, types.Dict, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.Dict), nil
}

func ensureOutlinesRoot(ctx *model.Context) (*types.IndirectRef, types.Dict, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.Dict), nil
}

func sourceOutlines(ctxSrc, ctxDest *model.Context) (*types.IndirectRef, *types.IndirectRef, int, error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, nil
}

func reparentOutlineItems(ctx *model.Context, first, parent *types.IndirectRef) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func mergeOutlinesPreserve(ctxSrc, ctxDest *model.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func handleNeedAppearances(ctxSrc *model.Context, dSrc, dDest types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func handleCO(ctxSrc, ctxDest *model.Context, dSrc, dDest types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func handleDR(ctxSrc *model.Context, dSrc, dDest types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func handleDA(ctxSrc *model.Context, dSrc, dDest types.Dict, arrFieldsSrc types.Array) error {
	_ = "STUB: not implemented"
	return nil
}

func handleQ(ctxSrc *model.Context, dSrc, dDest types.Dict, arrFieldsSrc types.Array) error {
	_ = "STUB: not implemented"
	return nil
}

func handleFormAttributes(ctxSrc, ctxDest *model.Context, dSrc, dDest types.Dict, arrFieldsSrc types.Array) error {
	_ = "STUB: not implemented"
	return nil
}

func rootDicts(ctxSrc, ctxDest *model.Context) (types.Dict, types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), *new(types.Dict), nil
}

func mergeInFields(ctxDest *model.Context, arrFieldsSrc, arrFieldsDest types.Array, dDest types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func mergeDests(ctxSource, ctxDest *model.Context) error { _ = "STUB: not implemented"; return nil }

func mergeNames(ctxSrc, ctxDest *model.Context) error { _ = "STUB: not implemented"; return nil }

func mergeForms(ctxSrc, ctxDest *model.Context) error { _ = "STUB: not implemented"; return nil }

func patchIndRef(ir *types.IndirectRef, lookup map[int]int) { _ = "STUB: not implemented"; return }

func patchObject(o types.Object, lookup map[int]int) types.Object {
	_ = "STUB: not implemented"
	return *new(types.Object)
}

func patchDict(d types.Dict, lookup map[int]int) { _ = "STUB: not implemented"; return }

func patchArray(a *types.Array, lookup map[int]int) { _ = "STUB: not implemented"; return }

func objNrsIntSet(ctx *model.Context) types.IntSet {
	_ = "STUB: not implemented"
	return *new(types.IntSet)
}

func lookupTable(keys types.IntSet, i int) map[int]int { _ = "STUB: not implemented"; return nil }

func patchObjects(s types.IntSet, lookup map[int]int) types.IntSet {
	_ = "STUB: not implemented"
	return *new(types.IntSet)
}

func patchNameTree(n *model.Node, lookup map[int]int) error { _ = "STUB: not implemented"; return nil }

func patchSourceObjectNumbers(ctxSrc, ctxDest *model.Context) { _ = "STUB: not implemented"; return }

func createDividerPagesDict(ctx *model.Context, parentIndRef types.IndirectRef) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func appendSourcePageTreeToDestPageTreeOrig(ctxSrc, ctxDest *model.Context, dividerPage bool) error {
	_ = "STUB: not implemented"
	return nil
}

func pageTreeRoot(ctx *model.Context) (*types.IndirectRef, types.Dict, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.Dict), nil
}

func hasInheritedPageAttrs(d types.Dict) bool { _ = "STUB: not implemented"; return false }

func ensureNeutralPageTreeRoot(ctx *model.Context, indRef *types.IndirectRef, d types.Dict) (*types.IndirectRef, types.Dict, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.Dict), nil
}

func pageTreeKids(d types.Dict, indRef types.IndirectRef) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func appendSourcePageTreeToDestPageTree(ctxSrc, ctxDest *model.Context, dividerPage bool) error {
	_ = "STUB: not implemented"
	return nil
}

func zipSourcePageTreeIntoDestPageTree(ctxSrc, ctxDest *model.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func appendSourceObjectsToDest(ctxSrc, ctxDest *model.Context) { _ = "STUB: not implemented"; return }

func mergeIntSets(src, dest types.IntSet) { _ = "STUB: not implemented"; return }

func mergeDuplicateObjNumberIntSets(ctxSrc, ctxDest *model.Context) {
	_ = "STUB: not implemented"
	return
}

func MergeXRefTables(fName string, ctxSrc, ctxDest *model.Context, zip, dividerPage bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}
