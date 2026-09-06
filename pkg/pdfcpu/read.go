package pdfcpu

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/filter"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

const (
	defaultBufSize     = 1 << 10
	maxBufSize         = 1 << 20
	maxObjectBufferLen = 64 << 20
)

var (
	ErrCorruptHeader = errors.New("no header version available")

	ErrPostScriptInput = errors.New("PostScript input is not supported; convert to PDF before processing")

	ErrEmptyInput = errors.New("the file could not be opened because it is empty")

	ErrMissingXRefSection = errors.New("can't detect last xref section")

	ErrReferenceDoesNotExist = errors.New("referenced object does not exist")

	ErrWrongPassword = errors.New("please provide the correct password")

	ErrNotEncrypted = errors.New("this file is not encrypted")

	ErrEncrypted = errors.New("this file is encrypted")

	ErrOwnerPasswordRequired = errors.New("please provide owner password and optional user password")

	ErrPermissionDenied = errors.New("operation restricted via permission bits setting")

	errMissingReadSeeker            = errors.New("missing PDF read seeker")
	errMissingXRefEOF               = errors.New("no matching %%EOF for startxref")
	errInvalidLastXRefSection       = errors.New("invalid last xref section")
	errInvalidXRefStreamWArray      = errors.New("invalid xref stream W array")
	errCorruptXRefStream            = errors.New("corrupt xref stream")
	errCorruptXRefSubsection        = errors.New("corrupt xref subsection")
	errIncompleteXRefSubsection     = errors.New("incomplete xref subsection")
	errMissingScannerLine           = errors.New("missing scanner line")
	errMissingTrailerSize           = errors.New("missing trailer Size")
	errMissingTrailerRoot           = errors.New("missing trailer Root")
	errInvalidTrailerID             = errors.New("invalid trailer ID")
	errMissingTrailerID             = errors.New("missing trailer ID")
	errCorruptTrailerDict           = errors.New("corrupt trailer dict")
	errMissingXRefStreamDict        = errors.New("missing xref stream dictionary")
	errMissingXRefStreamLength      = errors.New("missing xref stream Length")
	errCorruptStreamMarker          = errors.New("corrupt stream marker")
	errObjectBufferLimit            = errors.New("object buffer limit exceeded")
	errTruncatedStreamMarker        = errors.New("truncated stream marker")
	errUnregisteredObject           = errors.New("unregistered object")
	errNilDereferencedObject        = errors.New("nil dereferenced object")
	errCorruptIntegerObject         = errors.New("corrupt integer object")
	errCorruptDictObject            = errors.New("corrupt dict object")
	errMissingObjectStreamEntry     = errors.New("missing object stream xref entry")
	errMissingObjectStream          = errors.New("missing object stream")
	errCorruptStreamDict            = errors.New("corrupt stream dict")
	errMissingStreamLength          = errors.New("missing stream length")
	errMissingStreamOffset          = errors.New("missing stream offset")
	errCorruptFilterArray           = errors.New("corrupt filter array")
	errCorruptDecodeParms           = errors.New("corrupt decode parameters")
	errObjectStreamContainsStream   = errors.New("object stream contains stream object")
	errCorruptObjectStreamDict      = errors.New("corrupt object stream dict")
	errCorruptObjectStream          = errors.New("corrupt object stream")
	errMissingObjectStreamObjects   = errors.New("missing object stream objects")
	errObjectStreamObjectCountLimit = errors.New("object stream object count limit exceeded")
	errInvalidPermissions           = errors.New("invalid permissions")

	zero int64 = 0
)

func hasPostScriptHeader(buf []byte) bool { _ = "STUB: not implemented"; return false }

func ReadFile(inFile string, conf *model.Configuration) (*model.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReadFileWithContext(c context.Context, inFile string, conf *model.Configuration) (*model.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Read(rs io.ReadSeeker, conf *model.Configuration) (*model.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReadWithContext(c context.Context, rs io.ReadSeeker, conf *model.Configuration) (*model.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fillBuffer(r io.Reader, buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func newPositionedReader(rs io.ReadSeeker, offset *int64) (*bufio.Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func incrEpilogIndex(s string) int { _ = "STUB: not implemented"; return 0 }

func offsetLastXRefSection(ctx *model.Context, skip int64) (*int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createXRefTableEntry(entryType string, objNr int, offset, offExtra int64, generation, incr int) (model.XRefTableEntry, bool) {
	_ = "STUB: not implemented"
	return *new(model.XRefTableEntry), false
}

func decodeSubsection(fields []string) (int64, int, string, error) {
	_ = "STUB: not implemented"
	return 0, 0, "", nil
}

func parseXRefTableEntry(xRefTable *model.XRefTable, fields []string, objNr int, offExtra int64, incr int) error {
	_ = "STUB: not implemented"
	return nil
}

func parseXRefTableSubSection(xRefTable *model.XRefTable, s *bufio.Scanner, fields []string, offExtra int64, incr int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func compressedObject(c context.Context, s string) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func buildObjectArrayForObjectStream(
	c context.Context,
	osd *types.ObjectStreamDict,
	objs []string,
	fullContent bool,
	decodedContent []byte) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func parseObjectStream(c context.Context, osd *types.ObjectStreamDict, limits model.ResourceLimits) error {
	_ = "STUB: not implemented"
	return nil
}

func createXRefTableEntryFromXRefStream(entryType int64, objNr int, c2, c3, offExtra int64, objStreams types.IntSet, incr int) model.XRefTableEntry {
	_ = "STUB: not implemented"
	return *new(model.XRefTableEntry)
}

func extractXRefTableEntriesFromXRefStream(buf []byte, offExtra int64, xsd *types.XRefStreamDict, ctx *model.Context, incr int) error {
	_ = "STUB: not implemented"
	return nil
}

func xRefStreamDict(c context.Context, ctx *model.Context, o types.Object, objNr int, streamOffset int64) (*types.XRefStreamDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func processXRefStream(ctx *model.Context, xsd *types.XRefStreamDict, objNr *int, offset *int64, offExtra int64, incr int) (prevOffset *int64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseXRefStream(c context.Context, ctx *model.Context, rd io.Reader, offset *int64, offExtra int64, incr int) (prevOffset *int64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseHybridXRefStream(c context.Context, ctx *model.Context, offset *int64, offExtra int64, incr int) error {
	_ = "STUB: not implemented"
	return nil
}

func parseTrailerSize(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func parseTrailerRoot(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func parseTrailerInfo(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func parseTrailerID(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func parseTrailer(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func scanForPreviousXref(ctx *model.Context, offset *int64) *int64 {
	_ = "STUB: not implemented"
	return nil
}

func handleAdditionalStreams(trailerDict types.Dict, xRefTable *model.XRefTable) {
	_ = "STUB: not implemented"
	return
}

func offsetPrev(ctx *model.Context, trailerDict types.Dict, offCurXRef *int64) *int64 {
	_ = "STUB: not implemented"
	return nil
}

func parseTrailerDict(c context.Context, ctx *model.Context, trailerDict types.Dict, offCurXRef *int64, offExtra int64, incr int, repairing bool) (*int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func scanLineRaw(s *bufio.Scanner) (string, error) { _ = "STUB: not implemented"; return "", nil }

func scanLine(s *bufio.Scanner) (s1 string, err error) { _ = "STUB: not implemented"; return "", nil }

func scanTrailerDictStart(s *bufio.Scanner, line *string) error {
	_ = "STUB: not implemented"
	return nil
}

func scanTrailerDictRemainder(s *bufio.Scanner, line string, buf bytes.Buffer) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func scanTrailer(s *bufio.Scanner, line string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func processTrailer(c context.Context, ctx *model.Context, s *bufio.Scanner, line string, offCurXRef *int64, offExtra int64, incr int, repairing bool) (*int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseXRefSection(c context.Context, ctx *model.Context, s *bufio.Scanner, fields []string, ssCount *int, offCurXRef *int64, offExtra int64, incr int) (*int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func scanForVersion(rs io.ReadSeeker, prefix string) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func headerVersion(rs io.ReadSeeker) (v *model.Version, eolCount int, offset int64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, nil
}

func parseAndLoad(c context.Context, ctx *model.Context, line string, offset *int64, incr int, offsetPrev *int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func processObject(c context.Context, ctx *model.Context, line string, offset *int64, incr int, offsetPrev *int64) (*bufio.Scanner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func objCandidate(withinObj bool, line string) bool { _ = "STUB: not implemented"; return false }

func checkEndObj(withinObj *bool, line *string) { _ = "STUB: not implemented"; return }

func ensureNoStartXRef(line string, i int) bool { _ = "STUB: not implemented"; return false }

func registerXRefEntry(ctx *model.Context, line string, offset int64, incr int) bool {
	_ = "STUB: not implemented"
	return false
}

func registerXRefObjects(ctx *model.Context, incr int) error { _ = "STUB: not implemented"; return nil }

type xrefRepairState struct {
	scanner       *bufio.Scanner
	offset        int64
	withinObj     bool
	withinXref    bool
	withinTrailer bool
	bb            []byte
	offsetPrev    *int64
}

func processXRefRepairLine(c context.Context, ctx *model.Context, state *xrefRepairState, line string, length, eolCount int, offExtra int64, incr int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func bypassXrefSection(c context.Context, ctx *model.Context, offExtra int64, wasErr error, incr int) error {
	_ = "STUB: not implemented"
	return nil
}

func postProcess(ctx *model.Context, xrefSectionCount int) { _ = "STUB: not implemented"; return }

func tryXRefSection(c context.Context, ctx *model.Context, rs io.ReadSeeker, offset *int64, offExtra int64, xrefSectionCount *int, incr int) (*int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseXRefStreamOrRepair(c context.Context, ctx *model.Context, rs io.ReadSeeker, offset *int64, offExtra int64, incr int) (*int64, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func buildXRefTableStartingAt(c context.Context, ctx *model.Context, offset *int64) error {
	_ = "STUB: not implemented"
	return nil
}

func readXRefTable(c context.Context, ctx *model.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func isMissingXRefSection(err error) bool { _ = "STUB: not implemented"; return false }

func repairableLastXRefSection(ctx *model.Context, err error) bool {
	_ = "STUB: not implemented"
	return false
}

func growBufBy(buf []byte, size int, rd io.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func nextStreamOffset(line string, streamInd int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func lastStreamMarker(streamInd *int, endInd int, line string) { _ = "STUB: not implemented"; return }

func growObjectBuffer(buf []byte, growSize int, rd io.Reader, maxObjectBytes int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func streamOffsetForBuffer(buf []byte, line string, streamInd int, rd io.Reader, maxObjectBytes int64) ([]byte, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func buffer(c context.Context, rd io.Reader, maxObjectBytes int64) (buf []byte, endInd int, streamInd int, streamOffset int64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, 0, nil
}

func keywordStreamRightAfterEndOfDict(buf string, streamInd int) bool {
	_ = "STUB: not implemented"
	return false
}

var inlineImageFilterAliases = map[string]string{
	"AHx": filter.ASCIIHex,
	"A85": filter.ASCII85,
	"LZW": filter.LZW,
	"Fl":  filter.Flate,
	"RL":  filter.RunLength,
	"CCF": filter.CCITTFax,
	"DCT": filter.DCT,
}

func streamFilterName(ctx *model.Context, name string) string { _ = "STUB: not implemented"; return "" }

func normalizeStreamFilterName(ctx *model.Context, dict types.Dict, name string) string {
	_ = "STUB: not implemented"
	return ""
}

func normalizeStreamFilterArray(ctx *model.Context, dict types.Dict, filters types.Array) types.Array {
	_ = "STUB: not implemented"
	return *new(types.Array)
}

func buildFilterPipeline(c context.Context, ctx *model.Context, filterArray, decodeParmsArr types.Array) ([]types.PDFFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func singleFilter(c context.Context, ctx *model.Context, filterName string, d types.Dict) ([]types.PDFFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func filterArraySupportsDecodeParms(ctx *model.Context, filters types.Array) bool {
	_ = "STUB: not implemented"
	return false
}

func pdfFilterPipeline(c context.Context, ctx *model.Context, dict types.Dict) ([]types.PDFFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func streamDictForObject(c context.Context, ctx *model.Context, d types.Dict, objNr, streamInd int, streamOffset, offset int64) (sd types.StreamDict, err error) {
	_ = "STUB: not implemented"
	return *new(types.StreamDict), nil
}

func dict(ctx *model.Context, d1 types.Dict, objNr, genNr, endInd, streamInd int) (d2 types.Dict, err error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func object(c context.Context, ctx *model.Context, offset int64, objNr, genNr int) (o types.Object, endInd, streamInd int, streamOffset int64, err error) {
	_ = "STUB: not implemented"
	return *new(types.Object), 0, 0, 0, nil
}

func resolveObject(c context.Context, ctx *model.Context, obj types.Object, offset int64, objNr, genNr, endInd, streamInd int, streamOffset int64) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func ParseObjectWithContext(c context.Context, ctx *model.Context, offset int64, objNr, genNr int) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func dereferencedObject(c context.Context, ctx *model.Context, objNr int) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func dereferencedInteger(c context.Context, ctx *model.Context, objNr int) (*types.Integer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dereferencedDict(c context.Context, ctx *model.Context, objNr int) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func int64Object(c context.Context, ctx *model.Context, objNr int) (*int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readStreamContentBlindly(rd io.Reader, maxStreamBytes int64) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readStreamContent(rd io.Reader, streamLength int, maxStreamBytes int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ensureStreamLength(sd *types.StreamDict, fixLength bool) error {
	_ = "STUB: not implemented"
	return nil
}

func decodeLimit(ctx *model.Context) int64 { _ = "STUB: not implemented"; return 0 }

func streamLimit(ctx *model.Context) int64 { _ = "STUB: not implemented"; return 0 }

func recursionLimit(ctx *model.Context) int { _ = "STUB: not implemented"; return 0 }

func ensureIndirectStreamLength(c context.Context, ctx *model.Context, sd *types.StreamDict, fixLength bool) error {
	_ = "STUB: not implemented"
	return nil
}

func loadEncodedStreamContent(c context.Context, ctx *model.Context, sd *types.StreamDict, fixLength bool) error {
	_ = "STUB: not implemented"
	return nil
}

func metadataStream(sd *types.StreamDict) bool { _ = "STUB: not implemented"; return false }

func unencryptedMetadata(ctx *model.Context, sd *types.StreamDict) bool {
	_ = "STUB: not implemented"
	return false
}

func repairablePlaintextMetadata(ctx *model.Context, sd *types.StreamDict, err error) bool {
	_ = "STUB: not implemented"
	return false
}

func validPlaintextMetadata(ctx *model.Context, sd *types.StreamDict) bool {
	_ = "STUB: not implemented"
	return false
}

func decryptStreamContent(ctx *model.Context, sd *types.StreamDict, objNr, genNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func saveDecodedStreamContent(ctx *model.Context, sd *types.StreamDict, objNr, genNr int, decode bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func decompressXRefTableEntry(xRefTable *model.XRefTable, objNr int, entry *model.XRefTableEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func logStream(o types.Object) { _ = "STUB: not implemented"; return }

func decodeObjectStreamObjects(c context.Context, ctx *model.Context, sd *types.StreamDict, objNr int) (*types.ObjectStreamDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decodeObjectStream(c context.Context, ctx *model.Context, objNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func decodeObjectStreams(c context.Context, ctx *model.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func handleLinearizationParmDict(ctx *model.Context, obj types.Object, objNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func loadStreamDict(c context.Context, ctx *model.Context, sd *types.StreamDict, objNr, genNr int, fixLength bool) error {
	_ = "STUB: not implemented"
	return nil
}

func updateBinaryTotalSize(ctx *model.Context, o types.Object) { _ = "STUB: not implemented"; return }

func dereferenceAndLoad(c context.Context, ctx *model.Context, objNr int, entry *model.XRefTableEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func dereferenceObject(c context.Context, ctx *model.Context, objNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func dereferenceObjectsSorted(c context.Context, ctx *model.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func dereferenceObjectsRaw(c context.Context, ctx *model.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func dereferenceObjects(c context.Context, ctx *model.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func identifyRootVersion(xRefTable *model.XRefTable) error { _ = "STUB: not implemented"; return nil }

func dereferenceXRefTable(c context.Context, ctx *model.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func handleUnencryptedFile(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func needsOwnerAndUserPassword(cmd model.CommandMode) bool { _ = "STUB: not implemented"; return false }

func handlePermissions(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

func setupEncryptionKey(ctx *model.Context, d types.Dict) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func checkForEncryption(c context.Context, ctx *model.Context) error {
	_ = "STUB: not implemented"
	return nil
}
