package validate

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"github.com/pkg/errors"
)

var ErrBookmarksRepair = errors.New("pdfcpu: bookmarks repair failed")

func validateOutlineItemDictTitle(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOutlineItemDictParent(xRefTable *model.XRefTable, d types.Dict, dictName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOutlineItemDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func handleOutlineItemDict(xRefTable *model.XRefTable, ir types.IndirectRef, objNumber int) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func leaf(firstChild, lastChild *types.IndirectRef, objNumber, validationMode int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func evalOutlineCount(xRefTable *model.XRefTable, c, visc int, count int, total, visible *int) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOutlineTree(xRefTable *model.XRefTable, first, last *types.IndirectRef, m map[int]bool, fixed *bool) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func validateOutlineTreeDepth(xRefTable *model.XRefTable, first, last *types.IndirectRef, m map[int]bool, fixed *bool, depth int) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func validateVisibleOutlineCount(xRefTable *model.XRefTable, total, visible int, count *int) error {
	_ = "STUB: not implemented"
	return nil
}

func validateInvisibleOutlineCount(xRefTable *model.XRefTable, total int, count *int) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOutlineCount(xRefTable *model.XRefTable, total, visible int, count *int) error {
	_ = "STUB: not implemented"
	return nil
}

func firstOfRemainder(xRefTable *model.XRefTable, last *types.IndirectRef, duplObjNr, oneBeforeDuplObj int) (int, types.Dict, error) {
	_ = "STUB: not implemented"
	return 0, *new(types.Dict), nil
}

func removeDuplFirst(xRefTable *model.XRefTable, first, last *types.IndirectRef, duplObjNr, oneBeforeDuplObj int) error {
	_ = "STUB: not implemented"
	return nil
}

func handleCircular(xRefTable *model.XRefTable, dict types.Dict, first *types.IndirectRef, fixed *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func handleCorruptDict(xRefTable *model.XRefTable) error { _ = "STUB: not implemented"; return nil }

func handleDuplicate(
	xRefTable *model.XRefTable,
	ir, first, last *types.IndirectRef,
	prevDict types.Dict,
	objNr, prevObjNr int) error {
	_ = "STUB: not implemented"
	return nil
}

func scanAndFixOutlineItems(xRefTable *model.XRefTable, first, last *types.IndirectRef, seen map[int]bool, fixed *bool) error {
	_ = "STUB: not implemented"
	return nil
}

func removeOutlines(xRefTable *model.XRefTable, rootDict types.Dict) {
	_ = "STUB: not implemented"
	return
}

func validateOutlinesGeneral(xRefTable *model.XRefTable, rootDict types.Dict) (*types.IndirectRef, *types.IndirectRef, *int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func handleCorruptOutlineItems(xRefTable *model.XRefTable, rootDict types.Dict) {
	_ = "STUB: not implemented"
	return
}

func scanAndFixOutlines(xRefTable *model.XRefTable, rootDict types.Dict, first, last *types.IndirectRef, count *int) error {
	_ = "STUB: not implemented"
	return nil
}

func validateOutlines(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}
