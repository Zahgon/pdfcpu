package model

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type Zoom struct {
	Factor  float64
	HMargin float64
	VMargin float64
	Unit    types.DisplayUnit
	Border  bool
	BgColor *color.SimpleColor
}

func (z *Zoom) EnsureFactorAndMargins(w, h float64) error { _ = "STUB: not implemented"; return nil }

func parseHMargin(s string, zoom *Zoom) error { _ = "STUB: not implemented"; return nil }

func parseVMargin(s string, zoom *Zoom) error { _ = "STUB: not implemented"; return nil }

func parseZoomFactor(s string, zoom *Zoom) (err error) { _ = "STUB: not implemented"; return nil }

func parseBackgroundColorZoom(s string, zoom *Zoom) error { _ = "STUB: not implemented"; return nil }

func parseBorderZoom(s string, zoom *Zoom) error { _ = "STUB: not implemented"; return nil }

type zoomParameterMap map[string]func(string, *Zoom) error

var ZoomParamMap = zoomParameterMap{
	"factor":  parseZoomFactor,
	"hmargin": parseHMargin,
	"vmargin": parseVMargin,
	"bgcolor": parseBackgroundColorZoom,
	"border":  parseBorderZoom,
}
