package api

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pkg/errors"
)

var (
	ErrNoOutlines = errors.New("pdfcpu: no outlines available")
	ErrOutlines   = errors.New("pdfcpu: existing outlines")
)

func Bookmarks(rs io.ReadSeeker, conf *model.Configuration) (bms []pdfcpu.Bookmark, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExportBookmarksJSON(rs io.ReadSeeker, w io.Writer, source string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ExportBookmarksFile(inFilePDF, outFileJSON string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ImportBookmarks(rs io.ReadSeeker, rd io.Reader, w io.Writer, replace bool, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ImportBookmarksFile(inFilePDF, inFileJSON, outFilePDF string, replace bool, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AddBookmarks(rs io.ReadSeeker, w io.Writer, bms []pdfcpu.Bookmark, replace bool, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func AddBookmarksFile(inFile, outFile string, bms []pdfcpu.Bookmark, replace bool, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemoveBookmarks(rs io.ReadSeeker, w io.Writer, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func RemoveBookmarksFile(inFile, outFile string, conf *model.Configuration) (err error) {
	_ = "STUB: not implemented"
	return nil
}
