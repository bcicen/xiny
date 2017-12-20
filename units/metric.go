package units

import (
	"fmt"
	"math"
)

func makeMagnitudeUnits(name, symbol string, quantity Quantity) (units []Unit) {
	for _, mag := range Magnitudes {
		ratio := 1.0 * math.Pow(10.0, mag.Power)
		name := fmt.Sprintf("%s%s", mag.Prefix, name)
		symbol := fmt.Sprintf("%s%s", mag.Symbol, symbol)
		units = append(units, Unit{name, symbol, quantity, ratio})
	}
	return units
}

type magnitude struct {
	Symbol string
	Prefix string
	Power  float64
}

var Magnitudes = []magnitude{
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
