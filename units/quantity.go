package units

import (
	"fmt"

	valuate "github.com/Knetic/govaluate"
	"github.com/bcicen/xiny/bfstree"
	"github.com/bcicen/xiny/log"
)

type ConversionFn func(float64) float64

type Conversion struct {
	from    Unit
	to      Unit
	Fn      ConversionFn
	Formula string
}

// Conversion implements bfstree.Edge interface
func (c Conversion) To() string   { return c.to.Name }
func (c Conversion) From() string { return c.from.Name }

type Quantity struct {
	Name      string
	Formatter ValueFormatter
	conv      []Conversion
}

func NewQuantity(name string, formatter ValueFormatter) *Quantity {
	q := &Quantity{
		Name:      name,
		Formatter: formatter,
	}
	log.Debugf("loaded quantity %s", name)
	return q
}

// Create a new Unit within this quantity and return it
func (q *Quantity) NewUnit(name, symbol string) Unit {
	if _, ok := UnitMap[name]; ok {
		panic(fmt.Errorf("duplicate unit name: %s", name))
	}
	u := Unit{name, symbol, q}
	UnitMap[name] = u
	log.Debugf("loaded unit %s", name)
	return u
}

// Create a conversion and the inverse, given a ratio of from Unit
// in to Unit
func (q *Quantity) NewRatioConv(from, to Unit, ratio float64) {
	q.NewConv(from, to, fmt.Sprintf("x * %.12f", ratio))
	q.NewConv(to, from, fmt.Sprintf("x / %.12f", ratio))
}

// Create a new conversion from one unit to another
func (q *Quantity) NewConv(from, to Unit, formula string) {
	expr, err := valuate.NewEvaluableExpression(formula)
	if err != nil {
		panic(err)
	}

	// create conversion function
	fn := func(x float64) float64 {
		params := make(map[string]interface{})
		params["x"] = x

		res, err := expr.Evaluate(params)
		if err != nil {
			panic(err)
		}
		return res.(float64)
	}

	log.Debugf("loaded conversion %s -> %s", from.Name, to.Name)
	q.conv = append(q.conv, Conversion{from, to, fn, formula})
}

// Resolve a path of one or more conversions between two units
func (q *Quantity) Resolve(from, to Unit) (fns []ConversionFn, err error) {
	tree := bfstree.NewBFSTree()
	for _, cnv := range q.conv {
		tree.AddEdge(cnv)
	}

	path, err := tree.FindPath(from.Name, to.Name)
	if err != nil {
		return fns, err
	}

	for _, edge := range path.Edges() {
		conv, err := q.lookup(edge.From(), edge.To())
		if err != nil {
			return fns, err
		}
		log.Infof("%s -> %s [%s]", edge.From(), edge.To(), conv.Formula)
		fns = append(fns, conv.Fn)
	}

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
