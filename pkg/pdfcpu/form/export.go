package form

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/primitives"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

const (
	REQUIRED = true

	OPTIONAL = false
)

type Header struct {
	Source   string   `json:"source"`
	Version  string   `json:"version"`
	Creation string   `json:"creation"`
	ID       []string `json:"id,omitempty"`
	Title    string   `json:"title,omitempty"`
	Author   string   `json:"author,omitempty"`
	Creator  string   `json:"creator,omitempty"`
	Producer string   `json:"producer,omitempty"`
	Subject  string   `json:"subject,omitempty"`
	Keywords string   `json:"keywords,omitempty"`
}

type TextField struct {
	Pages     []int  `json:"pages"`
	ID        string `json:"id"`
	Name      string `json:"name,omitempty"`
	AltName   string `json:"altname,omitempty"`
	Default   string `json:"default,omitempty"`
	Value     string `json:"value"`
	MaxLen    int    `json:"maxlen,omitempty"`
	Multiline bool   `json:"multiline"`
	Locked    bool   `json:"locked"`
}

type DateField struct {
	Pages   []int  `json:"pages"`
	ID      string `json:"id"`
	Name    string `json:"name,omitempty"`
	AltName string `json:"altname,omitempty"`
	Format  string `json:"format"`
	Default string `json:"default,omitempty"`
	Value   string `json:"value"`
	Locked  bool   `json:"locked"`
}

type CheckBox struct {
	Pages   []int  `json:"pages"`
	ID      string `json:"id"`
	Name    string `json:"name,omitempty"`
	AltName string `json:"altname,omitempty"`
	Default bool   `json:"default"`
	Value   bool   `json:"value"`
	Locked  bool   `json:"locked"`
}

type RadioButtonGroup struct {
	Pages   []int    `json:"pages"`
	ID      string   `json:"id"`
	Name    string   `json:"name,omitempty"`
	AltName string   `json:"altname,omitempty"`
	Options []string `json:"options"`
	Default string   `json:"default,omitempty"`
	Value   string   `json:"value"`
	Locked  bool     `json:"locked"`
}

type ComboBox struct {
	Pages    []int    `json:"pages"`
	ID       string   `json:"id"`
	Name     string   `json:"name,omitempty"`
	AltName  string   `json:"altname,omitempty"`
	Editable bool     `json:"editable"`
	Options  []string `json:"options"`
	Default  string   `json:"default,omitempty"`
	Value    string   `json:"value"`
	Locked   bool     `json:"locked"`
}

type ListBox struct {
	Pages    []int    `json:"pages"`
	ID       string   `json:"id"`
	Name     string   `json:"name,omitempty"`
	AltName  string   `json:"altname,omitempty"`
	Multi    bool     `json:"multi"`
	Options  []string `json:"options"`
	Defaults []string `json:"defaults,omitempty"`
	Values   []string `json:"values,omitempty"`
	Locked   bool     `json:"locked"`
}

type Page struct {
	ImageBoxes []*primitives.ImageBox `json:"image,omitempty"`
}

type Form struct {
	TextFields        []*TextField        `json:"textfield,omitempty"`
	DateFields        []*DateField        `json:"datefield,omitempty"`
	CheckBoxes        []*CheckBox         `json:"checkbox,omitempty"`
	RadioButtonGroups []*RadioButtonGroup `json:"radiobuttongroup,omitempty"`
	ComboBoxes        []*ComboBox         `json:"combobox,omitempty"`
	ListBoxes         []*ListBox          `json:"listbox,omitempty"`
	Pages             map[string]*Page    `json:"pages,omitempty"`
	FileName          string              `json:"filename,omitempty"`
}

type FormGroup struct {
	Header Header `json:"header"`
	Forms  []Form `json:"forms"`
}

func (f Form) textFieldValueAndLock(id, name string) (string, bool, bool) {
	_ = "STUB: not implemented"
	return "", false, false
}

func (f Form) dateFieldValueAndLock(id, name string) (string, bool, bool) {
	_ = "STUB: not implemented"
	return "", false, false
}

func (f Form) checkBoxValueAndLock(id, name string) (bool, bool, bool) {
	_ = "STUB: not implemented"
	return false, false, false
}

func (f Form) radioButtonGroupValueAndLock(id, name string) (string, bool, bool) {
	_ = "STUB: not implemented"
	return "", false, false
}

func (f Form) comboBoxValueAndLock(id, name string) (string, bool, bool) {
	_ = "STUB: not implemented"
	return "", false, false
}

func (f Form) listBoxValuesAndLock(id, name string) ([]string, bool, bool) {
	_ = "STUB: not implemented"
	return nil, false, false
}

func locateAPN(xRefTable *model.XRefTable, d types.Dict) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func extractRadioButtonGroupOptions(xRefTable *model.XRefTable, d types.Dict) ([]string, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func resolveOption(s string, opts []string, explicit bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func extractRadioButtonGroup(xRefTable *model.XRefTable, page int, d types.Dict, id, name, altName string, locked bool) (*RadioButtonGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractCheckBox(xRefTable *model.XRefTable, page int, d types.Dict, id, name, altName string, locked bool) (*CheckBox, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractComboBox(xRefTable *model.XRefTable, page int, d types.Dict, id, name, altName string, locked bool) (*ComboBox, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dateFormatFromJSAction(d types.Dict) (*primitives.DateFormat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dateStringEntry(xRefTable *model.XRefTable, d types.Dict, key string) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func extractDateFormat(xRefTable *model.XRefTable, d types.Dict) (*primitives.DateFormat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractDateField(xRefTable *model.XRefTable, page int, d types.Dict, id, name, altName string, df *primitives.DateFormat, locked bool) (*DateField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractTextField(xRefTable *model.XRefTable, page int, d types.Dict, id, name, altName string, ff *int, locked bool) (*TextField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractListBox(xRefTable *model.XRefTable, page int, d types.Dict, id, name, altName string, locked, multi bool) (*ListBox, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func header(xRefTable *model.XRefTable, source string) Header {
	_ = "STUB: not implemented"
	return *new(Header)
}

func fieldsForAnnots(xRefTable *model.XRefTable, annots, fields types.Array) (map[string]fieldInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func exportBtn(
	xRefTable *model.XRefTable,
	i int,
	form *Form,
	d types.Dict,
	id, name, altName string,
	locked bool,
	ok *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func exportCh(
	xRefTable *model.XRefTable,
	i int,
	form *Form,
	d types.Dict,
	id, name, altName string,
	locked bool,
	ok *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func exportTx(
	xRefTable *model.XRefTable,
	i int,
	form *Form,
	d types.Dict,
	id, name, altName string,
	ff *int,
	locked bool,
	ok *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func exportPageField(ft string, xRefTable *model.XRefTable, i int, form *Form, d types.Dict, id, name, altName string, locked bool, ok *bool, ff *int) error {
	_ = "STUB: not implemented"
	return nil
}

func exportPageFields(xRefTable *model.XRefTable, i int, form *Form, m map[string]fieldInfo, ok *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func ExportForm(xRefTable *model.XRefTable, source string) (*FormGroup, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

type exportFormFunc func(*model.XRefTable, string) (*FormGroup, bool, error)

type marshalFormJSONFunc func(any, string, string) ([]byte, error)

func exportFormJSON(xRefTable *model.XRefTable, source string, w io.Writer, export exportFormFunc, marshal marshalFormJSONFunc) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func ExportFormJSON(xRefTable *model.XRefTable, source string, w io.Writer) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
