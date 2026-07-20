package font

import (
	"sync"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type TTFLight struct {
	PostscriptName     string
	Protected          bool
	UnitsPerEm         int
	Ascent             int
	Descent            int
	CapHeight          int
	FirstChar          uint16
	LastChar           uint16
	UnicodeRange       [4]uint32
	LLx, LLy, URx, URy float64
	ItalicAngle        float64
	FixedPitch         bool
	Bold               bool
	HorMetricsCount    int
	GlyphCount         int
	GlyphWidths        []int
	Chars              map[uint32]uint16
	ToUnicode          map[uint16]uint32
	Planes             map[int]bool
}

func (fd TTFLight) String() string { _ = "STUB: not implemented"; return "" }

func (fd TTFLight) supportsUnicodeBlock(bit int) bool { _ = "STUB: not implemented"; return false }

func (fd TTFLight) supportsUnicodeBlocks(bits []int) bool { _ = "STUB: not implemented"; return false }

func (fd TTFLight) unicodeRangeBits(id string) []int { _ = "STUB: not implemented"; return nil }

func (fd TTFLight) SupportsScript(id string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

var UserFontDir string

var UserFontMetrics = map[string]TTFLight{}
var UserFontMetricsLock = &sync.RWMutex{}

var loadUserFontsOnce sync.Once
var loadUserFontsErr error

func load(fileName string, fd *TTFLight) error { _ = "STUB: not implemented"; return nil }

func Read(fileName string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func isSupportedFontFile(filename string) bool { _ = "STUB: not implemented"; return false }

func doLoadUserFonts() error { _ = "STUB: not implemented"; return nil }

func LoadUserFonts() error { _ = "STUB: not implemented"; return nil }

func EnsureUserFontsLoaded() { _ = "STUB: not implemented"; return }

func BoundingBox(fontName string) *types.Rectangle { _ = "STUB: not implemented"; return nil }

func CharWidth(fontName string, r rune) int { _ = "STUB: not implemented"; return 0 }

func UserSpaceUnits(glyphSpaceUnits float64, fontScalingFactor int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func GlyphSpaceUnits(userSpaceUnits float64, fontScalingFactor int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func fontScalingFactor(glyphSpaceUnits, userSpaceUnits float64) int {
	_ = "STUB: not implemented"
	return 0
}

func Descent(fontName string, fontSize int) float64 { _ = "STUB: not implemented"; return 0 }

func Ascent(fontName string, fontSize int) float64 { _ = "STUB: not implemented"; return 0 }

func LineHeight(fontName string, fontSize int) float64 { _ = "STUB: not implemented"; return 0 }

func glyphSpaceWidth(text, fontName string) int { _ = "STUB: not implemented"; return 0 }

func TextWidth(text, fontName string, fontSize int) float64 { _ = "STUB: not implemented"; return 0 }

func Size(text, fontName string, width float64) int { _ = "STUB: not implemented"; return 0 }

func SizeForLineHeight(fontName string, lh float64) int { _ = "STUB: not implemented"; return 0 }

func UserSpaceFontBBox(fontName string, fontSize int) *types.Rectangle {
	_ = "STUB: not implemented"
	return nil
}

func IsCoreFont(fontName string) bool { _ = "STUB: not implemented"; return false }

func CoreFontNames() []string { _ = "STUB: not implemented"; return nil }

func IsUserFont(fontName string) bool { _ = "STUB: not implemented"; return false }

func UserFontNames() []string { _ = "STUB: not implemented"; return nil }

func UserFontNamesVerbose() []string { _ = "STUB: not implemented"; return nil }

func SupportedFont(fontName string) bool { _ = "STUB: not implemented"; return false }

func (fd TTFLight) Gids() []int { _ = "STUB: not implemented"; return nil }
