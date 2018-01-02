package bfstree

import (
	"fmt"
	"strings"
)

type Edge interface {
	From() string
	To() string
}

type BFSTree struct {
	edges []Edge
}

func NewBFSTree(edges ...Edge) *BFSTree { return &BFSTree{edges} }

func (b *BFSTree) Len() int       { return len(b.edges) }
func (b *BFSTree) Edges() []Edge  { return b.edges }
func (b *BFSTree) AddEdge(e Edge) { b.edges = append(b.edges, e) }

// Return unique node names
func (b *BFSTree) Nodes() []string {
	var names []string
	for _, e := range b.edges {
		names = append(names, e.To())
		names = append(names, e.From())
	}
	return uniq(names)
}

// return edges from a given start point
func (b *BFSTree) fromNode(start string) (res []Edge) {
	for _, e := range b.edges {
		if e.From() == start {
			res = append(res, e)
		}
	}
	return res
}

func (b *BFSTree) FindPath(start string, end string) (path *Path, err error) {
	var paths []*Path

	// Create start paths from origin node
	for _, e := range b.fromNode(start) {
		p := NewPath(e)
		if e.To() == end {
			return p, nil
		}
		paths = append(paths, p)
	}

	for len(paths) > 0 {
		var newPaths []*Path

		fmt.Printf("nopaths %d\n", len(paths))

		for _, p := range paths {
			children := b.fromNode(p.Last().To())

			// maximum path depth reached, drop
			if len(children) == 0 {
				continue
			}

			// branch path for each child node
			for _, e := range children {
				// drop circular paths
				if p.IsCircular(e) {
					continue
				}

				np := NewPath(p.edges...)
				np.AddEdge(e)

				if e.To() == end {
					return np, nil
				}
				newPaths = append(newPaths, np)
			}
		}
		paths = newPaths
	}

	return path, fmt.Errorf("no path found")
}

type Path struct {
	*BFSTree
}

func NewPath(edges ...Edge) *Path { return &Path{&BFSTree{edges}} }

func (p *Path) Last() Edge { return p.edges[len(p.edges)-1] }

// Returns names for all path nodes in the order they are transversed
func (p *Path) Nodes() []string {
	names := []string{p.edges[0].From()}
	for _, e := range p.edges {
		names = append(names, e.To())
	}
	return names
}

func (p *Path) String() string { return strings.Join(p.Nodes(), "->") }

// Return whether a given edge, if added, would result in
// a circular or recursive path
func (p *Path) IsCircular(edge Edge) bool {
	child := edge.To()
	for _, e := range p.edges {
		if e.From() == child || e.To() == child {
			return true
		}
	}
	return false
}

// Return whether this path transverses a given node name
func (p *Path) HasNode(s string) bool {
	for _, e := range p.edges {
		if e.From() == s || e.To() == s {
			return true
		}
	}
	return false
}

// uniq returns a unique subset of the string slice provided.
func uniq(a []string) []string {
	u := make([]string, 0, len(a))
	m := make(map[string]bool)

	for _, val := range a {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}
