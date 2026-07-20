package validate

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"github.com/pkg/errors"
)

var errInvalidPageAnnotArray = errors.New("pdfcpu: validatePageAnnotations: page annotation array without indirect references.")

func validateBorderEffectDictEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateBorderStyleDict(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateIconFitDictEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAppearanceCharacteristicsDictEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictText(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateActionOrDestination(xRefTable *model.XRefTable, d types.Dict, dictName string, sinceVersion model.Version) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func validateURIActionDictEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictLink(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAPAndDA(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictFreeTextPart1(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictFreeTextPart2(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictFreeText(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateEntryMeasure(xRefTable *model.XRefTable, d types.Dict, dictName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCP(s string) bool { _ = "STUB: not implemented"; return false }

func validateAnnotationDictLinePart1(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictLinePart2(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictLine(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictCircleOrSquare(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateEntryIT(xRefTable *model.XRefTable, d types.Dict, dictName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictPolyLine(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateTextMarkupAnnotation(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictStamp(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictCaret(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictInk(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictPopup(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictFileAttachment(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictSound(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMovieDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictMovie(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictWidget(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictScreen(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictPrinterMark(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictTrapNet(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictWatermark(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDict3D(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateEntryIC(xRefTable *model.XRefTable, d types.Dict, dictName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictRedact(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateRichMediaAnnotation(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateExDataDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePopupEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateIRTEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMarkupAnnotationPart1(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMarkupAnnotationPart2(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMarkupAnnotation(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateEntryP(xRefTable *model.XRefTable, d types.Dict, dictName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAppearDictEntry(xRefTable *model.XRefTable, d types.Dict, dictName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDashPatternArray(xRefTable *model.XRefTable, arr types.Array) bool {
	_ = "STUB: not implemented"
	return false
}

func validateBorderArray(xRefTable *model.XRefTable, a types.Array) bool {
	_ = "STUB: not implemented"
	return false
}

func validateAnnotationDictGeneralPart1(xRefTable *model.XRefTable, d types.Dict, dictName string) (*types.Name, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateAnnotationDictGeneralPart2(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDictGeneral(xRefTable *model.XRefTable, d types.Dict, dictName string) (*types.Name, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateAnnotationDictConcrete(xRefTable *model.XRefTable, d types.Dict, dictName string, subtype types.Name) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationDict(xRefTable *model.XRefTable, d types.Dict) (isTrapNet bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func addAnnotation(ann model.AnnotationRenderer, pgAnnots model.PgAnnots, i int, hasIndRef bool, indRef types.IndirectRef) {
	_ = "STUB: not implemented"
	return
}

func detectSignature(xRefTable *model.XRefTable, annotDict types.Dict, objNr, incr int) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAnnotationsArray(xRefTable *model.XRefTable, a types.Array) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func validatePageAnnotations(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePagesAnnotations(xRefTable *model.XRefTable, d types.Dict, curPage int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
