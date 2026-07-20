package pdfcpu

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func migrateIndRef(ir *types.IndirectRef, ctxSource, ctxDest *model.Context, migrated map[int]int) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func migrateObject(o types.Object, ctxSource, ctxDest *model.Context, migrated map[int]int) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func migrateAnnots(o types.Object, pageIndRef types.IndirectRef, ctxSrc, ctxDest *model.Context, migrated map[int]int) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func migratePageDict(d types.Dict, pageIndRef types.IndirectRef, ctxSrc, ctxDest *model.Context, migrated map[int]int) error {
	_ = "STUB: not implemented"
	return nil
}

func migrateAnnot(indRef *types.IndirectRef, fieldsSrc, fieldsDest *types.Array, ctxSrc *model.Context, migrated map[int]int) error {
	_ = "STUB: not implemented"
	return nil
}

func migrateFields(d types.Dict, fieldsSrc, fieldsDest *types.Array, ctxSrc, ctxDest *model.Context, migrated map[int]int) error {
	_ = "STUB: not implemented"
	return nil
}

func migrateFormDict(d types.Dict, fields types.Array, ctxSrc, ctxDest *model.Context, migrated map[int]int) error {
	_ = "STUB: not implemented"
	return nil
}

func detectMigratedAnnot(ctxSrc *model.Context, indRef *types.IndirectRef, kids types.Array, migrated map[int]int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
