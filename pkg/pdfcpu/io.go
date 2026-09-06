package pdfcpu

import (
	"io"
	"os"
)

type namedWriteCloser interface {
	io.WriteCloser
	Name() string
}

func removeStagedFile(path, tmpPath string, remove func(string) error) error {
	_ = "STUB: not implemented"
	return nil
}

func finishStagedFile(
	path string,
	w namedWriteCloser,
	writeErr error,
	closeInput func() error,
	replace func(string, string) error,
	remove func(string) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func writeReader(
	path string,
	r io.Reader,
	createTemp func(string) (namedWriteCloser, error),
	replace func(string, string) error,
	remove func(string) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func createStagedFile(path string) (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

func createWriteReaderTemp(path string) (namedWriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(namedWriteCloser), nil
}

func openStagedFile(dir, prefix string) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isNilReader(r io.Reader) bool { _ = "STUB: not implemented"; return false }

func writeNewFile(rd io.Reader, filePath string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func WriteReader(path string, r io.Reader) error { _ = "STUB: not implemented"; return nil }

func Write(rd io.Reader, filePath string, overwrite bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func CopyFile(srcFilename, destFilename string, overwrite bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
