package model

import (
	"errors"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

var errNameTreeDuplicateKey = errors.New("name: duplicate key")

const maxEntries = 3

type Node struct {
	Kids       []*Node
	Names      []entry
	Kmin, Kmax string
	D          types.Dict
}

type entry struct {
	k string
	v types.Object
}

func (n Node) leaf() bool { _ = "STUB: not implemented"; return false }

func (n Node) emptyLeaf() bool { _ = "STUB: not implemented"; return false }

func keyLess(k, s string) bool { _ = "STUB: not implemented"; return false }

func keyLessOrEqual(k, s string) bool { _ = "STUB: not implemented"; return false }

func (n Node) withinLimits(k string) bool { _ = "STUB: not implemented"; return false }

func (n Node) Value(k string) (types.Object, bool) {
	_ = "STUB: not implemented"
	return *new(types.Object), false
}

func (n *Node) AppendToNames(k string, v types.Object) { _ = "STUB: not implemented"; return }

type NameMap map[string][]types.Dict

func (m NameMap) Add(k string, d types.Dict) { _ = "STUB: not implemented"; return }

func (n *Node) insertIntoLeaf(k string, v types.Object, m NameMap) error {
	_ = "STUB: not implemented"
	return nil
}

func updateNameRef(d types.Dict, keys []string, nameOld, nameNew string) error {
	_ = "STUB: not implemented"
	return nil
}

func updateNameRefDicts(dd []types.Dict, nameRefDictKeys []string, nameOld, nameNew string) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *Node) insertUniqueIntoLeaf(k string, v types.Object, m NameMap, nameRefDictKeys []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (n *Node) HandleLeaf(xRefTable *XRefTable, k string, v types.Object, m NameMap, nameRefDictKeys []string) error {
	_ = "STUB: not implemented"
	return nil
}

func updateNameTreeLimits(path []*Node) { _ = "STUB: not implemented"; return }

func (n *Node) Add(xRefTable *XRefTable, k string, v types.Object, m NameMap, nameRefDictKeys []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *Node) AddTree(xRefTable *XRefTable, tree *Node, m NameMap, nameRefDictKeys []string) error {
	_ = "STUB: not implemented"
	return nil
}

func deleteNameTreeValueGraph(xRefTable *XRefTable, k string, o types.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *Node) removeFromNames(xRefTable *XRefTable, k string) (ok bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (n *Node) removeSingleFromParent(xRefTable *XRefTable, k string) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *Node) removeFromLeaf(xRefTable *XRefTable, k string) (empty, ok bool, err error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

func (n *Node) removeKid(xRefTable *XRefTable, k string, kid *Node, i int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (n *Node) removeFromKids(xRefTable *XRefTable, k string) (ok bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (n *Node) Remove(xRefTable *XRefTable, k string) (empty, ok bool, err error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

func (n *Node) Process(xRefTable *XRefTable, handler func(*XRefTable, string, *types.Object) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (n Node) KeyList() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (n Node) String() string { _ = "STUB: not implemented"; return "" }
