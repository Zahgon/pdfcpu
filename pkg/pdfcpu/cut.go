package pdfcpu

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func ParseCutConfigForPoster(s string, u types.DisplayUnit) (*model.Cut, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseCutConfigForN(n int, s string, u types.DisplayUnit) (*model.Cut, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseCutConfig(s string, u types.DisplayUnit) (*model.Cut, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func drawOutlineCuts(w io.Writer, cropBox, cb *types.Rectangle, cut *model.Cut) {
	_ = "STUB: not implemented"
	return
}

func createOutline(
	ctxSrc, ctxDest *model.Context,
	pagesIndRef types.IndirectRef,
	pagesDict, d types.Dict,
	pageNr int,
	cropBox *types.Rectangle,
	migrated map[int]int,
	cut *model.Cut) error {
	_ = "STUB: not implemented"
	return nil
}

func prepForCut(ctxSrc *model.Context, pageNr int) (
	*model.Context,
	*types.Rectangle,
	*types.IndirectRef,
	types.Dict,
	types.Dict,
	*model.InheritedPageAttrs,
	error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, *new(types.Dict), *new(types.Dict), nil, nil
}

func internPageRot(ctxSrc *model.Context, rotate int, cropBox *types.Rectangle, d types.Dict, pageNr int, trans []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func handleCutMargin(ctxSrc *model.Context, d, d1 types.Dict, pageNr int, cropBox, cb *types.Rectangle, i, j int, w, h float64, sc *float64, cut *model.Cut) error {
	_ = "STUB: not implemented"
	return nil
}

func createTiles(
	ctxSrc, ctxDest *model.Context,
	pagesIndRef types.IndirectRef,
	pagesDict, d types.Dict,
	pageNr int,
	cropBox *types.Rectangle,
	inhPAttrs *model.InheritedPageAttrs,
	migrated map[int]int,
	cut *model.Cut) error {
	_ = "STUB: not implemented"
	return nil
}

func CutPage(ctxSrc *model.Context, pageNr int, cut *model.Cut) (*model.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createNDownCuts(n int, cropBox *types.Rectangle, cut *model.Cut) {
	_ = "STUB: not implemented"
	return
}

func NDownPage(ctxSrc *model.Context, pageNr, n int, cut *model.Cut) (*model.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createPosterCuts(cropBox *types.Rectangle, cut *model.Cut) { _ = "STUB: not implemented"; return }

func PosterPage(ctxSrc *model.Context, pageNr int, cut *model.Cut) (*model.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
