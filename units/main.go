package units

import (
	"fmt"
)

var (
	All = make(map[string]*Quantity)
)

type Value struct {
	Val  float64
	Unit Unit
}

type Unit struct {
	Name   string
	Symbol string
}

// Return a Value for this Unit
func (u Unit) MakeValue(v float64) Value { return Value{v, u} }

// return unit matching name or symbol provided
func Find(s string) (*Quantity, Unit, error) {
	for _, q := range All {
		u, err := q.FindUnit(s)
		if err == nil {
			return q, u, nil
		}
	}
	panic(fmt.Errorf("oops"))
}
