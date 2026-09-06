package types

type Array []Object

func NewStringLiteralArray(sVars ...string) Array { _ = "STUB: not implemented"; return *new(Array) }

func NewHexLiteralArray(sVars ...string) Array { _ = "STUB: not implemented"; return *new(Array) }

func NewNameArray(sVars ...string) Array { _ = "STUB: not implemented"; return *new(Array) }

func NewNumberArray(fVars ...float64) Array { _ = "STUB: not implemented"; return *new(Array) }

func NewIntegerArray(fVars ...int) Array { _ = "STUB: not implemented"; return *new(Array) }

func (a Array) Clone() Object { _ = "STUB: not implemented"; return *new(Object) }

func (a Array) indentedString(level int) string { _ = "STUB: not implemented"; return "" }

func (a Array) String() string { _ = "STUB: not implemented"; return "" }

func (a Array) PDFString() string { _ = "STUB: not implemented"; return "" }

func (a Array) RemoveNulls() Array { _ = "STUB: not implemented"; return *new(Array) }
