package pdfcpu

import (
	"errors"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/color"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

var (
	ErrNoBookmarks = errors.New("no bookmarks available")

	ErrInvalidBookmark = errors.New("invalid bookmark")

	ErrInvalidBookmarkJSON = errors.New("invalid bookmark JSON")

	ErrExistingBookmarks = errors.New("existing bookmarks")

	ErrCircularBookmarks = errors.New("circular outline item list")

	errMissingBookmarkJSONReader = errors.New("missing bookmark JSON reader")
)

func validateBookmarkContext(ctx *model.Context) error { _ = "STUB: not implemented"; return nil }

type Header struct {
	Source   string   `json:"source,omitempty"`
	Version  string   `json:"version"`
	Creation string   `json:"creation"`
	ID       []string `json:"id,omitempty"`
	Title    string   `json:"title,omitempty"`
	Author   string   `json:"author,omitempty"`
	Creator  string   `json:"creator,omitempty"`
	Producer string   `json:"producer,omitempty"`
	Subject  string   `json:"subject,omitempty"`
	Keywords string   `json:"keywords,omitempty"`
}

type Bookmark struct {
	Title    string             `json:"title"`
	PageFrom int                `json:"page"`
	PageThru int                `json:"-"`
	Bold     bool               `json:"bold,omitempty"`
	Italic   bool               `json:"italic,omitempty"`
	Color    *color.SimpleColor `json:"color,omitempty"`
	Kids     []Bookmark         `json:"kids,omitempty"`
	Parent   *Bookmark          `json:"-"`
}

type BookmarkTree struct {
	Header    Header     `json:"header"`
	Bookmarks []Bookmark `json:"bookmarks"`
}

func header(xRefTable *model.XRefTable, source string) Header {
	_ = "STUB: not implemented"
	return *new(Header)
}

func (bm Bookmark) Style() int { _ = "STUB: not implemented"; return 0 }

func positionToFirstBookmark(ctx *model.Context) (*types.IndirectRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func outlineItemTitle(s string) string { _ = "STUB: not implemented"; return "" }

func destArray(ctx *model.Context, dest types.Object) (types.Array, error) {
	_ = "STUB: not implemented"
	return *new(types.Array), nil
}

func PageNrFromDestination(ctx *model.Context, dest types.Object) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func title(ctx *model.Context, d types.Dict) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func bookmark(d types.Dict, title string, pageFrom int, parent *Bookmark) Bookmark {
	_ = "STUB: not implemented"
	return *new(Bookmark)
}

func checkBookmarkRecursionDepth(ctx *model.Context, name string, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func checkBookmarkCycle(ir *types.IndirectRef, visited map[int]bool) error {
	_ = "STUB: not implemented"
	return nil
}

func outlineItemDict(ctx *model.Context, ir *types.IndirectRef, visited map[int]bool) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func bookmarksForOutlineItem(ctx *model.Context, item *types.IndirectRef, parent *Bookmark, depth int, visited map[int]bool) ([]Bookmark, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func outlineItemDestination(ctx *model.Context, d types.Dict) (types.Object, bool, error) {
	_ = "STUB: not implemented"
	return *new(types.Object), false, nil
}

func BookmarksForOutlineItem(ctx *model.Context, item *types.IndirectRef, parent *Bookmark) ([]Bookmark, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Bookmarks(ctx *model.Context) ([]Bookmark, error) { _ = "STUB: not implemented"; return nil, nil }

func bookmarkList(bms []Bookmark, level, maxDepth int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func BookmarkList(ctx *model.Context) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func ExportBookmarks(ctx *model.Context, source string) (*BookmarkTree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExportBookmarksJSON(ctx *model.Context, source string, w io.Writer) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func bmDict(ctx *model.Context, bm Bookmark, parent types.IndirectRef) (types.Dict, error) {
	_ = "STUB: not implemented"
	return *new(types.Dict), nil
}

func invalidBookmark(i int, bm Bookmark, bms []Bookmark, parentPageNr *int) bool {
	_ = "STUB: not implemented"
	return false
}

func createOutlineItemDictDepth(ctx *model.Context, bms []Bookmark, parent *types.IndirectRef, parentPageNr *int, depth int) (*types.IndirectRef, *types.IndirectRef, int, int, error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, 0, nil
}

func createOutlineItemDict(ctx *model.Context, bms []Bookmark, parent *types.IndirectRef, parentPageNr *int) (*types.IndirectRef, *types.IndirectRef, int, int, error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, 0, nil
}

func cleanupDestinations(ctx *model.Context, dNamesEmpty bool) error {
	_ = "STUB: not implemented"
	return nil
}

func removeDest(ctx *model.Context, name string) (bool, bool, error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

func removeNamedDestForOutlineItem(ctx *model.Context, d types.Dict, ir *types.IndirectRef, depth int, visited map[int]bool) (bool, bool, error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

func removeNamedDests(ctx *model.Context, item *types.IndirectRef, depth int, visited map[int]bool) error {
	_ = "STUB: not implemented"
	return nil
}

func RemoveBookmarks(ctx *model.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func AddBookmarks(ctx *model.Context, bms []Bookmark, replace bool) error {
	_ = "STUB: not implemented"
	return nil
}

func addBookmarkTree(ctx *model.Context, bmTree *BookmarkTree, replace bool) error {
	_ = "STUB: not implemented"
	return nil
}

func parseBookmarksFromJSON(bb []byte) (*BookmarkTree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ImportBookmarks(ctx *model.Context, rd io.Reader, replace bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
