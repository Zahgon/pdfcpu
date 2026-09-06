package validate

import (
	"errors"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

var errMissingNameTreeKidsOrNames = errors.New("missing Kids or Names")

func validateDestsNameTreeValue(xRefTable *model.XRefTable, o types.Object, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAPNameTreeValue(xRefTable *model.XRefTable, o types.Object, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateJavaScriptNameTreeValue(xRefTable *model.XRefTable, o types.Object, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePagesNameTreeValue(xRefTable *model.XRefTable, o types.Object, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateTemplatesNameTreeValue(xRefTable *model.XRefTable, o types.Object, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateURLAliasDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCommandSettingsDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCaptureCommandDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSourceInfoDictEntryAU(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSourceInfoDict(xRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateEntrySI(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateWebCaptureContentSetDict(XRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateIDSNameTreeValue(xRefTable *model.XRefTable, o types.Object, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateURLSNameTreeValue(xRefTable *model.XRefTable, o types.Object, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateEmbeddedFilesNameTreeValue(xRefTable *model.XRefTable, o types.Object, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSlideShowDict(XRefTable *model.XRefTable, d types.Dict) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAlternatePresentationsNameTreeValue(xRefTable *model.XRefTable, o types.Object, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateRenditionsNameTreeValue(xRefTable *model.XRefTable, o types.Object, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateIDTreeValue(xRefTable *model.XRefTable, o types.Object, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNameTreeValue(name string, xRefTable *model.XRefTable, o types.Object) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func validateNameTreeDictNamesEntry(xRefTable *model.XRefTable, d types.Dict, name string, node *model.Node) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func validateNameTreeDictLimitsEntry(xRefTable *model.XRefTable, d types.Dict, firstKey, lastKey string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNameTree(xRefTable *model.XRefTable, name string, d types.Dict, root bool) (string, string, *model.Node, error) {
	_ = "STUB: not implemented"
	return "", "", nil, nil
}

func nameTreeKidContext(name string, o types.Object, i int) string {
	_ = "STUB: not implemented"
	return ""
}

func validateNameTreeKids(
	xRefTable *model.XRefTable,
	name string,
	a types.Array,
	node *model.Node,
	depth int,
	specViolations *[]error,
) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func validateNameTreeDepth(xRefTable *model.XRefTable, name string, d types.Dict, root bool, depth int) (string, string, *model.Node, error) {
	_ = "STUB: not implemented"
	return "", "", nil, nil
}

func validateNameTreeDepthWithViolations(
	xRefTable *model.XRefTable,
	name string,
	d types.Dict,
	root bool,
	depth int,
	specViolations *[]error,
) (string, string, *model.Node, error) {
	_ = "STUB: not implemented"
	return "", "", nil, nil
}
