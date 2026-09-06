package api

import (
	"os"
)

type fileOperations struct {
	openExclusiveFn func(string, int, os.FileMode) (*os.File, error)
	createTempFn    func(string, string) (*os.File, error)
	statFn          func(string) (os.FileInfo, error)
	chmodFn         func(*os.File, os.FileMode) error
	closeFn         func(*os.File) error
	removeFn        func(string) error
	replaceFn       func(string, string) error
}

func (ops fileOperations) closeFile(f *os.File, context string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ops fileOperations) removeFile(fileName, context string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ops fileOperations) replaceFile(oldName, newName, context string) error {
	_ = "STUB: not implemented"
	return nil
}

func defaultFileOperations() fileOperations { _ = "STUB: not implemented"; return *new(fileOperations) }

func closeFile(f *os.File, context string) error { _ = "STUB: not implemented"; return nil }

func removeFile(fileName, context string) error { _ = "STUB: not implemented"; return nil }

func replaceFile(oldName, newName, context string) error { _ = "STUB: not implemented"; return nil }

type stagedFile struct {
	file    *os.File
	context string
}

type stagedOutput struct {
	output         stagedFile
	inputs         []stagedFile
	extraClosers   []func() error
	temporaryFile  string
	destination    string
	removeContext  string
	replaceContext string
	operations     fileOperations
}

func newStagedOutput(
	input, output *os.File,
	temporaryFile, inFile, outFile, replaceOut, operation string,
) stagedOutput {
	_ = "STUB: not implemented"
	return *new(stagedOutput)
}

func openStagedOutput(input *os.File, inFile, outFile, operation string) (stagedOutput, error) {
	_ = "STUB: not implemented"
	return *new(stagedOutput), nil
}

func openStagedOutputWithOperations(
	input *os.File,
	inFile, outFile, operation string,
	ops fileOperations,
) (stagedOutput, error) {
	_ = "STUB: not implemented"
	return *new(stagedOutput), nil
}

func (s stagedOutput) cleanup(processErr error) error { _ = "STUB: not implemented"; return nil }

func (s stagedOutput) withInput(file *os.File, context string) stagedOutput {
	_ = "STUB: not implemented"
	return *new(stagedOutput)
}

func (s stagedOutput) withCloser(closeFn func() error) stagedOutput {
	_ = "STUB: not implemented"
	return *new(stagedOutput)
}

func (s stagedOutput) closeInputs() error { _ = "STUB: not implemented"; return nil }

func (s stagedOutput) commit() error { _ = "STUB: not implemented"; return nil }

func outputAliasesInput(inFile, outFile string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func outputAliasesInputWith(
	inFile, outFile string,
	abs func(string) (string, error),
	stat func(string) (os.FileInfo, error),
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
