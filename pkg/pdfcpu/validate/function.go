package validate

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func validateExponentialInterpolationFunctionDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStitchingFunctionDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSampledFunctionStreamDictVersion(
	xRefTable *model.XRefTable,
	sd *types.StreamDict,
	version model.Version,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSampledFunctionStreamDict(xRefTable *model.XRefTable, sd *types.StreamDict) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePostScriptCalculatorFunctionStreamDictVersion(
	xRefTable *model.XRefTable,
	sd *types.StreamDict,
	version model.Version,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePostScriptCalculatorFunctionStreamDict(xRefTable *model.XRefTable, sd *types.StreamDict) error {
	_ = "STUB: not implemented"
	return nil
}

func processFunctionDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func processFunctionStreamDict(xRefTable *model.XRefTable, sd *types.StreamDict) error {
	_ = "STUB: not implemented"
	return nil
}

func processFunction(xRefTable *model.XRefTable, o types.Object) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func validateFunction(xRefTable *model.XRefTable, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}
