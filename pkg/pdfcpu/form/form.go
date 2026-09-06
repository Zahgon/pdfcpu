package form

import (
	"errors"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type FieldType int

const (
	FTText FieldType = iota
	FTDate
	FTCheckBox
	FTComboBox
	FTListBox
	FTRadioButtonGroup
)

var (
	ErrMissingJSONWriter = errors.New("missing JSON writer")

	errNoFormFieldsAvailable = errors.New("no form fields available")
	errFormFieldsNotRemoved  = errors.New("some form fields could not be removed")
)

func (ft FieldType) String() string { _ = "STUB: not implemented"; return "" }

type Field struct {
	Pages   []int
	Locked  bool
	Typ     FieldType
	ID      string
	Name    string
	AltName string
	Dv      string
	V       string
	Opts    string
}

func (f Field) pageString() string { _ = "STUB: not implemented"; return "" }

type FieldMeta struct {
	altName, def, val, opt                              bool
	pageMax, defMax, valMax, idMax, nameMax, altNameMax int
}

func Fields(xRefTable *model.XRefTable) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func indirectRef(o types.Object, context string, index int) (types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return *new(types.IndirectRef), nil
}

func dictNameEntry(xRefTable *model.XRefTable, d types.Dict, key string) (types.Name, bool, error) {
	_ = "STUB: not implemented"
	return *new(types.Name), false, nil
}

func fullyQualifiedFieldNameDepth(xRefTable *model.XRefTable, indRef types.IndirectRef, fields types.Array, id, name *string, depth int, visit *model.FormFieldVisit) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func fullyQualifiedFieldName(xRefTable *model.XRefTable, indRef types.IndirectRef, fields types.Array, id, name *string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type fieldInfo struct {
	id     string
	name   string
	ft     *string
	indRef *types.IndirectRef
}

func isField(xRefTable *model.XRefTable, indRef types.IndirectRef, fields types.Array) (bool, *fieldInfo, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func extractStringSlice(a types.Array) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseOptions(xRefTable *model.XRefTable, d types.Dict, required bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseStringLiteralArray(xRefTable *model.XRefTable, d types.Dict, key string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func collectRadioButtonGroupOptions(xRefTable *model.XRefTable, d types.Dict) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func collectRadioButtonGroup(xRefTable *model.XRefTable, d types.Dict, f *Field, fm *FieldMeta) error {
	_ = "STUB: not implemented"
	return nil
}

func collectBtn(xRefTable *model.XRefTable, d types.Dict, f *Field, fm *FieldMeta) error {
	_ = "STUB: not implemented"
	return nil
}

func collectComboBox(d types.Dict, f *Field, fm *FieldMeta) error {
	_ = "STUB: not implemented"
	return nil
}

func collectListBox(xRefTable *model.XRefTable, multi bool, d types.Dict, f *Field, fm *FieldMeta) error {
	_ = "STUB: not implemented"
	return nil
}

func collectCh(xRefTable *model.XRefTable, d types.Dict, f *Field, fm *FieldMeta) error {
	_ = "STUB: not implemented"
	return nil
}

func inheritedV(xRefTable *model.XRefTable, d types.Dict) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getV(xRefTable *model.XRefTable, d types.Dict) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func inheritedDV(xRefTable *model.XRefTable, d types.Dict) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getDV(xRefTable *model.XRefTable, d types.Dict) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func cleanTextForListCmd(s string, maxWidth int) string { _ = "STUB: not implemented"; return "" }

func collectTx(xRefTable *model.XRefTable, d types.Dict, f *Field, fm *FieldMeta, maxWidth int) error {
	_ = "STUB: not implemented"
	return nil
}

func collectField(xRefTable *model.XRefTable, ft string, d types.Dict, f *Field, fm *FieldMeta, maxWidth int) error {
	_ = "STUB: not implemented"
	return nil
}

func locateField(fs *[]Field, fi *fieldInfo, fm *FieldMeta, pageNr int) bool {
	_ = "STUB: not implemented"
	return false
}

func collectPageField(
	xRefTable *model.XRefTable,
	d types.Dict,
	pageNr int,
	fi *fieldInfo,
	fm *FieldMeta,
	fs *[]Field,
	maxWidth int) error {
	_ = "STUB: not implemented"
	return nil
}

func collectPageFields(
	xRefTable *model.XRefTable,
	wAnnots model.Annot,
	fields types.Array,
	p int,
	fm *FieldMeta,
	fs *[]Field,
	maxWidth int) error {
	_ = "STUB: not implemented"
	return nil
}

func collectFields(xRefTable *model.XRefTable, fields types.Array, fm *FieldMeta, maxWidth int) ([]Field, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func calcListHeader(fm *FieldMeta) (string, []int) { _ = "STUB: not implemented"; return "", nil }

func multiPageFieldsMap(fs []Field) map[string][]Field { _ = "STUB: not implemented"; return nil }

func renderMultiPageFields(m map[string][]Field, fm *FieldMeta) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func renderFields(ctx *model.Context, fs []Field, fm *FieldMeta) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FormFields(ctx *model.Context) ([]Field, *FieldMeta, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func ListFormFields(ctx *model.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func annotIndRefsDepth(xRefTable *model.XRefTable, fields types.Array, depth int, visit *model.FormFieldVisit) ([]types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func annotIndRefs(xRefTable *model.XRefTable, fields types.Array) ([]types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func annotIndRefSameLevel(xRefTable *model.XRefTable, fields types.Array, fieldIDOrName string) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func annotIndRefForFieldDepth(xRefTable *model.XRefTable, fields types.Array, fieldIDOrName string, depth int, visit *model.FormFieldVisit) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func annotIndRefForField(xRefTable *model.XRefTable, fields types.Array, fieldIDOrName string) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func annotIndRefsForFields(xRefTable *model.XRefTable, f []string, fields types.Array) ([]types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func removeIndRefByIndex(indRefs []types.IndirectRef, i int) []types.IndirectRef {
	_ = "STUB: not implemented"
	return nil
}

func removeFormFieldsDepth(xRefTable *model.XRefTable, indRefs *[]types.IndirectRef, fields *types.Array, depth int, visit *model.FormFieldVisit) error {
	_ = "STUB: not implemented"
	return nil
}

func removeFormFields(xRefTable *model.XRefTable, indRefs *[]types.IndirectRef, fields *types.Array) error {
	_ = "STUB: not implemented"
	return nil
}

func deletePageAnnots(xRefTable *model.XRefTable, m map[types.IndirectRef]bool, ok *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func RemoveFormFields(ctx *model.Context, fieldIDsOrNames []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func resetBtn(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func resetComboBoxOrRegularListBox(d types.Dict, opts []string, ff *int) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func resetMultiListBox(xRefTable *model.XRefTable, d types.Dict, opts []string) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func resetCh(ctx *model.Context, d types.Dict, fonts map[string]types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func resetTx(ctx *model.Context, d types.Dict, fonts map[string]types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func matchField(fi *fieldInfo, fieldIDsOrNames []string) bool {
	_ = "STUB: not implemented"
	return false
}

func resetPageFields(
	ctx *model.Context,
	fieldIDsOrNames []string,
	wAnnots model.Annot,
	fields types.Array,
	fonts map[string]types.IndirectRef,
	ok *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func ResetFormFields(ctx *model.Context, fieldIDsOrNames []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func lockFormField(d types.Dict) { _ = "STUB: not implemented"; return }

func ensureAP(ctx *model.Context, d types.Dict, fi *fieldInfo, fonts map[string]types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func lockPageFields(
	ctx *model.Context,
	fieldIDsOrNames []string,
	fields types.Array,
	wAnnots model.Annot,
	fonts map[string]types.IndirectRef,
	ok *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func LockFormFields(ctx *model.Context, fieldIDsOrNames []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func unlockFormField(d types.Dict) { _ = "STUB: not implemented"; return }

func deleteAP(d types.Dict, fi *fieldInfo) error { _ = "STUB: not implemented"; return nil }

func unlockPageFields(
	xRefTable *model.XRefTable,
	fieldIDsOrNames []string,
	fields types.Array,
	wAnnots model.Annot,
	ok *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func UnlockFormFields(ctx *model.Context, fieldIDsOrNames []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
