package pdfcpu

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func PropertiesList(ctx *model.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PropertiesAdd(ctx *model.Context, properties map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func PropertiesRemove(ctx *model.Context, properties []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func removeAllProperties(ctx *model.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
