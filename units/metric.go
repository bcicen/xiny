package units

import (
	"fmt"
	"math"
	"strings"
)

var (
	exa   = Magnitude{"E", "exa", 18.0}
	peta  = Magnitude{"P", "peta", 15.0}
	tera  = Magnitude{"T", "tera", 12.0}
	giga  = Magnitude{"G", "giga", 9.0}
	mega  = Magnitude{"M", "mega", 6.0}
	kilo  = Magnitude{"k", "kilo", 3.0}
	hecto = Magnitude{"h", "hecto", 2.0}
	deca  = Magnitude{"da", "deca", 1.0}
	deci  = Magnitude{"d", "deci", -1.0}
	centi = Magnitude{"c", "centi", -2.0}
	milli = Magnitude{"m", "milli", -3.0}
	micro = Magnitude{"μ", "micro", -6.0}
	nano  = Magnitude{"n", "nano", -9.0}
	pico  = Magnitude{"p", "pico", -12.0}
	femto = Magnitude{"f", "femto", -15.0}
	atto  = Magnitude{"a", "atto", -18.0}
)

var (
	magnitudes = []Magnitude{
		exa,
		peta,
		tera,
		giga,
		mega,
		kilo,
		hecto,
		deca,
		deci,
		centi,
		milli,
		micro,
		nano,
		pico,
		femto,
		atto,
	}
)

type Magnitude struct {
	Symbol string
	Prefix string
	Power  float64
}

// Create magnitude unit and conversion given a base unit
func (mag Magnitude) makeUnit(q *Quantity, base Unit, addOpts ...UnitOption) Unit {
	name := fmt.Sprintf("%s%s", mag.Prefix, base.Name)
	symbol := fmt.Sprintf("%s%s", mag.Symbol, base.Symbol)

	// set system to metric by default
	opts := []UnitOption{SI}

	// create prefixed aliases if needed
	for _, alias := range base.aliases {
		magAlias := fmt.Sprintf("%s%s", mag.Prefix, alias)
		opts = append(opts, UnitOptionAliases(magAlias))
	}

	// append any supplmental options
	for _, opt := range addOpts {
		opts = append(opts, opt)
	}

	u := q.NewUnit(name, symbol, opts...)

	// only create conversions to and from base unit
	ratio := 1.0 * math.Pow(10.0, mag.Power)
	q.NewRatioConv(u, base, ratio)

	return u
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
func MakeMagnitudeUnits(q *Quantity, base Unit) {
	for _, mag := range magnitudes {
		name := fmt.Sprintf("%s%s", mag.Prefix, base.Name)
		symbol := fmt.Sprintf("%s%s", mag.Symbol, base.Symbol)
		u := q.NewUnit(name, symbol)

		// only create conversions to and from base unit
		ratio := 1.0 * math.Pow(10.0, mag.Power)
		q.NewRatioConv(u, base, ratio)
	}
}
