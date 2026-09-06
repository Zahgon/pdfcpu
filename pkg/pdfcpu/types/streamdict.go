package types

import (
	"context"
	"io"
)

type PDFFilter struct {
	Name        string
	DecodeParms Dict
}

type StreamDict struct {
	Dict
	StreamOffset      int64
	StreamLength      *int64
	StreamLengthObjNr *int
	FilterPipeline    []PDFFilter
	Raw               []byte
	Content           []byte

	IsPageContent bool
	CSComponents  int
}

func NewStreamDict(d Dict, streamOffset int64, streamLength *int64, streamLengthObjNr *int, filterPipeline []PDFFilter) StreamDict {
	_ = "STUB: not implemented"
	return *new(StreamDict)
}

func (sd StreamDict) Clone() Object { _ = "STUB: not implemented"; return *new(Object) }

func (sd StreamDict) HasSoleFilterNamed(filterName string) bool {
	_ = "STUB: not implemented"
	return false
}

func (sd StreamDict) Image() bool { _ = "STUB: not implemented"; return false }

type DecodeLazyObjectStreamObjectFunc func(c context.Context, s string) (Object, error)

type LazyObjectStreamObject struct {
	osd         *ObjectStreamDict
	startOffset int
	endOffset   int

	decodeFunc    DecodeLazyObjectStreamObjectFunc
	decodedObject Object
	decodedError  error
}

func NewLazyObjectStreamObject(osd *ObjectStreamDict, startOffset, endOffset int, decodeFunc DecodeLazyObjectStreamObjectFunc) Object {
	_ = "STUB: not implemented"
	return *new(Object)
}

func (l LazyObjectStreamObject) Clone() Object { _ = "STUB: not implemented"; return *new(Object) }

func (l LazyObjectStreamObject) PDFString() string { _ = "STUB: not implemented"; return "" }

func (l LazyObjectStreamObject) String() string { _ = "STUB: not implemented"; return "" }

func (l *LazyObjectStreamObject) GetData() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *LazyObjectStreamObject) DecodedObject(c context.Context) (Object, error) {
	_ = "STUB: not implemented"
	return *new(Object), nil
}

type ObjectStreamDict struct {
	StreamDict
	Prolog         []byte
	ObjCount       int
	FirstObjOffset int
	MaxDecodeBytes int64
	ObjArray       Array
}

func NewObjectStreamDict() *ObjectStreamDict { _ = "STUB: not implemented"; return nil }

func parmsForFilter(d Dict) map[string]int { _ = "STUB: not implemented"; return nil }

func (sd *StreamDict) Encode() error { _ = "STUB: not implemented"; return nil }

func fixParms(f PDFFilter, parms map[string]int, sd *StreamDict) error {
	_ = "STUB: not implemented"
	return nil
}

func preserveEncodedImageFilter(name string) bool { _ = "STUB: not implemented"; return false }

func decodedContent(r io.Reader) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (sd *StreamDict) Decode() error { _ = "STUB: not implemented"; return nil }

func (sd *StreamDict) DecodeWithLimit(maxDecodeBytes int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (sd *StreamDict) decodeLength(maxLen, maxDecodeBytes int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sd *StreamDict) DecodeLength(maxLen int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sd *StreamDict) DecodeLengthWithLimit(maxLen, maxDecodeBytes int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (osd *ObjectStreamDict) IndexedObject(index int) (Object, error) {
	_ = "STUB: not implemented"
	return *new(Object), nil
}

func (osd *ObjectStreamDict) AddObject(objNumber int, pdfString string) error {
	_ = "STUB: not implemented"
	return nil
}

func (osd *ObjectStreamDict) Finalize() { _ = "STUB: not implemented"; return }

type XRefStreamDict struct {
	StreamDict
	Size           int
	Objects        []int
	W              [3]int
	PreviousOffset *int64
}
