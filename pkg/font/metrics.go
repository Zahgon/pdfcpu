package font

import (
	"os"
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

func validatePostScriptName(name string) error { _ = "STUB: not implemented"; return nil }

func validateMetricCounts(fd TTFLight) error { _ = "STUB: not implemented"; return nil }

func validateFontBoundingBox(fd TTFLight) error { _ = "STUB: not implemented"; return nil }

func validUnicodeScalar(r uint32) bool { _ = "STUB: not implemented"; return false }

func validateUnicodeMaps(fd TTFLight) error { _ = "STUB: not implemented"; return nil }

func validateUnicodePlanes(fd TTFLight) error { _ = "STUB: not implemented"; return nil }

func ValidateTTFLight(fd TTFLight) error { _ = "STUB: not implemented"; return nil }

func (fd TTFLight) String() string { _ = "STUB: not implemented"; return "" }

func (fd TTFLight) supportsUnicodeBlock(bit int) bool { _ = "STUB: not implemented"; return false }

func (fd TTFLight) supportsUnicodeBlocks(bits []int) bool { _ = "STUB: not implemented"; return false }

func (fd TTFLight) unicodeRangeBits(id string) []int { _ = "STUB: not implemented"; return nil }

func (fd TTFLight) SupportsScript(id string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

var UserFontDir string

var userFontMetrics = map[string]TTFLight{}
var userFontMetricsLock = &sync.RWMutex{}

var loadUserFontsOnce sync.Once
var loadUserFontsMutex sync.Mutex
var loadUserFontsErr error

func openInstalledGob(fileName string) (f *os.File, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func load(fileName string, fd *TTFLight) (err error) { _ = "STUB: not implemented"; return nil }

func Read(fileName string) (bb []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func isSupportedFontFile(filename string) bool { _ = "STUB: not implemented"; return false }

func doLoadUserFonts() error { _ = "STUB: not implemented"; return nil }

func LoadUserFonts() error { _ = "STUB: not implemented"; return nil }

func ReloadUserFonts() error { _ = "STUB: not implemented"; return nil }

func BoundingBox(fontName string) (*types.Rectangle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CharWidth(fontName string, r rune) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func UserSpaceUnits(glyphSpaceUnits float64, fontScalingFactor int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func UserSpaceUnitsFloat(glyphSpaceUnits, fontScalingFactor float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func GlyphSpaceUnits(userSpaceUnits float64, fontScalingFactor int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func GlyphSpaceUnitsFloat(userSpaceUnits, fontScalingFactor float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func fontScalingFactor(glyphSpaceUnits, userSpaceUnits float64) int {
	_ = "STUB: not implemented"
	return 0
}

func Descent(fontName string, fontSize int) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func Ascent(fontName string, fontSize int) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func LineHeight(fontName string, fontSize int) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func glyphSpaceWidth(text, fontName string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func TextWidth(text, fontName string, fontSize int) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func TextWidthFloat(text, fontName string, fontSize float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func Size(text, fontName string, width float64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func SizeForLineHeight(fontName string, lh float64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func UserSpaceFontBBox(fontName string, fontSize int) (*types.Rectangle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func IsCoreFont(fontName string) bool { _ = "STUB: not implemented"; return false }

func CoreFontNames() []string { _ = "STUB: not implemented"; return nil }

func IsUserFont(fontName string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func cloneTTFLight(ttf TTFLight) TTFLight { _ = "STUB: not implemented"; return *new(TTFLight) }

func userFont(fontName string) (TTFLight, bool, error) {
	_ = "STUB: not implemented"
	return *new(TTFLight), false, nil
}

func UserFont(fontName string) (TTFLight, bool, error) {
	_ = "STUB: not implemented"
	return *new(TTFLight), false, nil
}

func UserFontNames() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func UserFontNamesVerbose() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func SupportedFont(fontName string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (fd TTFLight) Gids() []int { _ = "STUB: not implemented"; return nil }
