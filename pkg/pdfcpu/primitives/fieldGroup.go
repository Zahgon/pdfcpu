package primitives

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type FieldGroup struct {
	pdf               *PDF
	content           *Content
	Name              string
	Value             string
	Border            *Border
	Padding           *Padding
	BackgroundColor   string `json:"bgCol"`
	bgCol             *color.SimpleColor
	TextFields        []*TextField        `json:"textfield"`
	DateFields        []*DateField        `json:"datefield"`
	CheckBoxes        []*CheckBox         `json:"checkbox"`
	RadioButtonGroups []*RadioButtonGroup `json:"radiobuttongroup"`
	ComboBoxes        []*ComboBox         `json:"combobox"`
	ListBoxes         []*ListBox          `json:"listbox"`
	Hide              bool
}

func (fg *FieldGroup) validateBorder() error { _ = "STUB: not implemented"; return nil }

func (fg *FieldGroup) validatePadding() error { _ = "STUB: not implemented"; return nil }

func (fg *FieldGroup) validateBackgroundColor() error { _ = "STUB: not implemented"; return nil }

func (fg *FieldGroup) validateBorderPaddingBgCol() error { _ = "STUB: not implemented"; return nil }

func (fg *FieldGroup) validate() error { _ = "STUB: not implemented"; return nil }

func (fg *FieldGroup) mergeIn(fg0 *FieldGroup) { _ = "STUB: not implemented"; return }

func (fg *FieldGroup) calcBBoxFromTextFields(bbox **types.Rectangle, p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (fg *FieldGroup) calcBBoxFromDateFields(bbox **types.Rectangle, p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (fg *FieldGroup) calcBBoxFromCheckBoxes(bbox **types.Rectangle, p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (fg *FieldGroup) calcBBoxFromRadioButtonGroups(bbox **types.Rectangle, p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (fg *FieldGroup) calcBBoxFromComboBoxes(bbox **types.Rectangle, p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (fg *FieldGroup) calcBBoxFromListBoxes(bbox **types.Rectangle, p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (fg *FieldGroup) calcBBox(p *model.Page, pageNr int, fonts model.FontMap) (*types.Rectangle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fg *FieldGroup) renderBBox(bbox *types.Rectangle, p *model.Page) error {
	_ = "STUB: not implemented"
	return nil
}

func (fg *FieldGroup) renderTextFields(p *model.Page, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (fg *FieldGroup) renderDateFields(p *model.Page, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (fg *FieldGroup) renderCheckBoxes(p *model.Page, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (fg *FieldGroup) renderRadioButtonGroups(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (fg *FieldGroup) renderComboBoxes(p *model.Page, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (fg *FieldGroup) renderListBoxes(p *model.Page, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (fg *FieldGroup) renderFields(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (fg *FieldGroup) render(p *model.Page, pageNr int, fonts model.FontMap) error {
	_ = "STUB: not implemented"
	return nil
}
