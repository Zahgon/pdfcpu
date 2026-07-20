package pdfcpu

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func createMHBEDict() *types.Dict { _ = "STUB: not implemented"; return nil }

func createMediaPlayersDict() *types.Dict { _ = "STUB: not implemented"; return nil }

func createMediaOffsetDict() *types.Dict { _ = "STUB: not implemented"; return nil }

func createSectionMHBEDict() *types.Dict { _ = "STUB: not implemented"; return nil }

func createMediaClipDataDict(xRefTable *model.XRefTable) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createMediaPlayParamsMHBE() *types.Dict { _ = "STUB: not implemented"; return nil }

func createMediaPlayParamsDict() *types.Dict { _ = "STUB: not implemented"; return nil }

func createFloatingWindowsParamsDict() *types.Dict { _ = "STUB: not implemented"; return nil }

func createScreenParamsDict() *types.Dict { _ = "STUB: not implemented"; return nil }

func createMediaRendition(mediaClipDataDict *types.IndirectRef) *types.Dict {
	_ = "STUB: not implemented"
	return nil
}

func createSectionMediaRendition(mediaClipDataDict *types.IndirectRef) *types.Dict {
	_ = "STUB: not implemented"
	return nil
}

func createSelectorRendition(mediaClipDataDict *types.IndirectRef) *types.Dict {
	_ = "STUB: not implemented"
	return nil
}
