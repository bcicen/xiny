package main

import (
	"fmt"
	"math"
)

type Magnitude struct {
	Symbol string  `json:"symbol"`
	Prefix string  `json:"prefix"`
	Power  float64 `json:"power"`
}

var (
	exa   = Magnitude{"E", "exa", 18.0}
	peta  = Magnitude{"P", "peta", 15.0}
	tera  = Magnitude{"T", "tera", 12.0}
	giga  = Magnitude{"G", "giga", 9.0}
	mega  = Magnitude{"M", "mega", 6.0}
	kilo  = Magnitude{"k", "kilo", 3.0}
	hecto = Magnitude{"h", "hecto", 2.0}
	deca  = Magnitude{"da", "deca", 1.0}
	none  = Magnitude{"", "", 0.0}
	deci  = Magnitude{"d", "deci", -1.0}
	centi = Magnitude{"c", "centi", -2.0}
	milli = Magnitude{"m", "milli", -3.0}
	micro = Magnitude{"μ", "micro", -6.0}
	nano  = Magnitude{"n", "nano", -9.0}
	pico  = Magnitude{"p", "pico", -12.0}
	femto = Magnitude{"f", "femto", -15.0}
	atto  = Magnitude{"a", "atto", -18.0}
)

type Value struct {
	v    float64
	mag  Magnitude
	unit string
}

func (v *Value) String() string {
	return fmt.Sprintf("%.6g %s%s", v.v, v.mag.Prefix, v.unit)
}

func (v *Value) ToMagnitude(newMag Magnitude) {
	factor := math.Pow(10.0, (v.mag.Power-newMag.Power)*-1)
	//fmt.Printf("old: %f new: %f factor: %f\n", v.mag.Power, newMag.Power, factor)
	v.v = v.v / factor
	v.mag = newMag
}

func main() {
	val := &Value{1.0, none, "gram"}
	fmt.Println(val)

	list := []Magnitude{deci, centi, micro, nano, pico, femto, atto, none, deca, hecto, kilo, mega, giga, tera, peta, exa, atto, none}
	for _, m := range list {
		val.ToMagnitude(m)
		fmt.Println(val)
	}
	//val.ToMagnitude(kilo)
	//fmt.Println(val)
	//val.ToMagnitude(none)
	//fmt.Println(val)
	//val.ToMagnitude(milli)
	//fmt.Println(val)
	//val.ToMagnitude(mega)
	//fmt.Println(val)
	//val.ToMagnitude(none)
	//fmt.Println(val)
}
