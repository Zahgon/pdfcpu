package main

import (
	"os"
	"runtime/debug"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

var (
	version = model.VersionStr
	commit  = "?"
	date    = "?"
)

func init() {
	updateVersionInfoFromBuildInfo()
}

func updateVersionInfoFromBuildInfo() { _ = "STUB: not implemented"; return }

func setVersionInfoFromBuildSettings(settings []debug.BuildSetting) {
	_ = "STUB: not implemented"
	return
}

func shortCommit(s string) string { _ = "STUB: not implemented"; return "" }

func printError(err error) { _ = "STUB: not implemented"; return }

func main() {
	if err := Execute(); err != nil {
		printError(err)
		os.Exit(1)
	}
}
