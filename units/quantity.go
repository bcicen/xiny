package units

import (
	"fmt"
	valuate "github.com/Knetic/govaluate"
)

type ConversionFn func(float64) float64

type Conversion struct {
	From Unit
	To   Unit
	Fn   ConversionFn
}

type Quantity struct {
	Name  string
	Units []Unit
	conv  []Conversion
}

func NewQuantity(name string) *Quantity {
	if _, ok := All[name]; !ok {
		All[name] = &Quantity{Name: name}
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

	q.conv = append(q.conv, Conversion{from, to, fn})
}

// Convert provided value from one unit to another
func (q *Quantity) Convert(val float64, from, to Unit) (newVal float64, err error) {
	fn, err := q.lookup(from, to)
	if err != nil {
		return newVal, err
	}
	return fn(val), nil
}

// find conversion function between two units
func (q *Quantity) lookup(from, to Unit) (ConversionFn, error) {
	for _, c := range q.conv {
		if c.From == from && c.To == to {
			return c.Fn, nil
		}
	}
	return nil, fmt.Errorf("conversion not found")
}
