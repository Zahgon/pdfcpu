package form

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/primitives"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type DataFormat int

const (
	CSV DataFormat = iota
	JSON
)

func cacheResIDs(ctx *model.Context, pdf *primitives.PDF) error {
	_ = "STUB: not implemented"
	return nil
}

func cacheResourceDict(resources types.Dict, name string, pageNr int, cache map[int]types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func addImages(ctx *model.Context, pages map[string]*Page) ([]*model.Page, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type CSVFieldAttributes struct {
	Values []string
	Lock   bool
}

func parsePageNr(s string, ib *primitives.ImageBox) error { _ = "STUB: not implemented"; return nil }

func parseWidth(s string, ib *primitives.ImageBox) error { _ = "STUB: not implemented"; return nil }

func parseHeight(s string, ib *primitives.ImageBox) error { _ = "STUB: not implemented"; return nil }

func parsePositionAnchor(s string, ib *primitives.ImageBox) error {
	_ = "STUB: not implemented"
	return nil
}

func parsePositionOffset(s string, ib *primitives.ImageBox) error {
	_ = "STUB: not implemented"
	return nil
}

func parseImgBackgroundColor(s string, ib *primitives.ImageBox) error {
	_ = "STUB: not implemented"
	return nil
}

func parseImgBorder(s string, ib *primitives.ImageBox) error { _ = "STUB: not implemented"; return nil }

type imageBoxParamMap map[string]func(string, *primitives.ImageBox) error

var imgParamMap = imageBoxParamMap{
	"bgcolor":  parseImgBackgroundColor,
	"border":   parseImgBorder,
	"offset":   parsePositionOffset,
	"page":     parsePageNr,
	"position": parsePositionAnchor,
	"width":    parseWidth,
	"height":   parseHeight,
}

func (m imageBoxParamMap) processImageBoxArg(paramPrefix, paramValueStr string, ib *primitives.ImageBox) error {
	_ = "STUB: not implemented"
	return nil
}

func imageBox(s, src, url string) (*primitives.ImageBox, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func addImageBox(vv []string, fieldName string, im map[string]*Page) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCSVRecord(fieldNames, formRecord []string) error {
	_ = "STUB: not implemented"
	return nil
}

func FieldMap(fieldNames, formRecord []string) (map[string]CSVFieldAttributes, map[string]*Page, string, error) {
	_ = "STUB: not implemented"
	return nil, nil, "", nil
}

func FillDetails(form *Form, fieldMap map[string]CSVFieldAttributes) func(id, name string, fieldType FieldType, format DataFormat) ([]string, bool, bool) {
	_ = "STUB: not implemented"
	return nil
}

func fillRadioButtons(ctx *model.Context, d types.Dict, vNew string, v types.Name) error {
	_ = "STUB: not implemented"
	return nil
}

func fillRadioButtonGroup(
	ctx *model.Context,
	d types.Dict,
	id, name string,
	opts []string,
	locked bool,
	format DataFormat,
	fillDetails func(id, name string, fieldType FieldType, format DataFormat) ([]string, bool, bool),
	ok *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func checkBoxASNames(ctx *model.Context, d types.Dict, child bool) (types.Name, types.Name, error) {
	_ = "STUB: not implemented"
	return *new(types.Name), *new(types.Name), nil
}

func fillCheckBox(
	ctx *model.Context,
	d types.Dict,
	id, name string,
	locked bool,
	format DataFormat,
	fillDetails func(id, name string, fieldType FieldType, format DataFormat) ([]string, bool, bool),
	ok *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func fillBtn(
	ctx *model.Context,
	d types.Dict,
	id, name string,
	locked bool,
	format DataFormat,
	fillDetails func(id, name string, fieldType FieldType, format DataFormat) ([]string, bool, bool),
	ok *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func fillComboBox(
	ctx *model.Context,
	d types.Dict,
	id, name string,
	opts []string,
	locked bool,
	format DataFormat,
	fonts map[string]types.IndirectRef,
	fillDetails func(id, name string, fieldType FieldType, format DataFormat) ([]string, bool, bool),
	ok *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func updateListBoxValues(multi bool, d types.Dict, opts, vNew []string) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func fillListBox(
	ctx *model.Context,
	d types.Dict,
	id, name string,
	opts []string,
	locked bool,
	format DataFormat,
	fonts map[string]types.IndirectRef,
	fillDetails func(id, name string, fieldType FieldType, format DataFormat) ([]string, bool, bool),
	ff *int,
	ok *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func fillCh(
	ctx *model.Context,
	d types.Dict,
	id, name string,
	locked bool,
	format DataFormat,
	fonts map[string]types.IndirectRef,
	fillDetails func(id, name string, fieldType FieldType, format DataFormat) ([]string, bool, bool),
	ff *int,
	ok *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func fillDateField(
	ctx *model.Context,
	d types.Dict,
	id, name, vOld string,
	locked bool,
	format DataFormat,
	fonts map[string]types.IndirectRef,
	fillDetails func(id, name string, fieldType FieldType, format DataFormat) ([]string, bool, bool),
	ok *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func fillTextField(
	ctx *model.Context,
	d types.Dict,
	id, name, vOld string,
	locked bool,
	format DataFormat,
	fonts map[string]types.IndirectRef,
	fillDetails func(id, name string, fieldType FieldType, format DataFormat) ([]string, bool, bool),
	ff *int,
	ok *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func fillTx(
	ctx *model.Context,
	d types.Dict,
	id, name string,
	locked bool,
	format DataFormat,
	fonts map[string]types.IndirectRef,
	fillDetails func(id, name string, fieldType FieldType, format DataFormat) ([]string, bool, bool),
	ff *int,
	ok *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func fillWidgetAnnots(
	ctx *model.Context,
	fields types.Array,
	indRefs map[types.IndirectRef]bool,
	wAnnots model.Annot,
	format DataFormat,
	fonts map[string]types.IndirectRef,
	fillDetails func(id, name string, fieldType FieldType, format DataFormat) ([]string, bool, bool),
	ok *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func setupFillFonts(xRefTable *model.XRefTable) error { _ = "STUB: not implemented"; return nil }

func FillForm(
	ctx *model.Context,
	fillDetails func(id, name string, fieldType FieldType, format DataFormat) ([]string, bool, bool),
	imgs map[string]*Page,
	format DataFormat) (bool, []*model.Page, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}
