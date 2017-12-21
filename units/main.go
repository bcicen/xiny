package units

import (
	"fmt"
)

var (
	All        = []Unit{}
	Quantities = map[string]Quantity{}
)

type Unit struct {
	Name     string
	Symbol   string
	Quantity Quantity
	Ratio    float64 // ratio of one unit to quantity reference unit
}

func New(name, symbol string, quantity Quantity, ratio float64) {
	u := Unit{name, symbol, quantity, ratio}
	All = append(All, u)
}

// return unit matching name or symbol provided
func Find(s string) (Unit, error) {
	for _, u := range All {
		if u.Name == s || u.Symbol == s {
			return u, nil
		}
	}
	return Unit{}, fmt.Errorf("unit not found")
}

type Quantity struct {
	Name    string
	RefName string // name of reference unit
}

func NewQuantity(name, refunit string) Quantity {
	if _, ok := Quantities[name]; !ok {
		Quantities[name] = Quantity{name, refunit}
	}
	return Quantities[name]
}

// reference unit for conversion ratio
func (q Quantity) RefUnit() (Unit, error) {
	return Find(q.RefName)
}

//var (
//Time        = Quantity{"time", "second"}
//Length      = Quantity{"length", "meter"}
//Temperature = Quantity{"temperature", "celsius"}
//Frequency   = Quantity{"frequency", "hertz"}
//Energy      = Quantity{"energy", "joule"}
//Power       = Quantity{"power", "watt"}
//)
