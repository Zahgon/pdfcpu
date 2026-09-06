package validate

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func validateTilingPatternDict(xRefTable *model.XRefTable, sd *types.StreamDict, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateShadingPatternDict(xRefTable *model.XRefTable, d types.Dict, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePattern(xRefTable *model.XRefTable, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePatternResourceDict(xRefTable *model.XRefTable, o types.Object, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}
