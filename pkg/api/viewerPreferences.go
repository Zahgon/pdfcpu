package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pkg/errors"
)

var ErrNoOp = errors.New("pdfcpu: no operation")

func ViewerPreferences(rs io.ReadSeeker, conf *model.Configuration) (vp *model.ViewerPreferences, v *model.Version, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func ViewerPreferencesFile(inFile string, all bool, conf *model.Configuration) (*model.ViewerPreferences, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListViewerPreferences(rs io.ReadSeeker, all bool, conf *model.Configuration) (ss []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListViewerPreferencesFileJSON(inFile string, all bool, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListViewerPreferencesFile(inFile string, all, json bool, conf *model.Configuration) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetViewerPreferences(rs io.ReadSeeker, w io.Writer, vp model.ViewerPreferences, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func SetViewerPreferencesFromJSONBytes(rs io.ReadSeeker, w io.Writer, jsonBytes []byte, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func SetViewerPreferencesFromJSONReader(rs io.ReadSeeker, w io.Writer, rd io.Reader, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func SetViewerPreferencesFile(inFile, outFile string, vp model.ViewerPreferences, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func SetViewerPreferencesFileFromJSONBytes(inFile, outFile string, jsonBytes []byte, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func SetViewerPreferencesFileFromJSONFile(inFilePDF, outFilePDF, inFileJSON string, conf *model.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func ResetViewerPreferences(rs io.ReadSeeker, w io.Writer, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ResetViewerPreferencesFile(inFile, outFile string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
