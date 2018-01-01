package bfstree

import (
	"fmt"
)

type Edge interface {
	From() string
	To() string
}

type Path struct {
	edges []Edge
}

func NewPath(edges ...Edge) *Path { return &Path{edges} }

func (p *Path) Edges() []Edge { return p.edges }
func (p *Path) String() string {
	s := fmt.Sprintf("%s->%s", p.edges[0].From(), p.edges[0].To())
	if len(p.edges) > 1 {
		for _, e := range p.edges[1:] {
			s += fmt.Sprintf("->%s", e.To())
		}
	}
	return s
}

func (p *Path) Append(e Edge) { p.edges = append(p.edges, e) }
func (p *Path) Last() Edge    { return p.edges[len(p.edges)-1] }

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

// return whether this path transverses a given node name
func (p *Path) HasNode(s string) bool {
	for _, e := range p.edges {
		if e.From() == s || e.To() == s {
			return true
		}
	}
	return false
}

type BFSTree struct {
	edges []Edge
}

func NewBFSTree(edges ...Edge) *BFSTree { return &BFSTree{edges} }

func (b *BFSTree) AddEdge(e Edge) { b.edges = append(b.edges, e) }

// Return unique node names
func (b *BFSTree) Nodes() (names []string) {
	nmap := make(map[string]int)
	for _, e := range b.edges {
		nmap[e.To()] = 0
		nmap[e.From()] = 0
	}

	for n, _ := range nmap {
		names = append(names, n)
	}
	return names
}

// return edges from a given start point
func (b *BFSTree) fromEdges(start string) (res []Edge) {
	for _, e := range b.edges {
		if e.From() == start {
			res = append(res, e)
		}
	}
	return res
}

func (b *BFSTree) FindPath(start string, end string) (path *Path, err error) {
	var paths []*Path

	// create start nodes
	for _, e := range b.fromEdges(start) {
		p := NewPath(e)
		if e.To() == end {
			return p, nil
		}
		paths = append(paths, p)
	}

	for {
		var newPaths []*Path

		if len(paths) == 0 {
			break
		}

		for _, p := range paths {
			children := b.fromEdges(p.Last().To())

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
				np.Append(e)

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
