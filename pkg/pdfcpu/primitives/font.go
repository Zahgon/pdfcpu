package primitives

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type FormFont struct {
	pdf      *PDF
	Name     string
	Lang     string
	Script   string
	Size     int
	Color    string `json:"col"`
	col      *color.SimpleColor
	FillFont bool
}

var ISO639Codes = []string{"ab", "aa", "af", "ak", "sq", "am", "ar", "an", "hy", "as", "av", "ae", "ay", "az", "bm", "ba", "eu", "be", "bn", "bi", "bs", "br", "bg",
	"my", "ca", "ch", "ce", "ny", "zh", "cu", "cv", "kw", "co", "cr", "hr", "cs", "da", "dv", "nl", "dz", "en", "eo", "et", "ee", "fo", "fj", "fi", "fr", "fy", "ff",
	"gd", "gl", "lg", "ka", "de", "el", "kl", "gn", "gu", "ht", "ha", "he", "hz", "hi", "ho", "hu", "is", "io", "ig", "id", "ia", "ie", "iu", "ik", "ga", "it", "ja",
	"jv", "kn", "kr", "ks", "kk", "km", "ki", "rw", "ky", "kv", "kg", "ko", "kj", "ku", "lo", "la", "lv", "li", "ln", "lt", "lu", "lb", "mk", "mg", "ms", "ml", "mt",
	"gv", "mi", "mr", "mh", "mn", "na", "nv", "nd", "nr", "ng", "ne", "no", "nb", "nn", "ii", "oc", "oj", "or", "om", "os", "pi", "ps", "fa", "pl", "pt", "pa", "qu",
	"ro", "rm", "rn", "ru", "se", "sm", "sg", "sa", "sc", "sr", "sn", "sd", "si", "sk", "sl", "so", "st", "es", "su", "sw", "ss", "sv", "tl", "ty", "tg", "ta", "tt",
	"te", "th", "bo", "ti", "to", "ts", "tn", "tr", "tk", "tw", "ug", "uk", "ur", "uz", "ve", "vi", "vo", "wa", "cy", "wo", "xh", "yi", "yo", "za", "zu"}

func (f *FormFont) validateISO639() error { _ = "STUB: not implemented"; return nil }

func (f *FormFont) validateScriptSupport() error { _ = "STUB: not implemented"; return nil }

func (f *FormFont) validate() error { _ = "STUB: not implemented"; return nil }

func (f *FormFont) mergeIn(f0 *FormFont) { _ = "STUB: not implemented"; return }

func (f *FormFont) SetCol(c color.SimpleColor) { _ = "STUB: not implemented"; return }

func (f FormFont) RTL() bool { _ = "STUB: not implemented"; return false }

func FormFontDetails(xRefTable *model.XRefTable, indRef types.IndirectRef) (string, string, string, error) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

func FormFontResDict(xRefTable *model.XRefTable) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func formFontIndRef(xRefTable *model.XRefTable, fontID string) *types.IndirectRef {
	_ = "STUB: not implemented"
	return nil
}

func FontIndRef(fName string, ctx *model.Context, fonts map[string]types.IndirectRef) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ensureUTF8FormFont(ctx *model.Context, fonts map[string]types.IndirectRef) (string, string, string, string, *types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return "", "", "", "", nil, nil
}

func extractFormFontDetails(
	ctx *model.Context,
	fontID string,
	fonts map[string]types.IndirectRef) (string, string, string, string, *types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return "", "", "", "", nil, nil
}

func fontFromDA(s string) (string, FormFont, error) {
	_ = "STUB: not implemented"
	return "", *new(FormFont), nil
}

func calcFontDetailsFromDA(ctx *model.Context, d types.Dict, da *string, needUTF8 bool, fonts map[string]types.IndirectRef) (string, *FormFont, bool, *types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return "", nil, false, nil, nil
}
