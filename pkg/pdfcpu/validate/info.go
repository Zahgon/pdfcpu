package validate

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func DocumentProperty(s string) bool { _ = "STUB: not implemented"; return false }

func validateInfoDictDate(xRefTable *model.XRefTable, name string, o types.Object) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func validateInfoDictTrapped(xRefTable *model.XRefTable, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func handleProperties(xRefTable *model.XRefTable, key string, val types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func validateKeywords(xRefTable *model.XRefTable, v types.Object) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func validateDocInfoDictEntry(xRefTable *model.XRefTable, k string, v types.Object) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func validateDocumentInfoDict(xRefTable *model.XRefTable, obj types.Object) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func validateDocumentInfoObject(xRefTable *model.XRefTable) error {
	_ = "STUB: not implemented"
	return nil
}

func DocumentPageLayout(s string) bool { _ = "STUB: not implemented"; return false }

func DocumentPageMode(s string) bool { _ = "STUB: not implemented"; return false }
