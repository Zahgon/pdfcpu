package model

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type Cut struct {
	Hor      []float64
	Vert     []float64
	Scale    float64
	PageSize string
	PageDim  *types.Dim
	Unit     types.DisplayUnit
	UserDim  bool
	Border   bool
	Margin   float64
	BgColor  *color.SimpleColor
	Origin   types.Corner
}

type cutParameterMap map[string]func(string, *Cut) error

func parseHorCut(v string, cut *Cut) (err error) { _ = "STUB: not implemented"; return nil }

func parseVertCut(v string, cut *Cut) (err error) { _ = "STUB: not implemented"; return nil }

func parsePageDimCut(v string, u types.DisplayUnit) (*types.Dim, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func parseDimensionsCut(s string, cut *Cut) (err error) { _ = "STUB: not implemented"; return nil }

func parsePageFormatCut(s string, cut *Cut) error { _ = "STUB: not implemented"; return nil }

func parseScaleFactorCut(s string, cut *Cut) (err error) { _ = "STUB: not implemented"; return nil }

func parseBackgroundColorCut(s string, cut *Cut) error { _ = "STUB: not implemented"; return nil }

func parseBorderCut(s string, cut *Cut) error { _ = "STUB: not implemented"; return nil }

func parseMarginCut(s string, cut *Cut) error { _ = "STUB: not implemented"; return nil }

var CutParamMap = cutParameterMap{
	"horizontalCut": parseHorCut,
	"verticalCut":   parseVertCut,
	"dimensions":    parseDimensionsCut,
	"formsize":      parsePageFormatCut,
	"papersize":     parsePageFormatCut,
	"scalefactor":   parseScaleFactorCut,
	"border":        parseBorderCut,
	"margin":        parseMarginCut,
	"bgcolor":       parseBackgroundColorCut,
}
