package pdfcpu

import (
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

var (
	testDir          = "../../testdata"
	testAudioFileWAV = filepath.Join(testDir, "resources", "test.wav")
)

func CreateXRefTableWithRootDict() (*model.XRefTable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateDemoXRef() (*model.XRefTable, error) { _ = "STUB: not implemented"; return nil, nil }

func addPageTreeForResourceDictInheritanceDemo(xRefTable *model.XRefTable, rootDict types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func CreateResourceDictInheritanceDemoXRef() (*model.XRefTable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createFunctionalShadingDict() types.Dict { _ = "STUB: not implemented"; return *new(types.Dict) }

func createRadialShadingDict() types.Dict { _ = "STUB: not implemented"; return *new(types.Dict) }

func createStreamObjForHalftoneDictType6(xRefTable *model.XRefTable) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createStreamObjForHalftoneDictType10(xRefTable *model.XRefTable) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createStreamObjForHalftoneDictType16(xRefTable *model.XRefTable) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createPostScriptCalculatorFunctionStreamDict(xRefTable *model.XRefTable) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addResources(xRefTable *model.XRefTable, pageDict types.Dict, fontName string) error {
	_ = "STUB: not implemented"
	return nil
}

func CreateTestPageContent(p model.Page) { _ = "STUB: not implemented"; return }

func addContents(xRefTable *model.XRefTable, pageDict types.Dict, p model.Page) error {
	_ = "STUB: not implemented"
	return nil
}

func createBoxColorDict() types.Dict { _ = "STUB: not implemented"; return *new(types.Dict) }

func addViewportDict(pageDict types.Dict) { _ = "STUB: not implemented"; return }

func annotRect(i int, w, h, d, l float64) *types.Rectangle { _ = "STUB: not implemented"; return nil }

func createAnnotsArray(xRefTable *model.XRefTable, pageIndRef types.IndirectRef, mediaBox types.Array) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func createPageWithAnnotations(xRefTable *model.XRefTable, parentPageIndRef types.IndirectRef, mediaBox *types.Rectangle, fontName string) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createPageWithForm(xRefTable *model.XRefTable, parentPageIndRef types.IndirectRef, annotsArray types.Array, mediaBox *types.Rectangle, fontName string) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addPageTreeWithoutPage(xRefTable *model.XRefTable, rootDict types.Dict, d *types.Dim) error {
	_ = "STUB: not implemented"
	return nil
}

func AddPageTreeWithSamplePage(xRefTable *model.XRefTable, rootDict types.Dict, p model.Page) error {
	_ = "STUB: not implemented"
	return nil
}

func addPageTreeWithAnnotations(xRefTable *model.XRefTable, rootDict types.Dict, fontName string) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addPageTreeWithFormFields(xRefTable *model.XRefTable, rootDict types.Dict, annotsArray types.Array, fontName string) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createThreadDict(xRefTable *model.XRefTable, pageIndRef types.IndirectRef) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addThreads(xRefTable *model.XRefTable, rootDict types.Dict, pageIndRef types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func addOpenAction(rootDict types.Dict) error { _ = "STUB: not implemented"; return nil }

func addURI(rootDict types.Dict) { _ = "STUB: not implemented"; return }

func addSpiderInfo(xRefTable *model.XRefTable, rootDict types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func addOCProperties(rootDict types.Dict) error { _ = "STUB: not implemented"; return nil }

func addRequirements(rootDict types.Dict) { _ = "STUB: not implemented"; return }

func CreateAnnotationDemoXRef() (*model.XRefTable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setBit(i uint32, pos uint) uint32 { _ = "STUB: not implemented"; return 0 }

func createNormalAppearanceForFormField(xRefTable *model.XRefTable, w, h float64) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createRolloverAppearanceForFormField(xRefTable *model.XRefTable, w, h float64) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createDownAppearanceForFormField(xRefTable *model.XRefTable, w, h float64) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createFormTextField(xRefTable *model.XRefTable, pageAnnots *types.Array, fontName string) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createYesAppearance(xRefTable *model.XRefTable, resourceDict types.Dict, w, h float64) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createOffAppearance(xRefTable *model.XRefTable, resourceDict types.Dict, w, h float64) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createCheckBoxButtonField(xRefTable *model.XRefTable, pageAnnots *types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createRadioButtonField(xRefTable *model.XRefTable, pageAnnots *types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createResetButton(xRefTable *model.XRefTable, pageAnnots *types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createSubmitButton(xRefTable *model.XRefTable, pageAnnots *types.Array) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func streamObjForXFAElement(xRefTable *model.XRefTable, s string) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createXFAArray(xRefTable *model.XRefTable) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func createFormDict(xRefTable *model.XRefTable, fontName string) (types.Dict, types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), *new(types.Array), nil
}

func CreateFormDemoXRef() (*model.XRefTable, error) { _ = "STUB: not implemented"; return nil, nil }

func CreateContext(xRefTable *model.XRefTable, conf *model.Configuration) *model.Context {
	_ = "STUB: not implemented"
	return nil
}

func CreateContextWithXRefTable(conf *model.Configuration, pageDim *types.Dim) (*model.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createDemoContentStreamDict(xRefTable *model.XRefTable, b []byte) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createDemoPage(xRefTable *model.XRefTable, parentPageIndRef types.IndirectRef, p model.Page) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
