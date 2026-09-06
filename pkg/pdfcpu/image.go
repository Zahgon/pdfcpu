package pdfcpu

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func Images(ctx *model.Context, selectedPages types.IntSet) ([]map[int]model.Image, *ImageListMaxLengths, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func prepHorSep(horSep *[]int, maxLen *ImageListMaxLengths) string {
	_ = "STUB: not implemented"
	return ""
}

func sortedObjNrs(ii map[int]model.Image) []int { _ = "STUB: not implemented"; return nil }

func attrs(img model.Image) (string, string, string, string, string) {
	_ = "STUB: not implemented"
	return "", "", "", "", ""
}

func listImages(mm []map[int]model.Image, maxLen *ImageListMaxLengths) ([]string, int, int64) {
	_ = "STUB: not implemented"
	return nil, 0, 0
}

type ImageListMaxLengths struct {
	PageNr, ObjNr, ID, Size, Filters int
}

func ListImages(ctx *model.Context, selectedPages types.IntSet) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateImageDimensions(ctx *model.Context, objNr, w, h int) error {
	_ = "STUB: not implemented"
	return nil
}

func UpdateImagesByObjNr(ctx *model.Context, rd io.Reader, objNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func imageResourceRef(d types.Dict, id, context string) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func inheritedImageResourceRef(ctx *model.Context, resources types.Dict, pageNr int, id string) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unknownImageResource(pageNr int, id string) error { _ = "STUB: not implemented"; return nil }

func requiredInheritedImageResourceRef(
	ctx *model.Context,
	resources types.Dict,
	pageNr int,
	id string,
) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func pageImageResource(
	ctx *model.Context,
	pageDict types.Dict,
	inheritedResources types.Dict,
	pageNr int,
	id string,
) (types.Dict, *types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil, nil
}

func UpdateImagesByPageNrAndId(ctx *model.Context, rd io.Reader, pageNr int, id string) error {
	_ = "STUB: not implemented"
	return nil
}
