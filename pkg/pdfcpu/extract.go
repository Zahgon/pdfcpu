package pdfcpu

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func sortedObjectNumbers[V any](m map[int]V) []int { _ = "STUB: not implemented"; return nil }

func ImageObjNrs(ctx *model.Context, pageNr int) []int { _ = "STUB: not implemented"; return nil }

func StreamLength(ctx *model.Context, sd *types.StreamDict) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ColorSpaceString(ctx *model.Context, sd *types.StreamDict) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func colorSpaceArrayName(cs types.Array) (types.Name, error) {
	_ = "STUB: not implemented"
	return *new(types.Name), nil
}

func colorSpaceNameComponents(cs types.Name) int { _ = "STUB: not implemented"; return 0 }

func colorSpaceArrayEntry(cs types.Array, index int) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func colorSpaceDereferenceError(err error, o types.Object, topic string) error {
	_ = "STUB: not implemented"
	return nil
}

func dereferenceRequiredStreamDict(xRefTable *model.XRefTable, o types.Object, topic string) (*types.StreamDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func indexedColorSpaceComponents(xRefTable *model.XRefTable, cs types.Array) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func iccBasedColorSpaceComponents(xRefTable *model.XRefTable, cs types.Array) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func deviceNColorSpaceComponents(cs types.Array) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func colorSpaceArrayComponents(xRefTable *model.XRefTable, cs types.Array) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ColorSpaceComponents(xRefTable *model.XRefTable, sd *types.StreamDict) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func imageWidth(ctx *model.Context, sd *types.StreamDict, objNr int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func imageHeight(ctx *model.Context, sd *types.StreamDict, objNr int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func imageStub(
	ctx *model.Context,
	sd *types.StreamDict,
	resourceId, filters, lastFilter string,
	decodeParms types.Dict,
	thumb, imgMask bool,
	objNr int) (*model.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func prepareExtractImage(sd *types.StreamDict) (string, string, types.Dict, bool) {
	_ = "STUB: not implemented"
	return "", "", *new(types.Dict), false
}

func decodeImage(ctx *model.Context, sd *types.StreamDict, filters, lastFilter string, objNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func img(
	ctx *model.Context,
	sd *types.StreamDict,
	thumb bool,
	resourceID, filters, lastFilter string,
	objNr int) (*model.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExtractImage(ctx *model.Context, sd *types.StreamDict, thumb bool, resourceID string, objNr int, stub bool) (*model.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validatePageNumber(ctx *model.Context, pageNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func failOnUnsupportedResource(ctx *model.Context) bool { _ = "STUB: not implemented"; return false }

func skipUnsupportedResource(ctx *model.Context, err error) bool {
	_ = "STUB: not implemented"
	return false
}

func ExtractPageImages(ctx *model.Context, pageNr int, stub bool) (map[int]model.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Font struct {
	io.Reader
	Name  string
	Type  string
	ObjNr int
}

func FontObjNrs(ctx *model.Context, pageNr int) []int { _ = "STUB: not implemented"; return nil }

func ExtractFont(ctx *model.Context, fontObject model.FontObject, objNr int) (*Font, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExtractPageFonts(ctx *model.Context, pageNr int, objNrs, skipped types.IntSet) ([]Font, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExtractFormFonts(ctx *model.Context) ([]Font, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExtractPages(ctx *model.Context, pageNrs []int, usePgCache bool) (*model.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExtractPageContent(ctx *model.Context, pageNr int) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

type Metadata struct {
	io.Reader
	ObjNr       int
	ParentObjNr int
	ParentType  string
}

func extractMetadataFromDict(ctx *model.Context, d types.Dict, parentObjNr int) (*Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExtractMetadata(ctx *model.Context) ([]Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
