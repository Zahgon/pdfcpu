package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/form"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func FormFields(rs io.ReadSeeker, conf *model.Configuration) (fields []form.Field, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListFormFields(rs io.ReadSeeker, conf *model.Configuration) (fields []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RemoveFormFields(rs io.ReadSeeker, w io.Writer, fieldIDsOrNames []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type formFieldMutation func(io.ReadSeeker, io.Writer, []string, *model.Configuration) error

func mutateFormFieldsFile(inFile, outFile string, fieldIDsOrNames []string, conf *model.Configuration, operation string, mutate formFieldMutation) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemoveFormFieldsFile(inFile, outFile string, fieldIDsOrNames []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func LockFormFields(rs io.ReadSeeker, w io.Writer, fieldIDsOrNames []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func LockFormFieldsFile(inFile, outFile string, fieldIDsOrNames []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func UnlockFormFields(rs io.ReadSeeker, w io.Writer, fieldIDsOrNames []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func UnlockFormFieldsFile(inFile, outFile string, fieldIDsOrNames []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ResetFormFields(rs io.ReadSeeker, w io.Writer, fieldIDsOrNames []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ResetFormFieldsFile(inFile, outFile string, fieldIDsOrNames []string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ExportForm(rs io.ReadSeeker, source string, conf *model.Configuration) (formGroup *form.FormGroup, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func exportedFormGroup(ctx *model.Context, source string) (*form.FormGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type formJSONExporter func(*model.XRefTable, string, io.Writer) (bool, error)

func exportFormJSONResult(xRefTable *model.XRefTable, source string, w io.Writer, export formJSONExporter) error {
	_ = "STUB: not implemented"
	return nil
}

func ExportFormJSON(rs io.ReadSeeker, w io.Writer, source string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ExportFormFile(inFilePDF, outFileJSON string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func validOptionValue(value string, options []string) bool { _ = "STUB: not implemented"; return false }

func unknownOptionValueError(fieldType string, fieldIndex int, name, value string, options []string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFormData(f form.Form) error { _ = "STUB: not implemented"; return nil }

func validateComboBoxValues(f form.Form) error { _ = "STUB: not implemented"; return nil }

func validateListBoxValues(f form.Form) error { _ = "STUB: not implemented"; return nil }

func validateRadioButtonGroupValues(f form.Form) error { _ = "STUB: not implemented"; return nil }

func validateOptionValues(f form.Form) error { _ = "STUB: not implemented"; return nil }

func fillPostProc(ctx *model.Context, pp []*model.Page) error {
	_ = "STUB: not implemented"
	return nil
}

func formGroupFromReader(rd io.Reader) (*form.FormGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validatedFillForm(formGroup *form.FormGroup) (form.Form, error) {
	_ = "STUB: not implemented"
	return *new(form.Form), nil
}

func FillForm(rs io.ReadSeeker, rd io.Reader, w io.Writer, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func FillFormFile(inFilePDF, inFileJSON, outFilePDF string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func parseFormGroup(rd io.Reader) (*form.FormGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func rollbackMultiFillOutputs(outFiles []string) error { _ = "STUB: not implemented"; return nil }

func mergeForms(outDir, fileName string, outFiles []string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func multiFillJSONOutputFile(outDir, fileName, requested string, formNr int) string {
	_ = "STUB: not implemented"
	return ""
}

func multiFillCSVOutputFile(outDir, fileName, requested string, recordNr int) string {
	_ = "STUB: not implemented"
	return ""
}

func multiFillPostProcess(ctx *model.Context, pp []*model.Page, context string, validate bool) error {
	_ = "STUB: not implemented"
	return nil
}

func writeMultiFillOutput(ctx *model.Context, outFile, context string) error {
	_ = "STUB: not implemented"
	return nil
}

func writeMultiFillOutputWith(ctx *model.Context, outFile, context string, writeContext func(*model.Context, io.Writer) error) error {
	_ = "STUB: not implemented"
	return nil
}

func multiFillJSONForm(inFilePDF string, f form.Form, outDir, fileName string, formNr int, conf *model.Configuration, writeContext func(*model.Context, io.Writer) error) (outFile string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func multiFillFormJSON(inFilePDF string, rd io.Reader, outDir, fileName string, merge bool, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func multiFillFormJSONWith(inFilePDF string, rd io.Reader, outDir, fileName string, merge bool, conf *model.Configuration, writeContext func(*model.Context, io.Writer) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func parseCSVLines(rd io.Reader) ([][]string, error) { _ = "STUB: not implemented"; return nil, nil }

func multiFillCSVRecord(inFilePDF string, fieldNames, formRecord []string, outDir, fileName string, recordNr, rowNr int, conf *model.Configuration, writeContext func(*model.Context, io.Writer) error) (outFile string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func multiFillFormCSV(inFilePDF string, rd io.Reader, outDir, fileName string, merge bool, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func multiFillFormCSVWith(inFilePDF string, rd io.Reader, outDir, fileName string, merge bool, conf *model.Configuration, writeContext func(*model.Context, io.Writer) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func MultiFillForm(inFilePDF string, rd io.Reader, outDir, fileName string, format form.DataFormat, merge bool, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func MultiFillFormFile(inFilePDF, inFileData, outDir, outFilePDF string, merge bool, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
