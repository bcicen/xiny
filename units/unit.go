package units

import (
	"fmt"

	"github.com/bcicen/xiny/log"
)

type Unit struct {
	Name     string
	Symbol   string
	Quantity *Quantity
	plural   bool
}

func NewUnit(name, symbol string, q *Quantity, opts ...UnitOption) Unit {
	if _, ok := UnitMap[name]; ok {
		panic(fmt.Errorf("duplicate unit name: %s", name))
	}

	u := Unit{
		Name:     name,
		Symbol:   symbol,
		Quantity: q,
		plural:   true,
	}

	for _, opt := range opts {
		u = opt(u)
	}

	UnitMap[name] = u
	log.Debugf("loaded unit %s", name)
	return u
}

// Return whether this unit can be described with a plural suffix
func (u Unit) Plural() bool { return u.plural }

// Return a Value for this Unit
func (u Unit) MakeValue(v float64) Value { return Value{v, u} }

type UnitOption func(Unit) Unit

func UnitOptionPlural(b bool) UnitOption {
	return func(u Unit) Unit {
		u.plural = b
		return u
	}
}
