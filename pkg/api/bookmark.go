package api

import (
	"io"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func bookmarkOpError(op string, err error) error { _ = "STUB: not implemented"; return nil }

func bookmarkSourceError(op, source string, err error) error { _ = "STUB: not implemented"; return nil }

func closeBookmarkInput(err error, f *os.File, context string) error {
	_ = "STUB: not implemented"
	return nil
}

func Bookmarks(rs io.ReadSeeker, conf *model.Configuration) (bms []pdfcpu.Bookmark, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListBookmarks(rs io.ReadSeeker, conf *model.Configuration) (ss []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ListBookmarksFile(inFile string, conf *model.Configuration) (ss []string, err error) {
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
