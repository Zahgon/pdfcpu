package types

type Dict map[string]Object

func NewDict() Dict { _ = "STUB: not implemented"; return *new(Dict) }

func (d Dict) Len() int { _ = "STUB: not implemented"; return 0 }

func (d Dict) Clone() Object { _ = "STUB: not implemented"; return *new(Object) }

func (d Dict) Insert(k string, v Object) bool { _ = "STUB: not implemented"; return false }

func (d Dict) InsertBool(key string, value bool) { _ = "STUB: not implemented"; return }

func (d Dict) InsertInt(key string, value int) { _ = "STUB: not implemented"; return }

func (d Dict) InsertFloat(key string, value float32) { _ = "STUB: not implemented"; return }

func (d Dict) InsertString(key, value string) { _ = "STUB: not implemented"; return }

func (d Dict) InsertName(key, value string) { _ = "STUB: not implemented"; return }

func (d Dict) Update(key string, value Object) { _ = "STUB: not implemented"; return }

func (d Dict) Find(key string) (Object, bool) {
	_ = "STUB: not implemented"
	return *new(Object), false
}

func (d Dict) HasEntry(key string) bool { _ = "STUB: not implemented"; return false }

func (d Dict) Delete(key string) (value Object) { _ = "STUB: not implemented"; return *new(Object) }

func (d Dict) NewIDForPrefix(prefix string, i int) string { _ = "STUB: not implemented"; return "" }

func (d Dict) Entry(dictName, key string, required bool) (Object, bool, error) {
	_ = "STUB: not implemented"
	return *new(Object), false, nil
}

func (d Dict) BooleanEntry(key string) *bool { _ = "STUB: not implemented"; return nil }

func (d Dict) StringEntry(key string) *string { _ = "STUB: not implemented"; return nil }

func (d Dict) NameEntry(key string) *string { _ = "STUB: not implemented"; return nil }

func (d Dict) IntEntry(key string) *int { _ = "STUB: not implemented"; return nil }

func (d Dict) Int64Entry(key string) *int64 { _ = "STUB: not implemented"; return nil }

func (d Dict) IndirectRefEntry(key string) *IndirectRef { _ = "STUB: not implemented"; return nil }

func (d Dict) DictEntry(key string) Dict { _ = "STUB: not implemented"; return *new(Dict) }

func (d Dict) StreamDictEntry(key string) *StreamDict { _ = "STUB: not implemented"; return nil }

func (d Dict) ArrayEntry(key string) Array { _ = "STUB: not implemented"; return *new(Array) }

func (d Dict) StringLiteralEntry(key string) *StringLiteral { _ = "STUB: not implemented"; return nil }

func (d Dict) HexLiteralEntry(key string) *HexLiteral { _ = "STUB: not implemented"; return nil }

func (d Dict) StringOrHexLiteralEntry(key string) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d Dict) Length() (*int64, *int) { _ = "STUB: not implemented"; return nil, nil }

func (d Dict) Type() *string { _ = "STUB: not implemented"; return nil }

func (d Dict) Subtype() *string { _ = "STUB: not implemented"; return nil }

func (d Dict) Size() *int { _ = "STUB: not implemented"; return nil }

func (d Dict) IsPage() bool { _ = "STUB: not implemented"; return false }

func (d Dict) IsObjStm() bool { _ = "STUB: not implemented"; return false }

func (d Dict) W() Array { _ = "STUB: not implemented"; return *new(Array) }

func (d Dict) Prev() *int64 { _ = "STUB: not implemented"; return nil }

func (d Dict) Index() Array { _ = "STUB: not implemented"; return *new(Array) }

func (d Dict) N() *int { _ = "STUB: not implemented"; return nil }

func (d Dict) First() *int { _ = "STUB: not implemented"; return nil }

func (d Dict) IsLinearizationParmDict() bool { _ = "STUB: not implemented"; return false }

func (d *Dict) IncrementBy(key string, i int) error { _ = "STUB: not implemented"; return nil }

func (d *Dict) Increment(key string) error { _ = "STUB: not implemented"; return nil }

func (d Dict) indentedString(level int) string { _ = "STUB: not implemented"; return "" }

func (d Dict) PDFString() string { _ = "STUB: not implemented"; return "" }

func (d Dict) String() string { _ = "STUB: not implemented"; return "" }

func (d Dict) StringEntryBytes(key string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
