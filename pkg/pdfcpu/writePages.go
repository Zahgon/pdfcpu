package pdfcpu

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func writePageEntry(ctx *model.Context, d types.Dict, dictName, entryName string, statsAttr int) error {
	_ = "STUB: not implemented"
	return nil
}

func writePageDict(ctx *model.Context, indRef *types.IndirectRef, pageDict types.Dict, pageNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func pageNodeDict(ctx *model.Context, o types.Object) (types.Dict, *types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil, nil
}

func writeKids(ctx *model.Context, a types.Array, pageNr *int, depth int, visit *model.PageTreeVisit) (types.Array, int, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), 0, nil
}

func writePageEntries(ctx *model.Context, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func writePagesDictDepth(ctx *model.Context, indRef *types.IndirectRef, pageNr *int, depth int, visit *model.PageTreeVisit) (skip bool, writtenPages int, err error) {
	_ = "STUB: not implemented"
	return false, 0, nil
}

func writePagesDict(ctx *model.Context, indRef *types.IndirectRef, pageNr *int) (skip bool, writtenPages int, err error) {
	_ = "STUB: not implemented"
	return false, 0, nil
}
