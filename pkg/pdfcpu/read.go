package pdfcpu

import (
	"bufio"
	"bytes"
	"context"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"github.com/pkg/errors"
)

const (
	defaultBufSize = 1 << 10
	maxBufSize     = 1 << 20
)

var (
	ErrCorruptHeader         = errors.New("pdfcpu: no header version available")
	ErrMissingXRefSection    = errors.New("pdfcpu: can't detect last xref section")
	ErrReferenceDoesNotExist = errors.New("pdfcpu: referenced object does not exist")
	ErrWrongPassword         = errors.New("pdfcpu: please provide the correct password")

	zero int64 = 0
)

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

func parseAndLoad(c context.Context, ctx *model.Context, line string, offset *int64, incr int, offsetPrev *int64) error {
	_ = "STUB: not implemented"
	return nil
}

func processObject(c context.Context, ctx *model.Context, line string, offset *int64, incr int, offsetPrev *int64) (*bufio.Scanner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func objCandidate(withinObj bool, line string) bool { _ = "STUB: not implemented"; return false }

func checkEndObj(withinObj *bool, line *string) { _ = "STUB: not implemented"; return }

func ensureNoStartXRef(line string, i int) bool { _ = "STUB: not implemented"; return false }

func bypassXrefSection(c context.Context, ctx *model.Context, offExtra int64, wasErr error, incr int) error {
	_ = "STUB: not implemented"
	return nil
}

func postProcess(ctx *model.Context, xrefSectionCount int) { _ = "STUB: not implemented"; return }

func tryXRefSection(c context.Context, ctx *model.Context, rs io.ReadSeeker, offset *int64, offExtra int64, xrefSectionCount *int, incr int) (*int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildXRefTableStartingAt(c context.Context, ctx *model.Context, offset *int64) error {
	_ = "STUB: not implemented"
	return nil
}

func readXRefTable(c context.Context, ctx *model.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func growBufBy(buf []byte, size int, rd io.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func nextStreamOffset(line string, streamInd int) (off int) { _ = "STUB: not implemented"; return 0 }

func lastStreamMarker(streamInd *int, endInd int, line string) { _ = "STUB: not implemented"; return }

func buffer(c context.Context, rd io.Reader) (buf []byte, endInd int, streamInd int, streamOffset int64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, 0, nil
}

func keywordStreamRightAfterEndOfDict(buf string, streamInd int) bool {
	_ = "STUB: not implemented"
	return false
}

func buildFilterPipeline(c context.Context, ctx *model.Context, filterArray, decodeParmsArr types.Array) ([]types.PDFFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func singleFilter(c context.Context, ctx *model.Context, filterName string, d types.Dict) ([]types.PDFFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func filterArraySupportsDecodeParms(filters types.Array) bool {
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
