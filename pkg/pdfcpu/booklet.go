package pdfcpu

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"github.com/pkg/errors"
)

var errInvalidBookletAdvanced = errors.New("pdfcpu booklet advanced cannot have binding along the top (portrait short-edge, landscape long-edge). use plain booklet instead.")

var NUpValuesForBooklets = []int{2, 4, 6, 8}

func DefaultBookletConfig() *model.NUp { _ = "STUB: not implemented"; return nil }

func PDFBookletConfig(val int, desc string, conf *model.Configuration) (*model.NUp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ImageBookletConfig(val int, desc string, conf *model.Configuration) (*model.NUp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getPageNumber(pageNumbers []int, n int) int { _ = "STUB: not implemented"; return 0 }

type pageNumberFunction func(inputPageNr int, pageCount int, pageNumbers []int, nup *model.NUp) (int, bool)

func nup2OutputPageNr(inputPageNr, inputPageCount int, pageNumbers []int, _ *model.NUp) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func get4upPos(pos int, isLandscape bool) (out int) { _ = "STUB: not implemented"; return 0 }

func nup4OutputPageNr(inputPageNr int, pageCount int, pageNumbers []int, nup *model.NUp) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func nup4BasicSideFoldOutputPageNr(positionNumber int, inputPageCount int, pageNumbers []int, nup *model.NUp) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func nup4BasicTopFoldOutputPageNr(positionNumber int, inputPageCount int, pageNumbers []int, nup *model.NUp) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func nup4AdvancedSideFoldOutputPageNr(inputPageNr int, inputPageCount int, pageNumbers []int, nup *model.NUp) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func nupLRTBOutputPageNr(positionNumber int, inputPageCount int, pageNumbers []int, nup *model.NUp) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func nup8OutputPageNr(portraitPositionNumber int, inputPageCount int, pageNumbers []int, nup *model.NUp) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func nupPerfectBound(positionNumber int, inputPageCount int, pageNumbers []int, nup *model.NUp) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func GetBookletOrdering(pages types.IntSet, nup *model.NUp) []model.BookletPage {
	_ = "STUB: not implemented"
	return nil
}

func getBookletPageOrdering(nup *model.NUp, pageNumbers []int, pageCount int) []model.BookletPage {
	_ = "STUB: not implemented"
	return nil
}

func bookletPages(
	ctx *model.Context,
	selectedPages types.IntSet,
	nup *model.NUp,
	pagesDict types.Dict,
	pagesIndRef *types.IndirectRef) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func BookletFromImages(ctx *model.Context, fileNames []string, nup *model.NUp, pagesDict types.Dict, pagesIndRef *types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func BookletFromPDF(ctx *model.Context, selectedPages types.IntSet, nup *model.NUp) error {
	_ = "STUB: not implemented"
	return nil
}
