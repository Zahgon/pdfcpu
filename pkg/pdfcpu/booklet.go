package pdfcpu

import (
	"errors"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

var errInvalidBookletAdvanced = errors.New("booklet advanced cannot have binding along the top (portrait short-edge, landscape long-edge). use plain booklet instead")

var nUpValuesForBooklets = []int{2, 4, 6, 8}

func NUpValuesForBooklets() []int { _ = "STUB: not implemented"; return nil }

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

func getBookletOrdering(pages types.IntSet, nup *model.NUp) []model.BookletPage {
	_ = "STUB: not implemented"
	return nil
}

func getBookletPageOrdering(nup *model.NUp, pageNumbers []int, pageCount int) []model.BookletPage {
	_ = "STUB: not implemented"
	return nil
}

func wrapBookletOutputPageError(pageNr int, err error) error { _ = "STUB: not implemented"; return nil }

func wrapBookletImageError(imageNr int, fileName, phase string, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func loadBookletImageResource(
	xRefTable *model.XRefTable,
	imageNr int,
	fileName string) (imgIndRef *types.IndirectRef, w, h int, err error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, nil
}

func loadBookletImageResourceWith(
	xRefTable *model.XRefTable,
	imageNr int,
	fileName string,
	createImageResource func(*model.XRefTable, io.Reader) (*types.IndirectRef, int, int, error),
) (imgIndRef *types.IndirectRef, w, h int, err error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, nil
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

func bookletImageNUp(nup *model.NUp) *model.NUp { _ = "STUB: not implemented"; return nil }

func BookletFromImages(ctx *model.Context, fileNames []string, nup *model.NUp, pagesDict types.Dict, pagesIndRef *types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func BookletFromPDF(ctx *model.Context, selectedPages types.IntSet, nup *model.NUp) error {
	_ = "STUB: not implemented"
	return nil
}
