package model

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type Resize struct {
	Scale         float64
	Unit          types.DisplayUnit
	PageDim       *types.Dim
	PageSize      string
	EnforceOrient bool
	UserDim       bool
	Border        bool
	BgColor       *color.SimpleColor
}

func (r Resize) EnforceOrientation() bool { _ = "STUB: not implemented"; return false }

func parsePageDimRes(v string, u types.DisplayUnit) (*types.Dim, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func parseDimensionsRes(s string, res *Resize) (err error) { _ = "STUB: not implemented"; return nil }

func parseEnforceOrientation(s string, res *Resize) error { _ = "STUB: not implemented"; return nil }

func parsePageFormatRes(s string, res *Resize) error { _ = "STUB: not implemented"; return nil }

func parseScaleFactorSimple(s string) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func parseScaleFactorRes(s string, res *Resize) (err error) { _ = "STUB: not implemented"; return nil }

func parseBackgroundColorRes(s string, res *Resize) error { _ = "STUB: not implemented"; return nil }

func parseBorderRes(s string, res *Resize) error { _ = "STUB: not implemented"; return nil }

type resizeParameterMap map[string]func(string, *Resize) error

var ResizeParamMap = resizeParameterMap{
	"dimensions":  parseDimensionsRes,
	"enforce":     parseEnforceOrientation,
	"formsize":    parsePageFormatRes,
	"papersize":   parsePageFormatRes,
	"scalefactor": parseScaleFactorRes,
	"bgcolor":     parseBackgroundColorRes,
	"border":      parseBorderRes,
}
