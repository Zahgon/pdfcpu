package main

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/spf13/cobra"
)

type formMultifillOptions struct {
	mode string
}

type formListOptions struct {
	json bool
}

func formCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func listFormFiles(conf *model.Configuration, args []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func handleListFormFieldsCommand(conf *model.Configuration, args []string, opts *formListOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func formFieldArgs(conf *model.Configuration, args []string, rejectPDFAsOnlyField bool) (string, string, []string, error) {
	_ = "STUB: not implemented"
	return "", "", nil, nil
}

func handleRemoveFormFieldsCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleLockFormCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleUnlockFormCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleResetFormCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleExportFormCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func handleFillFormCommand(conf *model.Configuration, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func multifillMode(opts *formMultifillOptions) error { _ = "STUB: not implemented"; return nil }

func multifillArgs(conf *model.Configuration, args []string, opts *formMultifillOptions) (string, string, string, string, error) {
	_ = "STUB: not implemented"
	return "", "", "", "", nil
}

func handleMultiFillFormCommand(conf *model.Configuration, args []string, opts *formMultifillOptions) error {
	_ = "STUB: not implemented"
	return nil
}
