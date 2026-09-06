package model

import (
	"errors"
	"io"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

var ErrNoContent = errors.New("page without content")

var ErrPageNotFound = errors.New("page not found")

var ErrMissingEncryptDictObject = errors.New("missing encryption dictionary object")

var ErrWrongTypeEncryptDictObject = errors.New("wrong type for encryption dictionary object")

var zero int64 = 0

type XRefTableEntry struct {
	Free            bool
	Offset          *int64
	Generation      *int
	Incr            int
	RefCount        int
	Object          types.Object
	Compressed      bool
	ObjectStream    *int
	ObjectStreamInd *int
	Valid           bool
	BeingValidated  bool
}

func NewXRefTableEntryGen0(obj types.Object) *XRefTableEntry { _ = "STUB: not implemented"; return nil }

func NewFreeHeadXRefTableEntry() *XRefTableEntry { _ = "STUB: not implemented"; return nil }

type Enc struct {
	O, U       []byte
	OE, UE     []byte
	Perms      []byte
	L, P, R, V int
	Emd        bool
	ID         []byte
}

type AnnotMap map[int]AnnotationRenderer

type Annot struct {
	IndRefs *[]types.IndirectRef
	Map     AnnotMap
}

type PgAnnots map[AnnotationType]Annot

type XRefTable struct {
	Table               map[int]*XRefTableEntry
	Size                *int
	MaxObjNr            int
	PageCount           int
	Root                *types.IndirectRef
	RootDict            types.Dict
	Names               map[string]*Node
	Dests               types.Dict
	NameRefs            map[string]NameMap
	Encrypt             *types.IndirectRef
	E                   *Enc
	EncKey              []byte
	AES4Strings         bool
	AES4Streams         bool
	AES4EmbeddedStreams bool

	HeaderVersion *Version
	RootVersion   *Version

	ID             types.Array
	Info           *types.IndirectRef
	Title          string
	Subject        string
	Author         string
	Creator        string
	Producer       string
	CreationDate   string
	ModDate        string
	Keywords       string
	KeywordList    types.StringSet
	Properties     map[string]string
	CatalogXMPMeta *XMPMeta

	PageLayout *PageLayout
	PageMode   *PageMode
	ViewerPref *ViewerPreferences

	OffsetPrimaryHintTable  *int64
	OffsetOverflowHintTable *int64
	LinearizationObjs       types.IntSet

	PageAnnots map[int]PgAnnots

	PageThumbs map[int]types.IndirectRef

	Signatures           map[int]map[int]Signature
	URSignature          types.Dict
	URSignatureIncrement int
	CertifiedSigObjNr    int
	DSS                  types.Dict
	DTS                  time.Time

	AdditionalStreams *types.Array

	Stats PDFStats

	Tagged           bool
	CustomExtensions bool

	CurPage        int
	CurObj         int
	Conf           *Configuration
	ValidationMode int
	ValidateLinks  bool
	Valid          bool
	URIs           map[int]map[string]string

	Optimized      bool
	Watermarked    bool
	Form           types.Dict
	Outlines       types.Dict
	SignatureExist bool
	AppendOnly     bool

	UsedGIDs  map[string]map[uint16]bool
	FillFonts map[string]types.IndirectRef
}

func newXRefTable(conf *Configuration) (xRefTable *XRefTable) {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) Version() Version { _ = "STUB: not implemented"; return *new(Version) }

func (xRefTable *XRefTable) PDF20() bool { _ = "STUB: not implemented"; return false }

func (xRefTable *XRefTable) VersionString() string { _ = "STUB: not implemented"; return "" }

func (xRefTable *XRefTable) ParseRootVersion() (v *string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) ValidateVersion(element string, sinceVersion Version) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) currentCommand() CommandMode {
	_ = "STUB: not implemented"
	return *new(CommandMode)
}

func (xRefTable *XRefTable) IsMerging() bool { _ = "STUB: not implemented"; return false }

func (xRefTable *XRefTable) EnsureVersionForWriting() { _ = "STUB: not implemented"; return }

func (xRefTable *XRefTable) IsLinearizationObject(i int) bool {
	_ = "STUB: not implemented"
	return false
}

func (xRefTable *XRefTable) LinearizationObjsString() (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

func (xRefTable *XRefTable) Exists(objNr int) bool { _ = "STUB: not implemented"; return false }

func (xRefTable *XRefTable) Find(objNr int) (*XRefTableEntry, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (xRefTable *XRefTable) FindObject(objNr int) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func (xRefTable *XRefTable) Free(objNr int) (*XRefTableEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) NextForFree(objNr int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (xRefTable *XRefTable) FindTableEntryLight(objNr int) (*XRefTableEntry, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (xRefTable *XRefTable) FindTableEntry(objNr int, genNr int) (*XRefTableEntry, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (xRefTable *XRefTable) FindTableEntryForIndRef(indRef *types.IndirectRef) (*XRefTableEntry, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (xRefTable *XRefTable) IncrementRefCount(indRef *types.IndirectRef) {
	_ = "STUB: not implemented"
	return
}

func (xRefTable *XRefTable) InsertNew(xRefTableEntry XRefTableEntry) (objNr int) {
	_ = "STUB: not implemented"
	return 0
}

func (xRefTable *XRefTable) InsertAndUseRecycled(xRefTableEntry XRefTableEntry) (objNr int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (xRefTable *XRefTable) InsertObject(obj types.Object) (objNr int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (xRefTable *XRefTable) IndRefForObject(objNr int, obj types.Object) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) IndRefForNewObject(obj types.Object) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) NewStreamDictForBuf(buf []byte) (*types.StreamDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) NewStreamDictForFile(filename string) (*types.StreamDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) finalizeEmbeddedStreamDict(
	sd *types.StreamDict,
	size int,
	modDate time.Time,
) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) NewEmbeddedStreamDict(r io.Reader, modDate time.Time) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) locateObjForIndRef(ir types.IndirectRef) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func (xRefTable *XRefTable) FreeObject(objNr int) error { _ = "STUB: not implemented"; return nil }

func (xRefTable *XRefTable) DeleteObject(o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) DeleteObjectGraph(o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) NewEmbeddedFileStreamDict(filename string) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) NewSoundStreamDict(filename string, samplingRate int, fileSpecDict types.Dict) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) NewFileSpecDict(f, uf, desc string, indRefStreamDict types.IndirectRef) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func (xRefTable *XRefTable) freeObjects() types.IntSet {
	_ = "STUB: not implemented"
	return *new(types.IntSet)
}

func anyKey(m types.IntSet) int { _ = "STUB: not implemented"; return 0 }

func (xRefTable *XRefTable) handleDanglingFree(m types.IntSet, head *XRefTableEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) validateFreeList(f int, m types.IntSet, e *XRefTableEntry) (*XRefTableEntry, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (xRefTable *XRefTable) EnsureValidFreeList() error { _ = "STUB: not implemented"; return nil }

func (xRefTable *XRefTable) DeleteDictEntry(d types.Dict, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) UndeleteObject(objectNumber int) error {
	_ = "STUB: not implemented"
	return nil
}

func validateObjNrAndGenNr(objNr, genNr int) error { _ = "STUB: not implemented"; return nil }

func (xRefTable *XRefTable) IsObjValid(objNr, genNr int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (xRefTable *XRefTable) IsValid(ir types.IndirectRef) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (xRefTable *XRefTable) IsObjBeingValidated(objNr, genNr int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (xRefTable *XRefTable) IsBeingValidated(ir types.IndirectRef) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (xRefTable *XRefTable) SetValid(ir types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) SetBeingValidated(ir types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) DereferenceStreamDict(o types.Object) (*types.StreamDict, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (xRefTable *XRefTable) DereferenceXObjectDict(indRef types.IndirectRef) (*types.StreamDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) Catalog() (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func (xRefTable *XRefTable) EncryptDict() (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func (xRefTable *XRefTable) CatalogHasPieceInfo() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (xRefTable *XRefTable) Pages() (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) MissingObjects() (int, *string) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (xRefTable *XRefTable) sortedKeys() []int { _ = "STUB: not implemented"; return nil }

func objStr(entry *XRefTableEntry, objNr int) string { _ = "STUB: not implemented"; return "" }

func (xRefTable *XRefTable) DumpObject(objNr, mode int) { _ = "STUB: not implemented"; return }

func (xRefTable *XRefTable) list(logStr []string) []string { _ = "STUB: not implemented"; return nil }

func (xRefTable *XRefTable) freeList(logStr []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type nameTreeFrame struct {
	n       *Node
	root    bool
	visited bool
}

func appendNameTreeChildFrames(stack []nameTreeFrame, n *Node) []nameTreeFrame {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) bindNameTreeNodeDict(name string, n *Node, root bool) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func bindNameTreeNodeLimits(dict types.Dict, n *Node, root bool) { _ = "STUB: not implemented"; return }

func bindNameTreeLeafNode(dict types.Dict, n *Node) { _ = "STUB: not implemented"; return }

func (xRefTable *XRefTable) bindNameTreeIntermediateNode(dict types.Dict, n *Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) bindNameTreeNodeFrame(name string, f nameTreeFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) bindNameTreeNode(name string, n *Node, root bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) BindNameTrees() error { _ = "STUB: not implemented"; return nil }

func (xRefTable *XRefTable) LocateNameTree(nameTreeName string, ensure bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) NamesDict() (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func (xRefTable *XRefTable) RemoveNameTree(nameTreeName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) RemoveCollection() error { _ = "STUB: not implemented"; return nil }

func (xRefTable *XRefTable) EnsureCollection() error { _ = "STUB: not implemented"; return nil }

func (xRefTable *XRefTable) RemoveEmbeddedFilesNameTree() error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) IDFirstElement() (id []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type InheritedPageAttrs struct {
	Resources types.Dict
	MediaBox  *types.Rectangle
	CropBox   *types.Rectangle
	Rotate    int
}

func rect(xRefTable *XRefTable, a types.Array) (*types.Rectangle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func weaveResourceSubDict(d1, d2 types.Dict) { _ = "STUB: not implemented"; return }

func (xRefTable *XRefTable) consolidateResources(obj types.Object, pAttrs *InheritedPageAttrs) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) checkInheritedPageAttrs(pageDict types.Dict, pAttrs *InheritedPageAttrs, consolidateRes bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) decodeContentStream(sd *types.StreamDict, pageNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) PageContent(d types.Dict, pageNr int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) consolidateResourceSubDict(d types.Dict, key string, prn PageResourceNames, pageNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) consolidateResourceDict(d types.Dict, prn PageResourceNames, pageNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) consolidateResourcesWithContent(pageDict, resDict types.Dict, pageNr int, consolidateRes bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) pageObjType(indRef types.IndirectRef) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func errForUnexpectedPageObjectType(validationMode int, objType string, indRef types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) processPageTreeKidForPageDict(indRef types.IndirectRef, pAttrs *InheritedPageAttrs, p *int, page int, consolidateRes bool, depth int, visit *PageTreeVisit) (types.Dict, *types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil, nil
}

func (xRefTable *XRefTable) processPageTreeForPageDictDepth(root *types.IndirectRef, pAttrs *InheritedPageAttrs, p *int, page int, consolidateRes bool, depth int, visit *PageTreeVisit) (types.Dict, *types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil, nil
}

func (xRefTable *XRefTable) processPageTreeForPageDict(root *types.IndirectRef, pAttrs *InheritedPageAttrs, p *int, page int, consolidateRes bool) (types.Dict, *types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil, nil
}

func (xRefTable *XRefTable) PageDict(pageNr int, consolidateRes bool) (types.Dict, *types.IndirectRef, *InheritedPageAttrs, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil, nil, nil
}

func cloneInheritedPageAttrs(pAttrs InheritedPageAttrs) InheritedPageAttrs {
	_ = "STUB: not implemented"
	return *new(InheritedPageAttrs)
}

func (xRefTable *XRefTable) consolidatePageResourcesForKid(
	o types.Object,
	pAttrs InheritedPageAttrs,
	pageNr *int,
	depth int,
	visit *PageTreeVisit,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) consolidatePageResourcesForNode(
	root *types.IndirectRef,
	pAttrs InheritedPageAttrs,
	pageNr *int,
	depth int,
	visit *PageTreeVisit,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) ConsolidatePageResources() error { _ = "STUB: not implemented"; return nil }

func (xRefTable *XRefTable) PageDictIndRef(page int) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) processPageTreeForPageNumberDepth(root *types.IndirectRef, pageCount *int, pageObjNr int, depth int, visit *PageTreeVisit) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (xRefTable *XRefTable) processPageTreeForPageNumber(root *types.IndirectRef, pageCount *int, pageObjNr int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (xRefTable *XRefTable) PageNumber(pageObjNr int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (xRefTable *XRefTable) EnsurePageCount() error { _ = "STUB: not implemented"; return nil }

func (xRefTable *XRefTable) resolvePageBoundary(d types.Dict, boxName string) (*types.Rectangle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) collectPageBoundariesForPage(d types.Dict, pb []PageBoundaries, inhMediaBox, inhCropBox *types.Rectangle, rot, p int) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) collectMediaBoxAndCropBox(d types.Dict, inhMediaBox, inhCropBox **types.Rectangle) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) collectPageBoundariesForPageTreeKids(
	parentObjNr int,
	kids types.Array,
	inhMediaBox, inhCropBox **types.Rectangle,
	pb []PageBoundaries,
	r int,
	p *int,
	selectedPages types.IntSet,
	depth int,
	visit *PageTreeVisit) error {
	_ = "STUB: not implemented"
	return nil
}

func pageTreeNodeType(d types.Dict, objNr int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (xRefTable *XRefTable) pageTreeNodeRotation(d types.Dict, objNr, inherited int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (xRefTable *XRefTable) collectPageBoundariesForPageTree(
	root *types.IndirectRef,
	inhMediaBox, inhCropBox **types.Rectangle,
	pb []PageBoundaries,
	r int,
	p *int,
	selectedPages types.IntSet,
	depth int,
	visit *PageTreeVisit) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) PageBoundaries(selectedPages types.IntSet) ([]PageBoundaries, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) PageDims() ([]types.Dim, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) EmptyPage(parentIndRef *types.IndirectRef, mediaBox *types.Rectangle, objNr int) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) pageMediaBox(d types.Dict) (*types.Rectangle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) emptyPage(parent *types.IndirectRef, d types.Dict, dim *types.Dim, pAttrs *InheritedPageAttrs) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type blankPageInsertion struct {
	parent        *types.IndirectRef
	pAttrs        *InheritedPageAttrs
	p             *int
	selectedPages types.IntSet
	dim           *types.Dim
	before        bool
	depth         int
	visit         *PageTreeVisit
}

func (xRefTable *XRefTable) appendBlankPageForPage(a *types.Array, ir types.IndirectRef, pageNodeDict types.Dict, ctx blankPageInsertion) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (xRefTable *XRefTable) appendBlankPagesForKid(a *types.Array, o types.Object, ctx blankPageInsertion) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (xRefTable *XRefTable) insertBlankPagesDepth(
	parent *types.IndirectRef,
	pAttrs *InheritedPageAttrs,
	p *int, selectedPages types.IntSet,
	dim *types.Dim,
	before bool,
	depth int,
	visit *PageTreeVisit) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (xRefTable *XRefTable) insertBlankPages(
	parent *types.IndirectRef,
	pAttrs *InheritedPageAttrs,
	p *int, selectedPages types.IntSet,
	dim *types.Dim,
	before bool) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (xRefTable *XRefTable) InsertBlankPages(pages types.IntSet, dim *types.Dim, before bool) error {
	_ = "STUB: not implemented"
	return nil
}

func weaveInPage(ctx *Context, parent types.IndirectRef, pageNr int) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) insertPagesDepth(parent *types.IndirectRef, p *int, ctx *Context, depth int, visit *PageTreeVisit) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (xRefTable *XRefTable) InsertPages(parent *types.IndirectRef, p *int, ctx *Context) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (xRefTable *XRefTable) AppendPages(rootPageIndRef *types.IndirectRef, fromPageNr int, ctx *Context) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (xRefTable *XRefTable) StreamDictIndRef(bb []byte) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) insertContent(pageDict types.Dict, bb []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func appendToContentStream(sd *types.StreamDict, bb []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) AppendContent(pageDict types.Dict, bb []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) HasUsedGIDs(fontName string) bool {
	_ = "STUB: not implemented"
	return false
}

func (xRefTable *XRefTable) NameRef(nameType string) NameMap {
	_ = "STUB: not implemented"
	return *new(NameMap)
}

func (xRefTable *XRefTable) RemoveSignature() { _ = "STUB: not implemented"; return }

func removePageAnnotationForSig(xRefTable *XRefTable, pIndRef, indRef types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func removeSigAnnot(xRefTable *XRefTable, indRef types.IndirectRef, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func (xRefTable *XRefTable) RemoveAllSignatures() error { _ = "STUB: not implemented"; return nil }

func (xRefTable *XRefTable) BindPrinterPreferences(vp *ViewerPreferences, d types.Dict) {
	_ = "STUB: not implemented"
	return
}

func (xRefTable *XRefTable) BindViewerPreferences() { _ = "STUB: not implemented"; return }

func (xRefTable *XRefTable) RectForArray(a types.Array) (*types.Rectangle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
