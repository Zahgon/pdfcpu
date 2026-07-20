package cli

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func NUp(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func Booklet(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func Resize(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func Poster(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func NDown(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func Cut(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func Zoom(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func Rotate(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func InsertPages(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func RemovePages(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func Crop(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func listBoxes(rs io.ReadSeeker, selectedPages []string, pb *model.PageBoundaries, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListBoxesFile(inFile string, selectedPages []string, pb *model.PageBoundaries, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListBoxes(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func AddBoxes(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func RemoveBoxes(cmd *Command) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
