package validate

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func validateMetadataStream(xRefTable *model.XRefTable, d types.Dict, required bool, sinceVersion model.Version) (*types.StreamDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateMetadata(xRefTable *model.XRefTable, d types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func catalogMetaData(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) (*model.XMPMeta, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func populateKeywordList(xRefTable *model.XRefTable, d model.Description) {
	_ = "STUB: not implemented"
	return
}

func validateRootMetadata(xRefTable *model.XRefTable, rootDict types.Dict, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}
