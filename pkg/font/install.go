package font

import (
	"bytes"
	"io"
)

const (
	sfntVersionTrueType      = "\x00\x01\x00\x00"
	sfntVersionTrueTypeApple = "true"
	sfntVersionCFF           = "OTTO"
	ttfHeadMagicNumber       = 0x5F0F3CF5
	ttcTag                   = "ttcf"
)

type ttf struct {
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
	FontFile           []byte
}

func (fd ttf) String() string { _ = "STUB: not implemented"; return "" }

func (fd ttf) toPDFGlyphSpace(i int) int { _ = "STUB: not implemented"; return 0 }

type myUint32 []uint32

func (f myUint32) Len() int { _ = "STUB: not implemented"; return 0 }

func (f myUint32) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (f myUint32) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (fd ttf) PrintChars() string { _ = "STUB: not implemented"; return "" }

type table struct {
	chksum uint32
	off    uint32
	size   uint32
	padded uint32
	data   []byte
}

func (t table) uint16(off int) uint16 { _ = "STUB: not implemented"; return 0 }

func (t table) int16(off int) int16 { _ = "STUB: not implemented"; return 0 }

func (t table) uint32(off int) uint32 { _ = "STUB: not implemented"; return 0 }

func (t table) fixed32(off int) float64 { _ = "STUB: not implemented"; return 0 }

func (t table) parseFontHeaderTable(fd *ttf) error { _ = "STUB: not implemented"; return nil }

func uint16ToBigEndianBytes(i uint16) []byte { _ = "STUB: not implemented"; return nil }

func uint32ToBigEndianBytes(i uint32) []byte { _ = "STUB: not implemented"; return nil }

func utf16BEToString(bb []byte) string { _ = "STUB: not implemented"; return "" }

func (t table) parsePostScriptTable(fd *ttf) error { _ = "STUB: not implemented"; return nil }

func (t table) parseWindowsMetricsTable(fd *ttf) error { _ = "STUB: not implemented"; return nil }

func (t table) parseNamingTable(fd *ttf) error { _ = "STUB: not implemented"; return nil }

func (t table) parseHorizontalHeaderTable(fd *ttf) error { _ = "STUB: not implemented"; return nil }

func (t table) parseMaximumProfile(fd *ttf) error { _ = "STUB: not implemented"; return nil }

func (t table) parseHorizontalMetricsTable(fd *ttf) error { _ = "STUB: not implemented"; return nil }

func (t table) parseCMapFormat4(fd *ttf) error { _ = "STUB: not implemented"; return nil }

func (t table) parseCMapFormat12(fd *ttf) error { _ = "STUB: not implemented"; return nil }

func (t table) parseCharToGlyphMappingTable(fd *ttf) error { _ = "STUB: not implemented"; return nil }

func calcTableChecksum(tag string, b []byte) uint32 { _ = "STUB: not implemented"; return 0 }

func getNext32BitAlignedLength(i uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func headerAndTables(fn string, r io.ReaderAt, baseOff int64) ([]byte, map[string]*table, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func parse(tags map[string]*table, tag string, fd *ttf) error {
	_ = "STUB: not implemented"
	return nil
}

func writeGob(fileName string, fd ttf) error { _ = "STUB: not implemented"; return nil }

func readGob(fileName string, fd *ttf) error { _ = "STUB: not implemented"; return nil }

func installTrueTypeRep(fontDir, fontName string, header []byte, tables map[string]*table) error {
	_ = "STUB: not implemented"
	return nil
}

func InstallTrueTypeCollection(fontDir, fn string) error { _ = "STUB: not implemented"; return nil }

func InstallTrueTypeFont(fontDir, fontName string) error { _ = "STUB: not implemented"; return nil }

func InstallFontFromBytes(fontDir, fontName string, bb []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func ttfTables(tableCount int, bb []byte) (map[string]*table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func glyfOffset(loca *table, gid, indexToLocFormat int) int { _ = "STUB: not implemented"; return 0 }

func writeGlyfOffset(buf *bytes.Buffer, off, indexToLocFormat int) {
	_ = "STUB: not implemented"
	return
}

func pad(bb []byte) []byte { _ = "STUB: not implemented"; return nil }

func glyphOffsets(gid int, locaFull, glyfsFull *table, numGlyphs, indexToLocFormat int) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func resolveCompoundGlyph(fontName string, bb []byte, usedGIDs map[uint16]bool,
	locaFull, glyfsFull *table, numGlyphs, indexToLocFormat int) error {
	_ = "STUB: not implemented"
	return nil
}

func resolveCompoundGlyphs(fontName string, usedGIDs map[uint16]bool, locaFull, glyfsFull *table, numGlyphs, indexToLocFormat int) error {
	_ = "STUB: not implemented"
	return nil
}

func glyfAndLoca(fontName string, tables map[string]*table, usedGIDs map[uint16]bool) error {
	_ = "STUB: not implemented"
	return nil
}

func createTTF(header []byte, tables map[string]*table) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Subset(fontName string, usedGIDs map[uint16]bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
