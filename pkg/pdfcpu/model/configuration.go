package model

import (
	_ "embed"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

const (
	ValidationStrict int = iota

	ValidationRelaxed
)

type PermissionFlags int

const (
	UnusedFlag1 PermissionFlags = 1 << iota
	UnusedFlag2
	PermissionPrintRev2
	PermissionModify
	PermissionExtract
	PermissionModAnnFillForm
	UnusedFlag7
	UnusedFlag8
	PermissionFillRev3
	PermissionExtractRev3
	PermissionAssembleRev3
	PermissionPrintRev3
)

const (
	PermissionsNone  = PermissionFlags(0xF0C3)
	PermissionsPrint = PermissionsNone + PermissionPrintRev2 + PermissionPrintRev3
	PermissionsAll   = PermissionFlags(0xFFFF)
)

type MergeBookmarkMode string

const (
	MergeBookmarkModeWrap MergeBookmarkMode = "wrap"

	MergeBookmarkModePreserve MergeBookmarkMode = "preserve"
)

const (
	StatsFileNameDefault = "stats.csv"
)

type CommandMode int

const (
	VALIDATE CommandMode = iota
	LISTINFO
	OPTIMIZE
	SPLIT
	SPLITBYPAGENR
	MERGECREATE
	MERGECREATEZIP
	MERGEAPPEND
	EXTRACTIMAGES
	EXTRACTFONTS
	EXTRACTPAGES
	EXTRACTCONTENT
	EXTRACTMETADATA
	TRIM
	LISTATTACHMENTS
	EXTRACTATTACHMENTS
	ADDATTACHMENTS
	ADDATTACHMENTSPORTFOLIO
	REMOVEATTACHMENTS
	LISTPERMISSIONS
	SETPERMISSIONS
	ADDWATERMARKS
	REMOVEWATERMARKS
	IMPORTIMAGES
	INSERTPAGESBEFORE
	INSERTPAGESAFTER
	REMOVEPAGES
	LISTKEYWORDS
	ADDKEYWORDS
	REMOVEKEYWORDS
	LISTPROPERTIES
	ADDPROPERTIES
	REMOVEPROPERTIES
	COLLECT
	CROP
	LISTBOXES
	ADDBOXES
	REMOVEBOXES
	LISTANNOTATIONS
	ADDANNOTATIONS
	REMOVEANNOTATIONS
	ROTATE
	NUP
	BOOKLET
	LISTBOOKMARKS
	ADDBOOKMARKS
	REMOVEBOOKMARKS
	IMPORTBOOKMARKS
	EXPORTBOOKMARKS
	LISTIMAGES
	UPDATEIMAGES
	CREATE
	DUMP
	LISTFORMFIELDS
	REMOVEFORMFIELDS
	LOCKFORMFIELDS
	UNLOCKFORMFIELDS
	RESETFORMFIELDS
	EXPORTFORMFIELDS
	FILLFORMFIELDS
	MULTIFILLFORMFIELDS
	ENCRYPT
	DECRYPT
	CHANGEUPW
	CHANGEOPW
	CHEATSHEETSFONTS
	INSTALLFONTS
	LISTFONTS
	RESIZE
	POSTER
	NDOWN
	CUT
	LISTPAGELAYOUT
	SETPAGELAYOUT
	RESETPAGELAYOUT
	LISTPAGEMODE
	SETPAGEMODE
	RESETPAGEMODE
	LISTVIEWERPREFERENCES
	SETVIEWERPREFERENCES
	RESETVIEWERPREFERENCES
	ZOOM
	LISTCERTIFICATES
	INSPECTCERTIFICATES
	IMPORTCERTIFICATES
	VALIDATESIGNATURES
	REMOVESIGNATURES
	ADDSIGNATURE
)

func (cmd CommandMode) AllowRemoveEncryption() bool { _ = "STUB: not implemented"; return false }

func (cmd CommandMode) AllowRemoveSignatures() bool { _ = "STUB: not implemented"; return false }

type Configuration struct {
	Path string

	CreationDate string

	Version string

	CheckFileNameExt bool

	Reader15 bool

	DecodeAllStreams bool

	ValidationMode int

	PostProcessValidate bool

	ValidateLinks bool

	Eol string

	WriteObjectStream bool

	WriteXRefStream bool

	StatsFileName string

	UserPW    string
	UserPWNew *string

	OwnerPW    string
	OwnerPWNew *string

	PrivateKeyPW string

	EncryptUsingAES bool

	EncryptKeyLength int

	Permissions PermissionFlags

	Cmd CommandMode

	Unit types.DisplayUnit

	TimestampFormat string

	DateFormat string

	Optimize bool

	OptimizeBeforeWriting bool

	OptimizeResourceDicts bool

	OptimizeDuplicateContentStreams bool

	CreateBookmarks bool

	MergeBookmarkMode MergeBookmarkMode

	NeedAppearances bool

	Offline bool

	Timeout int

	TimeoutCRL int

	TimeoutOCSP int

	PreferredCertRevocationChecker int

	FormFieldListMaxColWidth int

	Limits ResourceLimits

	RemoveEncryption bool

	RemoveSignatures bool
}

type ResourceLimits struct {
	MaxStreamBytes int64

	MaxDecodeBytes int64

	MaxImagePixels int64

	MaxImageBytes int64

	MaxObjectCount int

	MaxObjectStreamCount int

	MaxObjectStreamFirst int64

	MaxXRefEntries int

	MaxRecursionDepth int
}

func DefaultResourceLimits() ResourceLimits { _ = "STUB: not implemented"; return *new(ResourceLimits) }

var ConfigPath string = "default"

var loadedDefaultConfig *Configuration

//go:embed resources/config.yml
var configFileBytes []byte

//go:embed resources/Roboto-Regular.ttf
var robotoFontFileBytes []byte

func ensureConfigFileAt(path string, override bool) error { _ = "STUB: not implemented"; return nil }

func onlyHidden(files []os.DirEntry) bool { _ = "STUB: not implemented"; return false }

func ensureFontDirInitialized() error { _ = "STUB: not implemented"; return nil }

func initCertificates() error { _ = "STUB: not implemented"; return nil }

func EnsureDefaultConfigAt(path string, override bool) error { _ = "STUB: not implemented"; return nil }

func newDefaultConfiguration() *Configuration { _ = "STUB: not implemented"; return nil }

func ResetConfig() error { _ = "STUB: not implemented"; return nil }

func NewDefaultConfiguration() *Configuration { _ = "STUB: not implemented"; return nil }

func NewAESConfiguration(userPW, ownerPW string, keyLength int) *Configuration {
	_ = "STUB: not implemented"
	return nil
}

func NewRC4Configuration(userPW, ownerPW string, keyLength int) *Configuration {
	_ = "STUB: not implemented"
	return nil
}

func (c *Configuration) EolString() string { _ = "STUB: not implemented"; return "" }

func (c *Configuration) ValidationModeString() string { _ = "STUB: not implemented"; return "" }

func (c *Configuration) PreferredCertRevocationCheckerString() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Configuration) UnitString() string { _ = "STUB: not implemented"; return "" }

func (c *Configuration) SetUnit(s string) { _ = "STUB: not implemented"; return }

func (c *Configuration) ApplyReducedFeatureSet() bool { _ = "STUB: not implemented"; return false }
