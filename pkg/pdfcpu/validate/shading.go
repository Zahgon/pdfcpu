package validate

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func validateBitsPerComponent(i int) bool { _ = "STUB: not implemented"; return false }

func validateBitsPerCoordinate(i int) bool { _ = "STUB: not implemented"; return false }

func validateBitsPerFlag(i int) bool { _ = "STUB: not implemented"; return false }

func validateShadingDictCommonEntries(xRefTable *model.XRefTable, dict types.Dict) (shadType int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func validateFunctionBasedShadingDict(xRefTable *model.XRefTable, dict types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAxialShadingDict(xRefTable *model.XRefTable, dict types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateRadialShadingDict(xRefTable *model.XRefTable, dict types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateShadingDict(xRefTable *model.XRefTable, dict types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateFreeFormGouroudShadedTriangleMeshesDict(xRefTable *model.XRefTable, dict types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateLatticeFormGouraudShadedTriangleMeshesDict(xRefTable *model.XRefTable, dict types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCoonsPatchMeshesDict(xRefTable *model.XRefTable, dict types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateTensorProductPatchMeshesDict(xRefTable *model.XRefTable, dict types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateShadingStreamDict(xRefTable *model.XRefTable, sd *types.StreamDict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateShading(xRefTable *model.XRefTable, obj types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func validateShadingResourceDict(xRefTable *model.XRefTable, obj types.Object, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}
