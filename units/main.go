package units

import (
	"fmt"
)

var (
	All = make(map[string]*Quantity)
)

type Unit struct {
	Name   string
	Symbol string
}

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
