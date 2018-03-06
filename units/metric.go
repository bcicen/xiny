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

var mags = map[string]Magnitude{
	"exa":   Magnitude{"E", "exa", 18.0},
	"peta":  Magnitude{"P", "peta", 15.0},
	"tera":  Magnitude{"T", "tera", 12.0},
	"giga":  Magnitude{"G", "giga", 9.0},
	"mega":  Magnitude{"M", "mega", 6.0},
	"kilo":  Magnitude{"k", "kilo", 3.0},
	"hecto": Magnitude{"h", "hecto", 2.0},
	"deca":  Magnitude{"da", "deca", 1.0},
	"deci":  Magnitude{"d", "deci", -1.0},
	"centi": Magnitude{"c", "centi", -2.0},
	"milli": Magnitude{"m", "milli", -3.0},
	"micro": Magnitude{"μ", "micro", -6.0},
	"nano":  Magnitude{"n", "nano", -9.0},
	"pico":  Magnitude{"p", "pico", -12.0},
	"femto": Magnitude{"f", "femto", -15.0},
	"atto":  Magnitude{"a", "atto", -18.0},
}

func Exa(b Unit, o ...UnitOption) Unit   { return mags["exa"].makeUnit(b, o...) }
func Peta(b Unit, o ...UnitOption) Unit  { return mags["peta"].makeUnit(b, o...) }
func Tera(b Unit, o ...UnitOption) Unit  { return mags["tera"].makeUnit(b, o...) }
func Giga(b Unit, o ...UnitOption) Unit  { return mags["giga"].makeUnit(b, o...) }
func Mega(b Unit, o ...UnitOption) Unit  { return mags["mega"].makeUnit(b, o...) }
func Kilo(b Unit, o ...UnitOption) Unit  { return mags["kilo"].makeUnit(b, o...) }
func Hecto(b Unit, o ...UnitOption) Unit { return mags["hecto"].makeUnit(b, o...) }
func Deca(b Unit, o ...UnitOption) Unit  { return mags["deca"].makeUnit(b, o...) }
func Deci(b Unit, o ...UnitOption) Unit  { return mags["deci"].makeUnit(b, o...) }
func Centi(b Unit, o ...UnitOption) Unit { return mags["centi"].makeUnit(b, o...) }
func Milli(b Unit, o ...UnitOption) Unit { return mags["milli"].makeUnit(b, o...) }
func Micro(b Unit, o ...UnitOption) Unit { return mags["micro"].makeUnit(b, o...) }
func Nano(b Unit, o ...UnitOption) Unit  { return mags["nano"].makeUnit(b, o...) }
func Pico(b Unit, o ...UnitOption) Unit  { return mags["pico"].makeUnit(b, o...) }
func Femto(b Unit, o ...UnitOption) Unit { return mags["femto"].makeUnit(b, o...) }
func Atto(b Unit, o ...UnitOption) Unit  { return mags["atto"].makeUnit(b, o...) }

// Create magnitude unit and conversion given a base unit
func (mag Magnitude) makeUnit(base Unit, addOpts ...UnitOption) Unit {
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

	u := base.Quantity.NewUnit(name, symbol, opts...)

	// only create conversions to and from base unit
	ratio := 1.0 * math.Pow(10.0, mag.Power)
	base.Quantity.NewRatioConv(u, base, ratio)

	return u
}

// find and return a magnitude by prefix or symbol
func GetMagnitude(name string) (Magnitude, error) {
	var m Magnitude

	for _, m := range mags {
		if strings.EqualFold(name, m.Symbol) || strings.EqualFold(name, m.Prefix) {
			return m, nil
		}
	}

	return m, fmt.Errorf("magnitude not found")
}

// Create individual units and conversions for all metric mags, given a base unit
func MakeMagnitudeUnits(base Unit) {
	for _, mag := range mags {
		name := fmt.Sprintf("%s%s", mag.Prefix, base.Name)
		symbol := fmt.Sprintf("%s%s", mag.Symbol, base.Symbol)
		u := base.Quantity.NewUnit(name, symbol)

		// only create conversions to and from base unit
		ratio := 1.0 * math.Pow(10.0, mag.Power)
		base.Quantity.NewRatioConv(u, base, ratio)
	}
}
