package pdfcpu

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func createTextAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createLinkAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createFreeTextAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createLineAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createSquareAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createCircleAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createPolygonAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createPolyLineAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createHighlightAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createUnderlineAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createSquigglyAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createStrikeOutAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createCaretAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createStampAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createInkAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createPopupAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createFileAttachmentAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createFileSpecDict(xRefTable *model.XRefTable, fileName string) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func createSoundObject(xRefTable *model.XRefTable) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createSoundAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createMovieDict(xRefTable *model.XRefTable) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createMovieAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createMediaRenditionAction(mediaClipDataDict *types.IndirectRef) types.Dict {
	_ = "STUB: not implemented"
	return *new(types.Dict)
}

func createSelectorRenditionAction(mediaClipDataDict *types.IndirectRef) types.Dict {
	_ = "STUB: not implemented"
	return *new(types.Dict)
}

func createScreenAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createWidgetAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createXObjectForPrinterMark(xRefTable *model.XRefTable) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createPrinterMarkAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createXObjectForWaterMark(xRefTable *model.XRefTable) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createWaterMarkAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func create3DAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createRedactAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createRemoteGoToAction(xRefTable *model.XRefTable) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createLinkAnnotationWithRemoteGoToAction(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createEmbeddedGoToAction(xRefTable *model.XRefTable) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createLinkAnnotationWithEmbeddedGoToAction(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createLinkAnnotationDictWithLaunchAction(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createLinkAnnotationDictWithThreadAction(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createLinkAnnotationDictWithSoundAction(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createLinkAnnotationDictWithMovieAction(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createLinkAnnotationDictWithHideAction(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createTrapNetAnnotation(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, annotRect types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
