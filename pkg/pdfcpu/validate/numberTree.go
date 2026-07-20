package validate

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func validatePageLabelDict(xRefTable *model.XRefTable, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNumberTreeDictNumsEntry(xRefTable *model.XRefTable, d types.Dict, name string, useIDs bool) (firstKey, lastKey int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func validateNumberTreeDictLimitsEntry(xRefTable *model.XRefTable, d types.Dict, firstKey, lastKey int) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNumberTree(xRefTable *model.XRefTable, name string, d types.Dict, root, useIDs bool) (firstKey, lastKey int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func validateNumberTreeDepth(xRefTable *model.XRefTable, name string, d types.Dict, root, useIDs bool, depth int) (firstKey, lastKey int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}
