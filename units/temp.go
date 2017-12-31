package units

import (
	"fmt"
	"strconv"
)

var Temp = NewQuantity("temperature", TempFormatter)

// TempFormatter is a ValueFormatter implementation
func TempFormatter(v Value, opts FmtOptions) string {
	vstr := strconv.FormatFloat(v.Val, 'f', opts.Precision, 64)

	if opts.Short {
		return fmt.Sprintf("%s %s", vstr, v.Unit.Symbol)
	}
	return fmt.Sprintf("%s %s", vstr, v.Unit.Name)
}

func init() {
	C := Temp.NewUnit("celsius", "C")
	F := Temp.NewUnit("farenheit", "F")
	K := Temp.NewUnit("kelvin", "K")

	Temp.NewConv(C, F, "x * 1.8 + 32")
	Temp.NewConv(F, C, "(x - 32) / 1.8")
	Temp.NewConv(C, K, "(x - 32) / 1.8")
}
