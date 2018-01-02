package bfstree

import (
	"testing"
)

var tree *BFSTree

// TestEdge is an Edge interface implementation
type TestEdge struct {
	from string
	to   string
}

func (t TestEdge) To() string   { return t.to }
func (t TestEdge) From() string { return t.from }

func TestCreateTree(t *testing.T) {
	tree = NewBFSTree(
		TestEdge{"a", "b"},
		TestEdge{"b", "d"},
		TestEdge{"d", "e"},
		TestEdge{"e", "f"},
		TestEdge{"f", "g"},
		TestEdge{"c", "d"},
		TestEdge{"c", "a"},
		TestEdge{"a", "c"},
		TestEdge{"b", "c"},
	)
	t.Logf("created tree with %d edges, %d nodes", tree.Len(), len(tree.Nodes()))
}

func TestFindLongPath(t *testing.T) {
	path, err := tree.FindPath("a", "g")
	if err != nil {
		t.Error(err)
	}
	t.Logf("found path: %s", path)
}

func TestFindShortPath(t *testing.T) {
	path, err := tree.FindPath("a", "b")
	if err != nil {
		t.Error(err)
	}
	t.Logf("found path: %s", path)
}

func TestFindNoPath(t *testing.T) {
	_, err := tree.FindPath("a", "z")
	if err == nil {
		t.Errorf("no error returned on missing path")
	}
	t.Logf("got expected error: %s", err)
}
