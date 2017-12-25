package units

import (
	"fmt"
	valuate "github.com/Knetic/govaluate"
)

type Conversion struct {
	From Unit
	To   Unit
}

type ConversionFn func(float64) float64

type ConversionMap map[Conversion]ConversionFn

func (cm ConversionMap) Add(from, to Unit, fn ConversionFn) {
	cm[Conversion{from, to}] = fn
}

func (cm ConversionMap) Lookup(from, to Unit) (ConversionFn, error) {
	for k, v := range cm {
		if k.From == from && k.To == to {
			return v, nil
		}
	}
	return nil, fmt.Errorf("conversion not found")
}

type Quantity struct {
	Name  string
	Units []Unit
	cmap  ConversionMap
}

func NewQuantity(name string) *Quantity {
	if _, ok := All[name]; !ok {
		All[name] = &Quantity{
			Name: name,
			cmap: make(ConversionMap),
		}
	}
	fmt.Printf("added new quantity %s\n", name)
	return All[name]
}

// Create a new Unit within this quantity and return it
func (q *Quantity) NewUnit(name, symbol string) Unit {
	u := Unit{name, symbol}
	q.Units = append(q.Units, u)
	fmt.Printf("added new unit %s\n", name)
	return u
}

// return unit matching name or symbol provided
func (q *Quantity) FindUnit(s string) (Unit, error) {
	for _, u := range q.Units {
		if u.Name == s || u.Symbol == s {
			return u, nil
		}
	}
	return Unit{}, fmt.Errorf("unit \"%s\"not found", s)
}

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

	q.cmap.Add(from, to, fn)
}

func (q *Quantity) Conv(from, to Unit) ConversionFn {
	fn, err := q.cmap.Lookup(from, to)
	if err != nil {
		panic(err)
	}
	return fn
}
