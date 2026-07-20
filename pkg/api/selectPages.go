package api

import (
	"regexp"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

var (
	selectedPagesRegExp *regexp.Regexp
)

func setupRegExpForPageSelection() *regexp.Regexp { _ = "STUB: not implemented"; return nil }

func init() {
	selectedPagesRegExp = setupRegExpForPageSelection()
}

func ParsePageSelection(s string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func handlePrefix(v string, negated bool, pageCount int, selectedPages types.IntSet) error {
	_ = "STUB: not implemented"
	return nil
}

func handleSuffix(v string, negated bool, pageCount int, selectedPages types.IntSet) error {
	_ = "STUB: not implemented"
	return nil
}

func handleSpecificPageOrLastXPages(s string, negated bool, pageCount int, selectedPages types.IntSet) error {
	_ = "STUB: not implemented"
	return nil
}

func negation(c byte) bool { _ = "STUB: not implemented"; return false }

func selectEvenPages(selectedPages types.IntSet, pageCount int) { _ = "STUB: not implemented"; return }

func selectOddPages(selectedPages types.IntSet, pageCount int) { _ = "STUB: not implemented"; return }

func parsePageRange(pr []string, pageCount int, negated bool, selectedPages types.IntSet) error {
	_ = "STUB: not implemented"
	return nil
}

func sortedPages(selectedPages types.IntSet) []int { _ = "STUB: not implemented"; return nil }

func logSelPages(selectedPages types.IntSet) { _ = "STUB: not implemented"; return }

func calcSelPages(pageCount int, pageSelection []string, selectedPages types.IntSet) error {
	_ = "STUB: not implemented"
	return nil
}

func selectedPages(pageCount int, pageSelection []string, log bool) (types.IntSet, error) {
	_ = "STUB: not implemented"
	return *new(types.IntSet), nil
}

func PagesForPageSelection(pageCount int, pageSelection []string, ensureAllforNone bool, log bool) (types.IntSet, error) {
	_ = "STUB: not implemented"
	return *new(types.IntSet), nil
}

func RemainingPagesForPageRemoval(pageCount int, pageSelection []string, log bool) (types.IntSet, error) {
	_ = "STUB: not implemented"
	return *new(types.IntSet), nil
}

func deletePageFromCollection(cp *[]int, p int) { _ = "STUB: not implemented"; return }

func processPageForCollection(cp *[]int, negated bool, i int) { _ = "STUB: not implemented"; return }

func collectEvenPages(cp *[]int, pageCount int) { _ = "STUB: not implemented"; return }

func collectOddPages(cp *[]int, pageCount int) { _ = "STUB: not implemented"; return }

func handlePrefixForCollection(v string, negated bool, pageCount int, cp *[]int) error {
	_ = "STUB: not implemented"
	return nil
}

func handleSuffixForCollection(v string, negated bool, pageCount int, cp *[]int) error {
	_ = "STUB: not implemented"
	return nil
}

func handleSpecificPageOrLastXPagesForCollection(s string, negated bool, pageCount int, cp *[]int) error {
	_ = "STUB: not implemented"
	return nil
}

func parsePageRangeForCollection(pr []string, pageCount int, negated bool, cp *[]int) error {
	_ = "STUB: not implemented"
	return nil
}

func calcPagesForPageCollection(pageCount int, pageSelection []string) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PagesForPageCollection(pageCount int, pageSelection []string) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PagesForPageRange(from, thru int) []int { _ = "STUB: not implemented"; return nil }
