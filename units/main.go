package units

import (
	"fmt"
)

var (
	UnitMap = make(map[string]Unit)
)

type Value struct {
	Val  float64
	Unit Unit
}

func (v Value) Fmt(opts FmtOptions) string { return v.Unit.Quantity.Formatter(v, opts) }

// Convert this Value to another Unit, returning the new Value
func (v Value) Convert(to Unit) (newVal Value, err error) {
	// allow converting to same unit
	if v.Unit == to {
		return v, nil
	}

	fns, err := v.Unit.Quantity.Resolve(v.Unit, to)
	if err != nil {
		return newVal, err
	}

	fVal := v.Val
	for _, fn := range fns {
		fVal = fn(fVal)
	}

	return Value{fVal, to}, nil
}

type Unit struct {
	Name     string
	Symbol   string
	Quantity *Quantity
}

// Return a Value for this Unit
func (u Unit) MakeValue(v float64) Value { return Value{v, u} }

// Find Unit matching name or symbol provided
func Find(s string) (Unit, error) {
	for _, u := range UnitMap {
		if u.Name == s || u.Symbol == s {
			return u, nil
		}
	}
	return Unit{}, fmt.Errorf("unit \"%s\"not found", s)
}
