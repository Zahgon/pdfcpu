package model

import (
	"context"

	"github.com/pkg/errors"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

var (
	errArrayCorrupt            = errors.New("pdfcpu: parse: corrupt array")
	errArrayNotTerminated      = errors.New("pdfcpu: parse: unterminated array")
	errDictionaryCorrupt       = errors.New("pdfcpu: parse: corrupt dictionary")
	errDictionaryNotTerminated = errors.New("pdfcpu: parse: unterminated dictionary")
	errDictionaryDuplicateKey  = errors.New("pdfcpu: parse: duplicate key")
	errHexLiteralCorrupt       = errors.New("pdfcpu: parse: corrupt hex literal")
	errHexLiteralNotTerminated = errors.New("pdfcpu: parse: hex literal not terminated")
	errNameObjectCorrupt       = errors.New("pdfcpu: parse: corrupt name object")
	errNoArray                 = errors.New("pdfcpu: parse: no array")
	errNoDictionary            = errors.New("pdfcpu: parse: no dictionary")
	errStringLiteralCorrupt    = errors.New("pdfcpu: parse: corrupt string literal, possibly unbalanced parenthesis")
	errBufNotAvailable         = errors.New("pdfcpu: parse: no buffer available")
	errXrefStreamMissingW      = errors.New("pdfcpu: parse: xref stream dict missing entry W")
	errXrefStreamCorruptW      = errors.New("pdfcpu: parse: xref stream dict corrupt entry W: expecting array of 3 int")
	errXrefStreamCorruptIndex  = errors.New("pdfcpu: parse: xref stream dict corrupt entry Index")
	errObjStreamMissingN       = errors.New("pdfcpu: parse: obj stream dict missing entry W")
	errObjStreamMissingFirst   = errors.New("pdfcpu: parse: obj stream dict missing entry First")
	ErrCorruptObjectOffset     = errors.New("pdfcpu: corrupt object offset")
)

func positionToNextWhitespace(s string) (int, string) { _ = "STUB: not implemented"; return 0, "" }

func positionToNextWhitespaceOrChar(s, chars string) (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

func positionToNextEOL(s string) (string, int) { _ = "STUB: not implemented"; return "", 0 }

func trimLeftSpace(s string, relaxed bool) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func hexString(s string) (*string, bool) { _ = "STUB: not implemented"; return nil, false }

func balancedParenthesesPrefix(s string) int { _ = "STUB: not implemented"; return 0 }

func forwardParseBuf(buf string, pos int) string { _ = "STUB: not implemented"; return "" }

func delimiter(b byte) bool { _ = "STUB: not implemented"; return false }

func detectObj(s string) (string, string, error) { _ = "STUB: not implemented"; return "", "", nil }

func cleanObjProlog(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func ParseObjectAttributes(line *string) (*int, *int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func parseArray(c context.Context, line *string, level, maxDepth int) (*types.Array, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseStringLiteral(line *string) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func parseHexLiteral(line *string) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func decodeNameHexSequence(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func parseName(line *string) (*types.Name, error) { _ = "STUB: not implemented"; return nil, nil }

func insertKey(d types.Dict, key string, val types.Object, relaxed bool) error {
	_ = "STUB: not implemented"
	return nil
}

func dictString(l string) bool { _ = "STUB: not implemented"; return false }

func processDictKeys(c context.Context, line *string, level, maxDepth int, relaxed bool) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func parseDict(c context.Context, line *string, level, maxDepth int, relaxed bool) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func noBuf(l *string) bool { _ = "STUB: not implemented"; return false }

func startParseNumericOrIndRef(l string) (string, string, int) {
	_ = "STUB: not implemented"
	return "", "", 0
}

func isRangeError(err error) bool { _ = "STUB: not implemented"; return false }

func parseIndRef(s, l, l1 string, line *string, i, i2 int) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func parseFloat(s string) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func parseNumericOrIndRef(line *string) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func parseHexLiteralOrDict(c context.Context, l *string, level, maxDepth int) (val types.Object, err error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func parseBooleanOrNull(l string) (types.Object, string, bool) {
	_ = "STUB: not implemented"
	return *new(types.Object), "", false
}

func ParseObject(line *string) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func parseObjectDepthLimit(maxDepth []int) int { _ = "STUB: not implemented"; return 0 }

func parseObjectValue(c context.Context, l *string, level, depthLimit int) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func ParseObjectContext(c context.Context, line *string, level int, maxDepth ...int) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func createXRefStreamDict(sd *types.StreamDict, objs []int) (*types.XRefStreamDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseXRefStreamDict(sd *types.StreamDict) (*types.XRefStreamDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func xRefStreamSize(sd *types.StreamDict, limits ResourceLimits) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func xRefStreamObjectsFromIndex(indArr types.Array, size int, limits ResourceLimits) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func xRefStreamObjectsFromSize(size int, limits ResourceLimits) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func xRefStreamObjects(sd *types.StreamDict, size int, limits ResourceLimits) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseXRefStreamDictWithLimits(sd *types.StreamDict, limits ResourceLimits) (*types.XRefStreamDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ObjectStreamDict(sd *types.StreamDict) (*types.ObjectStreamDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ObjectStreamDictWithLimits(sd *types.StreamDict, limits ResourceLimits) (*types.ObjectStreamDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isMarkerTerminated(r rune) bool { _ = "STUB: not implemented"; return false }

func detectMarker(line, marker string) int { _ = "STUB: not implemented"; return 0 }

func detectMarkers(line string, endInd, streamInd *int) { _ = "STUB: not implemented"; return }

func positionAfterStringLiteral(line string) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func posFloor(pos1, pos2 int) int { _ = "STUB: not implemented"; return 0 }

func detectNonEscaped(line, s string) int { _ = "STUB: not implemented"; return 0 }

func applyOffBoth(endInd, streamInd, off int) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func applyOffEndIndFirst(endInd, streamInd, off, floor int) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func applyOffStreamIndFirst(endInd, streamInd, off, floor int) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func isComment(commentPos, strLitPos int) bool { _ = "STUB: not implemented"; return false }

func DetectKeywords(line string) (endInd int, streamInd int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func skipComment(line string, commentPos int, off, endInd, streamInd *int) string {
	_ = "STUB: not implemented"
	return ""
}

func skipStringLit(line string, strLitPos int, off, endInd, streamInd *int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func skipCommentOrStringLiteral(line string, commentPos, slPos int, off, endInd, streamInd *int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func DetectKeywordsWithContext(c context.Context, line string) (endInd int, streamInd int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}
