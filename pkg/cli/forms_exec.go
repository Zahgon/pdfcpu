package cli

import (
	"io"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/form"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func listFormFields(rs io.ReadSeeker, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func exportFormGroup(rs io.ReadSeeker, source string, conf *model.Configuration) (*form.FormGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func listFormFieldsJSON(inFiles []string, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func formFieldReadSeeker(fn string) (io.ReadSeeker, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadSeeker), nil
}

func formFieldSource(fn string) string { _ = "STUB: not implemented"; return "" }

func ListFormFieldsFile(inFiles []string, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListFormFields(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func formInOut(cmd *Command) (io.ReadSeeker, io.Writer, func(), error) {
	_ = "STUB: not implemented"
	return *new(io.ReadSeeker), *new(io.Writer), nil, nil
}

func formDataReader(filename string) (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

func formTemplateFileFromStdin() (string, func(), error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func fillFormData(cmd *Command) (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

func formPDFFileCommand(inFile, outFile string, fileFn func() error, readerFn func(io.ReadSeeker, io.Writer) error) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func formPDFWithData(cmd *Command, fileFn func() error, readerFn func(io.ReadSeeker, io.Reader, io.Writer) error) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RemoveFormFields(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func LockFormFields(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func UnlockFormFields(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ResetFormFields(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ExportFormFields(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func FillFormFields(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func multiFillFormInputFile(cmd *Command) (string, func(), error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func multiFillFormOutputFile(cmd *Command) string { _ = "STUB: not implemented"; return "" }

func multiFillFormFieldsToStdout(cmd *Command, inFile string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MultiFillFormFields(cmd *Command) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
