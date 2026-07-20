package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/form"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pkg/errors"
)

var (
	ErrNoFormData           = errors.New("pdfcpu: missing form data")
	ErrNoFormFieldsAffected = errors.New("pdfcpu: no form fields affected")
	ErrInvalidCSV           = errors.New("pdfcpu: invalid csv input file")
	ErrInvalidJSON          = errors.New("pdfcpu: invalid JSON encoding")
)

func FormFields(rs io.ReadSeeker, conf *model.Configuration) (fields []form.Field, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RemoveFormFields(rs io.ReadSeeker, w io.Writer, fieldIDsOrNames []string, conf *model.Configuration) (err error) {
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

func ExportFormJSON(rs io.ReadSeeker, w io.Writer, source string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ExportFormFile(inFilePDF, outFileJSON string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func validateComboBoxValues(f form.Form) error { _ = "STUB: not implemented"; return nil }

func validateListBoxValues(f form.Form) error { _ = "STUB: not implemented"; return nil }

func validateRadioButtonGroupValues(f form.Form) error { _ = "STUB: not implemented"; return nil }

func validateOptionValues(f form.Form) error { _ = "STUB: not implemented"; return nil }

func fillPostProc(ctx *model.Context, pp []*model.Page) error {
	_ = "STUB: not implemented"
	return nil
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

func mergeForms(outDir, fileName string, outFiles []string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func multiFillFormJSON(inFilePDF string, rd io.Reader, outDir, fileName string, merge bool, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func parseCSVLines(rd io.Reader) ([][]string, error) { _ = "STUB: not implemented"; return nil, nil }

func multiFillFormCSV(inFilePDF string, rd io.Reader, outDir, fileName string, merge bool, conf *model.Configuration) error {
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
