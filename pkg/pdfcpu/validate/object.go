package validate

import (
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

const (
	REQUIRED = true

	OPTIONAL = false
)

func validateEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func validateArrayEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version, validate func(types.Array) bool) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func validateBooleanEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version, validate func(bool) bool) (*bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateFlexBooleanEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) (*bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateBooleanArrayEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version, validate func(types.Array) bool) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func timeOfDateObject(xRefTable *model.XRefTable, o types.Object, sinceVersion model.Version) (*time.Time, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateDateObject(xRefTable *model.XRefTable, o types.Object, sinceVersion model.Version) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func validateDateEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) (*time.Time, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateDictEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version, validate func(types.Dict) bool) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func validateFloat(xRefTable *model.XRefTable, o types.Object, validate func(float64) bool) (*types.Float, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateFunctionArrayEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version, validate func(types.Array) bool) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func validateFunctionOrArrayOfFunctionsEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateIndRefEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateIndRefArrayEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version, validate func(types.Array) bool) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func validateInteger(xRefTable *model.XRefTable, o types.Object, validate func(int) bool) (*types.Integer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateIntegerEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version, validate func(int) bool) (*types.Integer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateIntegerArray(xRefTable *model.XRefTable, o types.Object) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func validateIntegerArrayEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version, validate func(types.Array) bool) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func validateName(xRefTable *model.XRefTable, o types.Object, validate func(string) bool) (*types.Name, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateNameEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version, validate func(string) bool) (*types.Name, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateNameArray(xRefTable *model.XRefTable, o types.Object) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func validateNameArrayEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version, validate func(a types.Array) bool) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func validateNumber(xRefTable *model.XRefTable, o types.Object) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func validateNumberEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version, validate func(f float64) bool) (types.Object, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), nil
}

func validateNumberEntryToFloat(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version, validate func(f float64) bool) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func validateNumberArray(xRefTable *model.XRefTable, o types.Object) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func validateNumberArrayEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version, validate func(types.Array) bool) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func validateRectangleEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version, validate func(types.Array) bool) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func validateStreamDict(xRefTable *model.XRefTable, o types.Object) (*types.StreamDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateStreamDictEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version, validate func(types.StreamDict) bool) (*types.StreamDict, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decodeString(o types.Object, dictName, entryName string) (s string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func validateStringEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version, validate func(string) bool) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateStringArrayEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version, validate func(types.Array) bool) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func validateArrayArrayEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version, validate func(types.Array) bool) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func validateStringOrStreamEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNameOrStringEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateIntOrStringEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateBooleanOrStreamEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateStreamDictOrDictEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateIntegerOrArrayOfInteger(xRefTable *model.XRefTable, o types.Object, dictName, entryName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateIntegerOrArrayOfIntegerEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNameOrArrayOfName(xRefTable *model.XRefTable, o types.Object, dictName, entryName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateNameOrArrayOfNameEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}

func validateBooleanOrArrayOfBoolean(xRefTable *model.XRefTable, o types.Object, dictName, entryName string) error {
	_ = "STUB: not implemented"
	return nil
}

func validateBooleanOrArrayOfBooleanEntry(xRefTable *model.XRefTable, d types.Dict, dictName, entryName string, required bool, sinceVersion model.Version) error {
	_ = "STUB: not implemented"
	return nil
}
