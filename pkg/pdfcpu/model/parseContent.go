package model

import (
	"errors"
)

var (
	errPageContentCorrupt  = errors.New("corrupt page content")
	errTJExpressionCorrupt = errors.New("corrupt TJ expression")
	errBIExpressionCorrupt = errors.New("corrupt BI expression")
)

func whitespaceOrEOL(c rune) bool { _ = "STUB: not implemented"; return false }

func whitespaceOrEOLOrClosingBracket(c rune) bool { _ = "STUB: not implemented"; return false }

func skipDict(l *string) error { _ = "STUB: not implemented"; return nil }

func skipStringLiteral(l *string) error { _ = "STUB: not implemented"; return nil }

func skipHexStringLiteral(l *string) error { _ = "STUB: not implemented"; return nil }

func skipTJ(l *string) error { _ = "STUB: not implemented"; return nil }

func lookupEI(l *string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func skipBIColorSpaceName(s string, prn PageResourceNames) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func skipBI(l *string, prn PageResourceNames) error { _ = "STUB: not implemented"; return nil }

func skipInlineImage(l *string, prn PageResourceNames) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func positionToNextContentToken(line *string, prn PageResourceNames) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func nextContentToken(pre string, line *string, prn PageResourceNames) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func colorSpace(s, name string, prn PageResourceNames) bool {
	_ = "STUB: not implemented"
	return false
}

func resourceNameAtPos1(s, name string, prn PageResourceNames) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func resourceNameAtPos2(s, name string, prn PageResourceNames) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func parseContent(s string) (PageResourceNames, error) {
	_ = "STUB: not implemented"
	return *new(PageResourceNames), nil
}
