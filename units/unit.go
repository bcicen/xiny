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
	aliases  []string
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

// Returns all names and symbols this unit may be referred to
func (u Unit) Names() []string {
	names := []string{u.Name}
	if u.Symbol != "" {
		names = append(names, u.Symbol)
	}
	return append(names, u.aliases...)
}

// Return whether this unit can be described with a plural suffix
func (u Unit) Plural() bool { return u.plural }

// Return a Value for this Unit
func (u Unit) MakeValue(v float64) Value { return Value{v, u} }

// Option that may be passed to NewUnit
type UnitOption func(Unit) Unit

// If true, labels for this unit will be created with a plural suffix when appropriate
func UnitOptionPlural(b bool) UnitOption {
	return func(u Unit) Unit {
		u.plural = b
		return u
	}
}

// Additional names, spellings, or symbols that this unit may be referred to as
func UnitOptionAliases(a ...string) UnitOption {
	return func(u Unit) Unit {
		for _, s := range a {
			u.aliases = append(u.aliases, s)
		}
		return u
	}
}
