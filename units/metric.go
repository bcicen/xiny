package units

import (
	"fmt"
	"math"
	"strings"
)

type Magnitude struct {
	Symbol string
	Prefix string
	Power  float64
}

var magnitudes = []Magnitude{
	{"E", "exa", 18.0},
	{"P", "peta", 15.0},
	{"T", "tera", 12.0},
	{"G", "giga", 9.0},
	{"M", "mega", 6.0},
	{"k", "kilo", 3.0},
	{"h", "hecto", 2.0},
	{"da", "deca", 1.0},
	//{"", "", 0.0},
	{"d", "deci", -1.0},
	{"c", "centi", -2.0},
	{"m", "milli", -3.0},
	{"μ", "micro", -6.0},
	{"n", "nano", -9.0},
	{"p", "pico", -12.0},
	{"f", "femto", -15.0},
	{"a", "atto", -18.0},
}

// find and return a magnitude by prefix or symbol
func GetMagnitude(name string) (Magnitude, error) {
	var m Magnitude

	for _, m := range magnitudes {
		if strings.EqualFold(name, m.Symbol) || strings.EqualFold(name, m.Prefix) {
			return m, nil
		}
	}

	return m, fmt.Errorf("magnitude not found")
}

// Create individual units and conversions for all metric magnitudes, given a base unit
func MakeMagnitudeUnits(q *Quantity, baseUnit Unit) {
	for _, mag := range magnitudes {
		name := fmt.Sprintf("%s%s", mag.Prefix, baseUnit.Name)
		symbol := fmt.Sprintf("%s%s", mag.Symbol, baseUnit.Symbol)
		u := q.NewUnit(name, symbol)

		// only create conversions to and from base unit
		ratio := 1.0 * math.Pow(10.0, mag.Power)
		q.NewRatioConv(u, baseUnit, ratio)
	}
}
