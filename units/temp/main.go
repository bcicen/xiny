package temp

import (
	"github.com/bcicen/xiny/units"
)

var (
	Temp = units.NewQuantity("temperature")
)

func init() {
	C := Temp.NewUnit("celsius", "C")
	F := Temp.NewUnit("farenheit", "F")
	K := Temp.NewUnit("kelvin", "K")

	Temp.NewConv(C, F, "x * 1.8 + 32")
	Temp.NewConv(F, C, "(x - 32) / 1.8")
	Temp.NewConv(C, K, "(x - 32) / 1.8")
}
