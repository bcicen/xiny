package units

import (
	"fmt"
	"strings"

	"github.com/bcicen/bfstree"
	"github.com/bcicen/xiny/log"
)

var (
	QuantityMap = make(map[string]*Quantity)
)

type Quantity struct {
	Name string
	conv []Conversion
}

func NewQuantity(name string) *Quantity {
	q := &Quantity{Name: name}
	QuantityMap[name] = q
	log.Debugf("loaded quantity %s", name)
	return q
}

// Create a new Unit within this quantity and return it
func (q *Quantity) NewUnit(name, symbol string, opts ...UnitOption) Unit {
	return NewUnit(name, symbol, q, opts...)
}

// Create a conversion and the inverse, given a ratio of from Unit
// in to Unit
func (q *Quantity) NewRatioConv(from, to Unit, ratio float64) {
	ratioStr := fmt.Sprintf("%.62f", ratio)
	q.NewConv(from, to, fmt.Sprintf("x * %s", ratioStr))
	q.NewConv(to, from, fmt.Sprintf("x / %s", ratioStr))
}

// Create a new conversion from one unit to another, within this quantity
func (q *Quantity) NewConv(from, to Unit, formula string) {
	conv := newConversion(from, to, formula)
	log.Debugf("loaded conversion %s -> %s", from.Name, to.Name)
	q.conv = append(q.conv, conv)
}

// Resolve a path of one or more conversions between two units
func (q *Quantity) Resolve(from, to Unit) (fns []ConversionFn, err error) {
	tree := bfstree.New()
	for _, cnv := range q.conv {
		tree.AddEdge(cnv)
	}

	path, err := tree.FindPath(from.Name, to.Name)
	if err != nil {
		return fns, err
	}

	formula := ""
	for _, edge := range path.Edges() {
		conv, err := q.lookup(edge.From(), edge.To())
		if err != nil {
			return fns, err
		}
		if formula != "" {
			formula = fmt.Sprintf("(%s)", strings.Replace(conv.Formula, "x", formula, 1))
		} else {
			formula = fmt.Sprintf("(%s)", conv.Formula)
		}
		log.Debugf("%s -> %s: %s", edge.From(), edge.To(), conv.Formula)
		fns = append(fns, conv.Fn)
	}
	log.Infof("%s -> %s: %s", from.Name, to.Name, formula)

	return fns, nil
}

// find conversion function between two units
func (q *Quantity) lookup(from, to string) (c Conversion, err error) {
	for _, c := range q.conv {
		if c.From() == from && c.To() == to {
			return c, nil
		}
	}
	return c, fmt.Errorf("conversion not found")
}
