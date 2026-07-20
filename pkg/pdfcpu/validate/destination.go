package validate

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func validateDestinationArrayFirstElement(xRefTable *model.XRefTable, a types.Array) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func validateDestinationArrayLength(a types.Array) bool { _ = "STUB: not implemented"; return false }

func validateDestType(a types.Array, destType types.Name) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDestinationArray(xRefTable *model.XRefTable, a types.Array) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDestinationDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDestination(xRefTable *model.XRefTable, o types.Object, forAction bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func validateActionDestinationEntry(xRefTable *model.XRefTable, d types.Dict, dictName string, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}
