package types

func RemoveControlChars(s string) string { _ = "STUB: not implemented"; return "" }

func NewStringSet(slice []string) StringSet { _ = "STUB: not implemented"; return *new(StringSet) }

func ByteForOctalString(octalBytes string) byte { _ = "STUB: not implemented"; return 0 }

func Escape(s string) (*string, error) { _ = "STUB: not implemented"; return nil, nil }

func escaped(c byte) (bool, byte) { _ = "STUB: not implemented"; return false, 0 }

func regularChar(c byte, esc bool) bool { _ = "STUB: not implemented"; return false }

func Unescape(s string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func UTF8ToCP1252(s string) string { _ = "STUB: not implemented"; return "" }

func CP1252ToUTF8(s string) string { _ = "STUB: not implemented"; return "" }

func Reverse(s string) string { _ = "STUB: not implemented"; return "" }

func needsHexSequence(c byte) bool { _ = "STUB: not implemented"; return false }

func EncodeName(s string) string { _ = "STUB: not implemented"; return "" }

func DecodeName(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func TrimLeadingComment(s string) string { _ = "STUB: not implemented"; return "" }
