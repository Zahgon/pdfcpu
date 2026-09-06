package model

import (
	"bufio"
	"errors"
	"io"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

var ErrMissingPDFContext = errors.New("missing PDF context")

var ErrMissingXRefTable = errors.New("missing PDF cross-reference table")

type Context struct {
	*Configuration
	*XRefTable
	Read         *ReadContext
	Optimize     *OptimizationContext
	Write        *WriteContext
	WritingPages bool
	Dest         bool
}

func NewContext(rs io.ReadSeeker, conf *Configuration) (*Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctx *Context) ResetWriteContext() { _ = "STUB: not implemented"; return }

func (rc *ReadContext) logReadContext(logStr *[]string) { _ = "STUB: not implemented"; return }

func (ctx *Context) String() string { _ = "STUB: not implemented"; return "" }

func (ctx *Context) UnitString() string { _ = "STUB: not implemented"; return "" }

func (ctx *Context) ConvertToUnit(d types.Dim) types.Dim {
	_ = "STUB: not implemented"
	return *new(types.Dim)
}

type ReadContext struct {
	FileName            string
	FileSize            int64
	RS                  io.ReadSeeker
	EolCount            int
	RepairOffset        int64
	BinaryTotalSize     int64
	BinaryImageSize     int64
	BinaryFontSize      int64
	BinaryImageDuplSize int64
	BinaryFontDuplSize  int64
	Linearized          bool
	Hybrid              bool
	UsingObjectStreams  bool
	ObjectStreams       types.IntSet
	UsingXRefStreams    bool
	XRefStreams         types.IntSet
}

func newReadContext(rs io.ReadSeeker) (*ReadContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rc *ReadContext) IsObjectStreamObject(i int) bool { _ = "STUB: not implemented"; return false }

func (rc *ReadContext) ObjectStreamsString() (int, string) { _ = "STUB: not implemented"; return 0, "" }

func (rc *ReadContext) LogStats(optimized bool) { _ = "STUB: not implemented"; return }

func (rc *ReadContext) ReadFileSize() int { _ = "STUB: not implemented"; return 0 }

type OptimizationContext struct {
	PageFonts           []types.IntSet
	FontObjects         map[int]*FontObject
	FormFontObjects     map[int]*FontObject
	Fonts               map[string][]int
	DuplicateFonts      map[int]types.Dict
	DuplicateFontObjs   types.IntSet
	CorruptFontResDicts []types.Dict

	PageImages         []types.IntSet
	ImageObjects       map[int]*ImageObject
	ImageObjectHashes  map[[32]byte][]int
	DuplicateImages    map[int]*DuplicateImageObject
	DuplicateImageObjs types.IntSet

	ContentStreamCache map[int]*types.StreamDict

	FormStreamCache map[int]*types.StreamDict

	FormResourceCache map[int]types.IntSet

	DuplicateInfoObjects types.IntSet
	NonReferencedObjs    []int

	Cache     map[int]bool
	NullObjNr *int
}

func newOptimizationContext() *OptimizationContext { _ = "STUB: not implemented"; return nil }

func (oc *OptimizationContext) IsDuplicateFontObject(i int) bool {
	_ = "STUB: not implemented"
	return false
}

func (oc *OptimizationContext) DuplicateFontObjectsString() (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

func (oc *OptimizationContext) IsDuplicateImageObject(i int) bool {
	_ = "STUB: not implemented"
	return false
}

func (oc *OptimizationContext) DuplicateImageObjectsString() (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

func (oc *OptimizationContext) IsDuplicateInfoObject(i int) bool {
	_ = "STUB: not implemented"
	return false
}

func (oc *OptimizationContext) DuplicateInfoObjectsString() (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

func (oc *OptimizationContext) NonReferencedObjsString() (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

func (oc *OptimizationContext) collectFontInfo(logStr []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (oc *OptimizationContext) collectImageInfo(logStr []string) []string {
	_ = "STUB: not implemented"
	return nil
}

type WriteContext struct {
	*bufio.Writer
	Fp                  *os.File
	FileSize            int64
	DirName             string
	FileName            string
	SelectedPages       types.IntSet
	BinaryTotalSize     int64
	BinaryImageSize     int64
	BinaryFontSize      int64
	Table               map[int]int64
	Offset              int64
	OffsetSigByteRange  int64
	OffsetSigContents   int64
	WriteToObjectStream bool
	CurrentObjStream    *int
	Eol                 string
	Increment           bool
	ObjNrs              []int
	OffsetPrevXRef      *int64
}

func NewWriteContext(eol string) *WriteContext { _ = "STUB: not implemented"; return nil }

func (wc *WriteContext) SetWriteOffset(objNumber int) { _ = "STUB: not implemented"; return }

func (wc *WriteContext) HasWriteOffset(objNumber int) bool { _ = "STUB: not implemented"; return false }

func (wc *WriteContext) LogStats() { _ = "STUB: not implemented"; return }

func (wc *WriteContext) WriteEol() error { _ = "STUB: not implemented"; return nil }

func (wc *WriteContext) IncrementWithObjNr(i int) { _ = "STUB: not implemented"; return }
