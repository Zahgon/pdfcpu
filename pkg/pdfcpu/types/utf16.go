package types

import (
	"errors"
)

var (
	ErrInvalidUTF8 = errors.New("invalid UTF-8")

	ErrInvalidUTF16BE = errors.New("invalid UTF-16BE detected")
)

func IsStringUTF16BE(s string) bool { _ = "STUB: not implemented"; return false }

func IsUTF16BE(b []byte) bool { _ = "STUB: not implemented"; return false }

func decodeUTF16String(b []byte) (string, error) { _ = "STUB: not implemented"; return "", nil }

func DecodeUTF16String(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func EncodeUTF16String(s string) string { _ = "STUB: not implemented"; return "" }

func EscapedUTF16String(s string) (*string, error) { _ = "STUB: not implemented"; return nil, nil }

func StringLiteralToString(sl StringLiteral) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func HexLiteralToString(hl HexLiteral) (string, error) { _ = "STUB: not implemented"; return "", nil }

func StringOrHexLiteral(obj Object) (*string, error) { _ = "STUB: not implemented"; return nil, nil }
