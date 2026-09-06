package pdfcpu

import (
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func logWriteStats(ctx *model.Context) { _ = "STUB: not implemented"; return }

func statsHeadLine() *string { _ = "STUB: not implemented"; return nil }

func statsLine(ctx *model.Context) *string { _ = "STUB: not implemented"; return nil }

type statsAppendFile interface {
	WriteString(string) (int, error)
	Close() error
}

type openStatsAppendFile func(string, int, os.FileMode) (statsAppendFile, error)

func appendStatsFile(ctx *model.Context, openFile openStatsAppendFile) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AppendStatsFile(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }
