package model

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type FontObject struct {
	ResourceNames []string
	Prefix        string
	FontName      string
	FontDict      types.Dict
	Data          []byte
	Extension     string
	Embedded      bool
}

func (fo *FontObject) AddResourceName(resourceName string) { _ = "STUB: not implemented"; return }

func (fo FontObject) ResourceNamesString() string { _ = "STUB: not implemented"; return "" }

func (fo FontObject) SubType() string { _ = "STUB: not implemented"; return "" }

func (fo FontObject) Encoding() string { _ = "STUB: not implemented"; return "" }

func (fo FontObject) String() string { _ = "STUB: not implemented"; return "" }

type ImageObject struct {
	ResourceNames map[int]string
	ImageDict     *types.StreamDict
}

type DuplicateImageObject struct {
	ImageDict *types.StreamDict
	NewObjNr  int
}

func (io *ImageObject) AddResourceName(pageNr int, resourceName string) {
	_ = "STUB: not implemented"
	return
}

func (io ImageObject) ResourceNamesString() string { _ = "STUB: not implemented"; return "" }

var resourceTypes = types.NewStringSet([]string{"ColorSpace", "ExtGState", "Font", "Pattern", "Properties", "Shading", "XObject"})

type PageResourceNames map[string]types.StringSet

func NewPageResourceNames() PageResourceNames {
	_ = "STUB: not implemented"
	return *new(PageResourceNames)
}

func (prn PageResourceNames) Resources(s string) types.StringSet {
	_ = "STUB: not implemented"
	return *new(types.StringSet)
}

func (prn PageResourceNames) HasResources(s string) bool { _ = "STUB: not implemented"; return false }

func (prn PageResourceNames) HasContent() bool { _ = "STUB: not implemented"; return false }

func (prn PageResourceNames) String() string { _ = "STUB: not implemented"; return "" }
