package pdfcpu

import (
	"image"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"github.com/pkg/errors"
)

var (
	ErrUnsupported16BPC = errors.New("unsupported 16 bits per component")
)

type colValRange struct {
	min, max float64
}

type PDFImage struct {
	objNr     int
	sd        *types.StreamDict
	comp      int
	bpc       int
	w, h      int
	softMask  []byte
	decode    []colValRange
	imageMask bool
	thumb     bool
}

func decodeArr(a types.Array) []colValRange { _ = "STUB: not implemented"; return nil }

func checkedImageBytes(w, h, comp, bpc int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func imageLimits(xRefTable *model.XRefTable) model.ResourceLimits {
	_ = "STUB: not implemented"
	return *new(model.ResourceLimits)
}

func validatePDFImageDimensions(xRefTable *model.XRefTable, w, h, comp, bpc, objNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func pdfImage(xRefTable *model.XRefTable, sd *types.StreamDict, thumb bool, objNr int) (*PDFImage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func colorLookupTable(xRefTable *model.XRefTable, o types.Object) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func maxValForBits(bpc int) int { _ = "STUB: not implemented"; return 0 }

func decodePixelValue(v uint8, bpc int, r colValRange) uint8 { _ = "STUB: not implemented"; return 0 }

func streamBytes(sd *types.StreamDict) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func softMask(xRefTable *model.XRefTable, d *types.StreamDict, w, h, objNr int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func imageForCMYKWithoutSoftMask(im *PDFImage) image.Image {
	_ = "STUB: not implemented"
	return *new(image.Image)
}

func imageForCMYKWithSoftMask(im *PDFImage) image.Image {
	_ = "STUB: not implemented"
	return *new(image.Image)
}

func renderDeviceCMYKToTIFF(im *PDFImage) (io.Reader, string, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), "", nil
}

func scaleToBPC8(v uint8, bpc int) uint8 { _ = "STUB: not implemented"; return 0 }

func renderDeviceGrayToPNG(im *PDFImage) (io.Reader, string, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), "", nil
}

func renderDeviceRGBToPNG(im *PDFImage) (io.Reader, string, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), "", nil
}

func renderCalRGBToPNG(im *PDFImage) (io.Reader, string, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), "", nil
}

func renderICCBased(xRefTable *model.XRefTable, im *PDFImage, cs types.Array) (io.Reader, string, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), "", nil
}

func renderIndexedGrayToPNG(im *PDFImage, lookup []byte) (io.Reader, string, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), "", nil
}

func renderIndexedRGBToPNG(im *PDFImage, lookup []byte) (io.Reader, string, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), "", nil
}

func imageForIndexedCMYKWithoutSoftMask(im *PDFImage, lookup []byte) image.Image {
	_ = "STUB: not implemented"
	return *new(image.Image)
}

func imageForIndexedCMYKWithSoftMask(im *PDFImage, lookup []byte) image.Image {
	_ = "STUB: not implemented"
	return *new(image.Image)
}

func renderIndexedCMYKToTIFF(im *PDFImage, lookup []byte) (io.Reader, string, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), "", nil
}

func renderIndexedNameCS(im *PDFImage, cs types.Name, maxInd int, lookup []byte) (io.Reader, string, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), "", nil
}

func renderIndexedArrayCS(xRefTable *model.XRefTable, im *PDFImage, csa types.Array, maxInd int, lookup []byte) (io.Reader, string, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), "", nil
}

func renderIndexed(xRefTable *model.XRefTable, im *PDFImage, cs types.Array) (io.Reader, string, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), "", nil
}

func renderDeviceN(im *PDFImage, cs types.Array) (io.Reader, string, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), "", nil
}

func renderImage(xRefTable *model.XRefTable, sd *types.StreamDict, thumb bool, objNr int) (io.Reader, string, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), "", nil
}

func decodeCMYK(c, m, y, k uint8, decode []colValRange) (uint8, uint8, uint8, uint8) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

func renderCMYKToPng(im *PDFImage) (io.Reader, string, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), "", nil
}

func renderDCTToPNG(xRefTable *model.XRefTable, sd *types.StreamDict, thumb bool, objNr int) (io.Reader, string, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), "", nil
}

func RenderImage(xRefTable *model.XRefTable, sd *types.StreamDict, thumb bool, resourceName string, objNr int) (io.Reader, string, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), "", nil
}

func WriteReader(path string, r io.Reader) error { _ = "STUB: not implemented"; return nil }

func WriteImage(xRefTable *model.XRefTable, fileName string, sd *types.StreamDict, thumb bool, objNr int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
