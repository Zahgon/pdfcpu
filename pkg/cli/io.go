package cli

import (
	"io"
	"os"
)

type streamInOutFinalizer struct {
	input       *os.File
	temporaryIn *temporaryInput
	output      *os.File
	outFile     string
	replaceOut  string
}

type temporaryInput struct {
	file   *os.File
	path   string
	remove func(string) error
}

var (
	createTemporaryInputFile = os.CreateTemp
	rewindTemporaryInputFile = func(f *os.File) error {
		_, err := f.Seek(0, io.SeekStart)
		return err
	}
)

func (in *temporaryInput) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (in *temporaryInput) Seek(offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (in *temporaryInput) finalize(op string, opErr error) error {
	_ = "STUB: not implemented"
	return nil
}

func readSeekerFromStdin(op string) (*temporaryInput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func withStdinReadSeeker[T any](op string, fn func(io.ReadSeeker) (T, error)) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func closeStreamFile(f *os.File, context string) error { _ = "STUB: not implemented"; return nil }

func removeStreamOutput(fileName, context string) error { _ = "STUB: not implemented"; return nil }

func (f *streamInOutFinalizer) finalize(op string, opErr error) error {
	_ = "STUB: not implemented"
	return nil
}

func createStreamOutput(fileName string) (*os.File, string, string, error) {
	_ = "STUB: not implemented"
	return nil, "", "", nil
}

func streamInOutForOperation(inFile, outFile, op string) (io.ReadSeeker, io.Writer, func(error) error, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadSeeker), *new(io.Writer), nil, nil
}
