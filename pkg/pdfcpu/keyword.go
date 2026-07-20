package pdfcpu

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func KeywordsList(ctx *model.Context) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func removeKeywordsFromMetadata(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func finalizeKeywords(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func KeywordsAdd(ctx *model.Context, keywords []string) error {
	_ = "STUB: not implemented"
	return nil
}

func KeywordsRemove(ctx *model.Context, keywords []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
