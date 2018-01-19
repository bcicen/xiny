package units

var (
	Temp = NewQuantity("temperature")

	Celsius   = Temp.NewUnit("celsius", "C", UnitOptionPlural(false))
	Farenheit = Temp.NewUnit("farenheit", "F", UnitOptionPlural(false))
	Kelvin    = Temp.NewUnit("kelvin", "K", UnitOptionPlural(false))
)

func init() {
	Temp.NewConv(Celsius, Farenheit, "x * 1.8 + 32")
	Temp.NewConv(Farenheit, Celsius, "(x - 32) / 1.8")
	Temp.NewConv(Celsius, Kelvin, "x + 273.15")
	Temp.NewConv(Kelvin, Celsius, "x - 273.15")
}
