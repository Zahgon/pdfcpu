package model

import (
	"bytes"
	"encoding/binary"
	"image"
	"image/color"
	_ "image/png"

	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	_ "golang.org/x/image/webp"
)

type Image struct {
	io.Reader
	Name        string
	FileType    string
	PageNr      int
	ObjNr       int
	Width       int
	Height      int
	Bpc         int
	Cs          string
	Comp        int
	IsImgMask   bool
	HasImgMask  bool
	HasSMask    bool
	Thumb       bool
	Interpol    bool
	Size        int64
	Filter      string
	DecodeParms string
}

func validateImageResourceLimits(xRefTable *XRefTable, c image.Config) error {
	_ = "STUB: not implemented"
	return nil
}

func ImageFileName(fileName string) bool { _ = "STUB: not implemented"; return false }

func ImageFileNames(dir string, maxFileSize types.ByteSize) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createSMaskObject(xRefTable *XRefTable, buf []byte, w, h, bpc int) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateFlateImageStreamDict(xRefTable *XRefTable, buf, sm []byte, w, h, bpc int, cs string) (*types.StreamDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateDCTImageStreamDict(xRefTable *XRefTable, buf []byte, w, h, bpc int, cs string) (*types.StreamDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func writeRGBAImageBuf(img image.Image) ([]byte, []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

func writeRGBA64ImageBuf(img image.Image) []byte { _ = "STUB: not implemented"; return nil }

func writeNRGBAImageBuf(xRefTable *XRefTable, img image.Image) ([]byte, []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

func writeNRGBA64ImageBuf(xRefTable *XRefTable, img image.Image) ([]byte, []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

func writeSoftmask16(xRefTable *XRefTable, img *image.Alpha16) []byte {
	_ = "STUB: not implemented"
	return nil
}

func writeSoftmask(xRefTable *XRefTable, img *image.Alpha) []byte {
	_ = "STUB: not implemented"
	return nil
}

func writeGrayImageBuf(img image.Image) []byte { _ = "STUB: not implemented"; return nil }

func writeGray16ImageBuf(img image.Image) []byte { _ = "STUB: not implemented"; return nil }

func writeCMYKImageBuf(img image.Image) []byte { _ = "STUB: not implemented"; return nil }

func convertToRGBA(img image.Image) *image.RGBA { _ = "STUB: not implemented"; return nil }

func convertNYCbCrAToRGBA(img *image.NYCbCrA) *image.RGBA { _ = "STUB: not implemented"; return nil }

func extractAlpha(img image.Image) image.Image { _ = "STUB: not implemented"; return *new(image.Image) }

func checkIfGray(img image.Image) bool { _ = "STUB: not implemented"; return false }

func convertToGray(img image.Image) image.Image {
	_ = "STUB: not implemented"
	return *new(image.Image)
}

func convertToSepia(img image.Image) *image.RGBA { _ = "STUB: not implemented"; return nil }

func createImageStreamDict(xRefTable *XRefTable, buf, softMask []byte, w, h, bpc int, format, cs string) (*types.StreamDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeJPEG(img image.Image) ([]byte, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func handleRGBImage(xRefTable *XRefTable, img image.Image) ([]byte, []byte, int, string, error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, "", nil
}

func handleGrayImage(xRefTable *XRefTable, img image.Image, imgA image.Image) ([]byte, []byte, int, string, error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, "", nil
}

func handleCMYKImage(img *image.CMYK) ([]byte, []byte, int, string, error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, "", nil
}

func createImageBuf(xRefTable *XRefTable, img image.Image, imgA image.Image, format string) ([]byte, []byte, int, string, error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, "", nil
}

func colorSpaceForJPEGColorModel(cm color.Model) string { _ = "STUB: not implemented"; return "" }

func createDCTImageStreamDictForJPEG(xRefTable *XRefTable, c image.Config, bb bytes.Buffer) (*types.StreamDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createImageResourcesForJPEG(xRefTable *XRefTable, c image.Config, bb bytes.Buffer) ([]ImageResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decodeImage(xRefTable *XRefTable, buf *bytes.Reader, currentOffset int64, gray, sepia bool, byteOrder binary.ByteOrder, imgResources *[]ImageResource) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func createImageResourcesForTIFF(xRefTable *XRefTable, bb bytes.Buffer, gray, sepia bool) ([]ImageResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createImageResources(xRefTable *XRefTable, c image.Config, bb bytes.Buffer, gray, sepia bool) ([]ImageResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateImageResources(xRefTable *XRefTable, r io.Reader, gray, sepia bool) ([]ImageResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateImageStreamDict(xRefTable *XRefTable, r io.Reader) (*types.StreamDict, int, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, nil
}

func hasAlpha(cm color.Model) bool { _ = "STUB: not implemented"; return false }

func CreateImageResource(xRefTable *XRefTable, r io.Reader) (*types.IndirectRef, int, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, nil
}
