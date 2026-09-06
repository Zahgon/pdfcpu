package font

import (
	"bytes"
	"errors"
	"io"
	"os"
)

const (
	sfntVersionTrueType      = "\x00\x01\x00\x00"
	sfntVersionTrueTypeApple = "true"
	sfntVersionCFF           = "OTTO"
	ttfHeadMagicNumber       = 0x5F0F3CF5
	ttcTag                   = "ttcf"
	maxFontFileSize          = 256 << 20
	maxFontTableSize         = 128 << 20
	maxInstalledFontSize     = 512 << 20
	maxFontTableCount        = 1024
	maxFontCollectionFonts   = 256
	installedFontMode        = 0644
)

var (
	ErrMissingFontDir = errors.New("missing font directory")

	ErrMissingFontName = errors.New("missing font name")

	ErrMissingFontData = errors.New("missing font data")

	ErrInvalidFontData = errors.New("invalid font data")

	ErrUnsupportedFontFormat = errors.New("unsupported font format")

	ErrUnknownFont = errors.New("unknown font")

	ErrDuplicatePostScriptName = errors.New("duplicate PostScript name")
)

type InstallResult struct {
	PostScriptName string
	Member         int
}

type InstallReport struct {
	Fonts    []InstallResult
	Warnings []error
}

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

type ttfComparable struct {
	PostscriptName  string
	Protected       bool
	UnitsPerEm      int
	Ascent          int
	Descent         int
	CapHeight       int
	FirstChar       uint16
	LastChar        uint16
	UnicodeRange    [4]uint32
	LLx             float64
	LLy             float64
	URx             float64
	URy             float64
	ItalicAngle     float64
	FixedPitch      bool
	Bold            bool
	HorMetricsCount int
	GlyphCount      int
}

func comparableTTF(fd ttf) ttfComparable { _ = "STUB: not implemented"; return *new(ttfComparable) }

func (fd ttf) String() string { _ = "STUB: not implemented"; return "" }

func ttfEqual(fd1, fd2 ttf) bool { _ = "STUB: not implemented"; return false }

func (fd ttf) toPDFGlyphSpace(i int) int { _ = "STUB: not implemented"; return 0 }

func requireUnitsPerEm(fd *ttf, tag string) error { _ = "STUB: not implemented"; return nil }

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

func (t table) logicalLength(tag string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (t table) requireSize(tag string, size int) error { _ = "STUB: not implemented"; return nil }

func (t table) slice(tag string, off, size int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

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

type cmapFormat4Layout struct {
	segCount int
	endOff   int
	startOff int
	deltaOff int
	rangeOff int
}

func prepareCMapFormat4(t table) (table, cmapFormat4Layout, error) {
	_ = "STUB: not implemented"
	return *new(table), *new(cmapFormat4Layout), nil
}

func cmapFormat4Glyph(t table, layout cmapFormat4Layout, segment int, code uint32, index int) (uint16, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func parseCMapFormat4Segment(t table, fd *ttf, layout cmapFormat4Layout, segment int) error {
	_ = "STUB: not implemented"
	return nil
}

func (t table) parseCMapFormat4(fd *ttf) error { _ = "STUB: not implemented"; return nil }

func prepareCMapFormat12(t table) (table, uint32, error) {
	_ = "STUB: not implemented"
	return *new(table), 0, nil
}

func cmapFormat12Group(t table, fd *ttf, group uint32) (uint32, uint32, uint32, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, nil
}

func recordCMapPlanes(planes map[int]bool, startCode, endCode, prevCode uint32, first bool) {
	_ = "STUB: not implemented"
	return
}

func addCMapFormat12Group(fd *ttf, startCode, endCode, startGlyphID uint32) {
	_ = "STUB: not implemented"
	return
}

func (t table) parseCMapFormat12(fd *ttf) error { _ = "STUB: not implemented"; return nil }

func cmapSubtableLength(t table, record, start int, format uint16) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func cmapSubtable(t table, record int) (string, table, bool, error) {
	_ = "STUB: not implemented"
	return "", *new(table), false, nil
}

func (t table) parseCharToGlyphMappingTable(fd *ttf) error { _ = "STUB: not implemented"; return nil }

func calcTableChecksum(tag string, b []byte) uint32 { _ = "STUB: not implemented"; return 0 }

func getNext32BitAlignedLength(i uint32) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func readFontHeader(r io.ReaderAt, baseOff, size int64) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func readTableDirectory(r io.ReaderAt, baseOff, size int64, c int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readFontTable(r io.ReaderAt, size int64, b []byte) (string, *table, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func validateTableDirectory(b []byte, fontSize int64, tableCount int) error {
	_ = "STUB: not implemented"
	return nil
}

func headerAndTables(r io.ReaderAt, baseOff, size int64) ([]byte, map[string]*table, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func parse(tags map[string]*table, tag string, fd *ttf) error {
	_ = "STUB: not implemented"
	return nil
}

type gobPersistenceOperations struct {
	createTemp func(string, string) (*os.File, error)
	encode     func(*os.File, ttf) error
	chmod      func(*os.File, os.FileMode) error
	sync       func(*os.File) error
	syncDir    func(string) error
	close      func(*os.File) error
	rename     func(string, string) error
	remove     func(string) error
	verify     func(string, *ttf) error
}

func defaultGobPersistenceOperations() gobPersistenceOperations {
	_ = "STUB: not implemented"
	return *new(gobPersistenceOperations)
}

func encodeGobFile(f *os.File, fileName string, fd ttf, ops gobPersistenceOperations) error {
	_ = "STUB: not implemented"
	return nil
}

func removeTemporaryFont(fileName string, remove func(string) error) error {
	_ = "STUB: not implemented"
	return nil
}

func writeGobWithOperations(fileName string, fd ttf, ops gobPersistenceOperations) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func writeGob(fileName string, fd ttf) error { _ = "STUB: not implemented"; return nil }

func readGob(fileName string, fd *ttf) (err error) { _ = "STUB: not implemented"; return nil }

func (fd ttf) light() TTFLight { _ = "STUB: not implemented"; return *new(TTFLight) }

func validateDecodedFontFile(bb []byte) error { _ = "STUB: not implemented"; return nil }

func validateDecodedTTF(fd ttf) error { _ = "STUB: not implemented"; return nil }

func validateInstallTarget(fontDir, fontName string) error { _ = "STUB: not implemented"; return nil }

func validateInstallData(fontDir, fontName string, bb []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func installTrueTypeRep(fontDir, fontName string, header []byte, tables map[string]*table, logInstall bool, reserve func(string) error) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func readTrueTypeCollectionHeader(f *os.File, fn string) (int64, uint32, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, nil
}

func trueTypeCollectionMemberOffset(f *os.File, fn string, index uint32, size, offsetTableEnd int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type collectionInstallFileOperations struct {
	mkdirTemp     func(string, string) (string, error)
	lstat         func(string) (os.FileInfo, error)
	close         func(*os.File) error
	syncDir       func(string) error
	rename        func(string, string) error
	remove        func(string) error
	removeAll     func(string) error
	stageMembers  func(*os.File, string, string, int64, uint32, int64) ([]InstallResult, error)
	reportWarning func(error)
}

type committedCollectionFont struct {
	name        string
	hadOriginal bool
	committed   bool
}

func defaultCollectionInstallFileOperations() collectionInstallFileOperations {
	_ = "STUB: not implemented"
	return *new(collectionInstallFileOperations)
}

func syncCollectionDirectories(ops collectionInstallFileOperations, dirs ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func rollbackCollectionFonts(fontDir, backupDir string, files []committedCollectionFont, ops collectionInstallFileOperations) error {
	_ = "STUB: not implemented"
	return nil
}

func commitCollectionFonts(fontDir, stagingDir string, results []InstallResult, ops collectionInstallFileOperations) error {
	_ = "STUB: not implemented"
	return nil
}

func reserveCollectionPostScriptName(members map[string]int, postScriptName string, member int) error {
	_ = "STUB: not implemented"
	return nil
}

func installTrueTypeCollectionMembers(f *os.File, fontDir, fn string, size int64, count uint32, offsetTableEnd int64) ([]InstallResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func InstallTrueTypeCollection(fontDir, fn string) (InstallReport, error) {
	_ = "STUB: not implemented"
	return *new(InstallReport), nil
}

func installTrueTypeCollectionResults(fontDir, fn string, ops collectionInstallFileOperations) (report InstallReport, err error) {
	_ = "STUB: not implemented"
	return *new(InstallReport), nil
}

func InstallTrueTypeCollectionResults(fontDir, fn string) (InstallReport, error) {
	_ = "STUB: not implemented"
	return *new(InstallReport), nil
}

func InstallTrueTypeFont(fontDir, fontName string) (InstallReport, error) {
	_ = "STUB: not implemented"
	return *new(InstallReport), nil
}

type trueTypeFontInstallOperations struct {
	close           func(*os.File) error
	headerAndTables func(io.ReaderAt, int64, int64) ([]byte, map[string]*table, error)
	installRep      func(string, string, []byte, map[string]*table, bool, func(string) error) (string, error)
}

func defaultTrueTypeFontInstallOperations() trueTypeFontInstallOperations {
	_ = "STUB: not implemented"
	return *new(trueTypeFontInstallOperations)
}

func InstallTrueTypeFontResult(fontDir, fontName string) (InstallReport, error) {
	_ = "STUB: not implemented"
	return *new(InstallReport), nil
}

func installTrueTypeFontResult(fontDir, fontName string, ops trueTypeFontInstallOperations) (report InstallReport, err error) {
	_ = "STUB: not implemented"
	return *new(InstallReport), nil
}

func InstallFontFromBytes(fontDir, fontName string, bb []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func InstallFontFromBytesQuiet(fontDir, fontName string, bb []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func installFontFromBytes(fontDir, fontName string, bb []byte, logInstall bool) error {
	_ = "STUB: not implemented"
	return nil
}

func ttfTables(tableCount int, bb []byte) (map[string]*table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const (
	compoundArgWords       = 0x0001
	compoundScale          = 0x0008
	compoundMoreComponents = 0x0020
	compoundXYScale        = 0x0040
	compoundTwoByTwo       = 0x0080
	compoundInstructions   = 0x0100
	maxCompoundGlyphDepth  = 64
)

type subsetGlyphTables struct {
	loca          *table
	glyf          *table
	numGlyphs     int
	locaFormat    int
	locaEntrySize int
}

func tableUint16At(t *table, off int) (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

func tableUint32At(t *table, off int) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func requiredGlyphTable(tables map[string]*table, tag string, minSize int) (*table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func subsetGlyphTableInfo(tables map[string]*table) (subsetGlyphTables, error) {
	_ = "STUB: not implemented"
	return *new(subsetGlyphTables), nil
}

func locaOffset(info subsetGlyphTables, index int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func glyphRange(info subsetGlyphTables, gid int) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func glyphData(info subsetGlyphTables, gid int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeGlyfOffset(off uint64, indexToLocFormat int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func writeGlyfOffset(buf *bytes.Buffer, off uint64, indexToLocFormat int) error {
	_ = "STUB: not implemented"
	return nil
}

func nextGlyfOffset(off uint64, glyphLength, indexToLocFormat int) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func pad(bb []byte) []byte { _ = "STUB: not implemented"; return nil }

func rebuiltTableData(tag string, bb []byte) ([]byte, uint32, uint32, error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, nil
}

func compoundTransformSize(flags uint16) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func compoundComponent(bb []byte, off int) (uint16, uint16, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, nil
}

func validateCompoundInstructions(bb []byte, off int, flags uint16) error {
	_ = "STUB: not implemented"
	return nil
}

func compoundGlyph(bb []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func resolveCompoundGlyph(gid int, bb []byte, usedGIDs map[uint16]bool, info subsetGlyphTables, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func sortedUsedGlyphIDs(usedGIDs map[uint16]bool) []int { _ = "STUB: not implemented"; return nil }

func resolveCompoundGlyphs(usedGIDs map[uint16]bool, info subsetGlyphTables) error {
	_ = "STUB: not implemented"
	return nil
}

func glyfAndLoca(tables map[string]*table, usedGIDs map[uint16]bool) error {
	_ = "STUB: not implemented"
	return nil
}

type ttfOutputTable struct {
	tag      string
	table    *table
	checksum uint32
	offset   uint32
}

func invalidFontData(format string, args ...any) error { _ = "STUB: not implemented"; return nil }

func validateSFNTHeader(header []byte) error { _ = "STUB: not implemented"; return nil }

func validateTTFTable(tag string, t *table) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func nextTTFOffset(offset uint64, size uint32) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func prepareTTFOutput(header []byte, tables map[string]*table) ([]ttfOutputTable, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func createTTF(header []byte, tables map[string]*table) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Subset(fontName string, usedGIDs map[uint16]bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
