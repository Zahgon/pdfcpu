package validate

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func validateEntryV(xRefTable *model.XRefTable, d types.Dict, dictName string, required bool, sinceVersion model.Version, pBeadIndRef *types.IndirectRef, objNumber int) error {
	_ = "STUB: not implemented"
	return nil
}

func validateBeadDict(xRefTable *model.XRefTable, beadIndRef, threadIndRef, pBeadIndRef, lBeadIndRef *types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func soleBeadDict(beadIndRef, pBeadIndRef, nBeadIndRef *types.IndirectRef) bool {
	_ = "STUB: not implemented"
	return false
}

func validateBeadChainIntegrity(beadIndRef, pBeadIndRef, nBeadIndRef *types.IndirectRef) bool {
	_ = "STUB: not implemented"
	return false
}

func validateFirstBeadDict(xRefTable *model.XRefTable, beadIndRef, threadIndRef *types.IndirectRef) error {
	_ = "STUB: not implemented"
	return nil
}

func validateThreadDict(xRefTable *model.XRefTable, o types.Object, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateThreads(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}
