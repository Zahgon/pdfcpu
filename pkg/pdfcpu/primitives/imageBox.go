package primitives

import (
	"io"
	"net/url"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/matrix"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

var imageBoxUserAgent = "pdfcpu/" + model.VersionStr + " (+https://github.com/pdfcpu/pdfcpu)"

type ImageData struct {
	Payload       string
	Format        string
	Width, Height int
}

type ImageBox struct {
	pdf             *PDF
	content         *Content
	Name            string
	Src             string `json:"src"`
	Data            *ImageData
	Position        [2]float64 `json:"pos"`
	x, y            float64
	Dx, Dy          float64
	dest            *types.Rectangle
	Anchor          string
	anchor          types.Anchor
	anchored        bool
	Width           float64
	Height          float64
	Margin          *Margin
	Border          *Border
	Padding         *Padding
	BackgroundColor string `json:"bgCol"`
	bgCol           *color.SimpleColor
	Rotation        float64 `json:"rot"`
	Url             string
	Hide            bool
	PageNr          string `json:"-"`
}

func (ib *ImageBox) resolveFileName(s string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ib *ImageBox) parseAnchor() (types.Anchor, error) {
	_ = "STUB: not implemented"
	return *new(types.Anchor), nil
}

func (ib *ImageBox) validate() error { _ = "STUB: not implemented"; return nil }

func (ib *ImageBox) margin(name string) *Margin { _ = "STUB: not implemented"; return nil }

func (ib *ImageBox) border(name string) *Border { _ = "STUB: not implemented"; return nil }

func (ib *ImageBox) padding(name string) *Padding { _ = "STUB: not implemented"; return nil }

func (ib *ImageBox) missingPosition() bool { _ = "STUB: not implemented"; return false }

func (ib *ImageBox) mergeIn(ib0 *ImageBox) { _ = "STUB: not implemented"; return }

func (ib *ImageBox) cachedImg(img model.ImageResource, pageImages model.ImageMap, pageNr int) (int, int, string, error) {
	_ = "STUB: not implemented"
	return 0, 0, "", nil
}

func (ib *ImageBox) checkForExistingImage(sd *types.StreamDict, w, h int) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ib *ImageBox) resource() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (ib *ImageBox) remoteResource(u *url.URL) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (ib *ImageBox) imageResource(pageImages, images model.ImageMap, pageNr int) (*model.ImageResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ib *ImageBox) image(pageImages, images model.ImageMap, pageNr int) (int, int, string, error) {
	_ = "STUB: not implemented"
	return 0, 0, "", nil
}

func (ib *ImageBox) createLink(p *model.Page, pageNr int, r *types.Rectangle, m matrix.Matrix) {
	_ = "STUB: not implemented"
	return
}

func (ib *ImageBox) prepareMargin() (float64, float64, float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, nil
}

func (ib *ImageBox) prepareBorder() (float64, *color.SimpleColor, types.LineJoinStyle, error) {
	_ = "STUB: not implemented"
	return 0, nil, *new(types.LineJoinStyle), nil
}

func (ib *ImageBox) preparePadding() (float64, float64, float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, nil
}

func (ib *ImageBox) calcDim(rSrc, r *types.Rectangle, bWidth, pTop, pBot, pLeft, pRight float64) {
	_ = "STUB: not implemented"
	return
}

func (ib *ImageBox) calcTransform(
	mLeft, mBot, mRight, mTop,
	pLeft, pBot, pRight, pTop,
	bWidth float64, rSrc *types.Rectangle) (matrix.Matrix, float64, float64, float64, float64, *types.Rectangle) {
	_ = "STUB: not implemented"
	return *new(matrix.Matrix), 0, 0, 0, 0, nil
}

func (ib *ImageBox) render(p *model.Page, pageNr int, images model.ImageMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (ib *ImageBox) RenderForFill(pdf *PDF, p *model.Page, pageNr int, imageMap model.ImageMap) error {
	_ = "STUB: not implemented"
	return nil
}
