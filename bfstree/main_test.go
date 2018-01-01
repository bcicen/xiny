package bfstree

import "testing"

var tree BFSTree

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
	t.Logf("created tree with %d edges, %d nodes", len(tree), len(tree.Nodes()))
}

func TestFindPath(t *testing.T) {

	path, err := tree.FindPath("a", "g")
	if err != nil {
		t.Error(err)
	}
	t.Logf("found path: %s", path)
}
