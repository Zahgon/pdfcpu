package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

type PageSpan struct {
	From   int
	Thru   int
	Reader io.Reader
}

func pageSpan(ctx *model.Context, from, thru int) (*PageSpan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func spanFileName(fileName string, from, thru int) string { _ = "STUB: not implemented"; return "" }

func splitOutPath(outDir, fileName string, forBookmark bool, from, thru int) string {
	_ = "STUB: not implemented"
	return ""
}

func writePageSpan(ctx *model.Context, from, thru int, outPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func context(rs io.ReadSeeker, conf *model.Configuration) (*model.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func pageSpansSplitAlongBookmarks(ctx *model.Context) ([]*PageSpan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func pageSpans(ctx *model.Context, span int) ([]*PageSpan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func writePageSpans(ctx *model.Context, span int, outDir, fileName string) error {
	_ = "STUB: not implemented"
	return nil
}

func writePageSpansSplitAlongBookmarks(ctx *model.Context, outDir string) error {
	_ = "STUB: not implemented"
	return nil
}

func writePageSpansSplitAlongPages(ctx *model.Context, pageNrs []int, outDir, fileName string) error {
	_ = "STUB: not implemented"
	return nil
}

func SplitRaw(rs io.ReadSeeker, span int, conf *model.Configuration) (ps []*PageSpan, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Split(rs io.ReadSeeker, outDir, fileName string, span int, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func SplitFile(inFile, outDir string, span int, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func SplitByPageNr(rs io.ReadSeeker, outDir, fileName string, pageNrs []int, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func SplitByPageNrFile(inFile, outDir string, pageNrs []int, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}
