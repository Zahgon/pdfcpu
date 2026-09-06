package model

import (
	"errors"
	"io"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type Attachment struct {
	io.Reader
	ID       string
	FileName string
	Desc     string
	ModTime  *time.Time
}

func (a Attachment) String() string { _ = "STUB: not implemented"; return "" }

func decodeFileSpecStreamDict(sd *types.StreamDict) error { _ = "STUB: not implemented"; return nil }

func fileSpecStreamFileName(xRefTable *XRefTable, d types.Dict) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func fileSpecStreamDict(xRefTable *XRefTable, d types.Dict) (*types.StreamDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (xRefTable *XRefTable) NewFileSpecDictForAttachment(a Attachment) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

var (
	errInvalidModDateType = errors.New("invalid ModDate type")
	errInvalidModDate     = errors.New("invalid ModDate")
)

func getModDate(xRefTable *XRefTable, obj types.Object) (*time.Time, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fileSpecStreamDictInfo(xRefTable *XRefTable, id string, o types.Object, decode bool) (*types.StreamDict, string, string, *time.Time, error) {
	_ = "STUB: not implemented"
	return nil, "", "", nil, nil
}

func (ctx *Context) ListAttachments() ([]Attachment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctx *Context) AddAttachment(a Attachment, useCollection bool) error {
	_ = "STUB: not implemented"
	return nil
}

var errContentMatch = errors.New("name tree content match")

func (ctx *Context) SearchEmbeddedFilesNameTreeNodeByContent(s string) (*string, types.Object, error) {
	_ = "STUB: not implemented"
	return nil, *new(types.Object), nil
}

func (ctx *Context) removeAttachment(id string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ctx *Context) RemoveAttachments(ids []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ctx *Context) RemoveAttachment(a Attachment) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ctx *Context) ExtractAttachments(ids []string) ([]Attachment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctx *Context) ExtractAttachment(a Attachment) (*Attachment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctx *Context) AddAttachmentsToInfoDigest(ss *[]string) error {
	_ = "STUB: not implemented"
	return nil
}
