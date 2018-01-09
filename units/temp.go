package units

var Temp = NewQuantity("temperature")

func init() {
	C := Temp.NewUnit("celsius", "C", UnitOptionPlural(false))
	F := Temp.NewUnit("farenheit", "F", UnitOptionPlural(false))
	K := Temp.NewUnit("kelvin", "K", UnitOptionPlural(false))

	Temp.NewConv(C, F, "x * 1.8 + 32")
	Temp.NewConv(F, C, "(x - 32) / 1.8")
	Temp.NewConv(C, K, "x + 273.15")
	Temp.NewConv(K, C, "x - 273.15")
}
