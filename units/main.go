package units

import (
	"fmt"
	"strings"
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

	// first try case-sensitive match
	for _, u := range UnitMap {
		if matchUnitName(s, u, true) {
			return u, nil
		}
	}

	// then case-insensitive
	for _, u := range UnitMap {
		if matchUnitName(s, u, false) {
			return u, nil
		}
	}

	return Unit{}, fmt.Errorf("unit \"%s\" not found", s)
}

func matchUnitName(s string, u Unit, matchCase bool) bool {
	if u.Name == s || u.Symbol == s {
		return true
	}

	if !matchCase {
		ls := strings.ToLower(s)
		if strings.ToLower(u.Name) == ls || strings.ToLower(u.Symbol) == ls {
			return true
		}
	}

	return false
}
