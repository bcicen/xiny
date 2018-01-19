package units

var (
	Length = NewQuantity("length")

	// metric
	Meter      = Length.NewUnit("meter", "m", UnitOptionAliases("metre"))
	ExaMeter   = exa.makeUnit(Length, Meter)
	PetaMeter  = peta.makeUnit(Length, Meter)
	TeraMeter  = tera.makeUnit(Length, Meter)
	GigaMeter  = giga.makeUnit(Length, Meter)
	MegaMeter  = mega.makeUnit(Length, Meter)
	KiloMeter  = kilo.makeUnit(Length, Meter)
	HectoMeter = hecto.makeUnit(Length, Meter)
	DecaMeter  = deca.makeUnit(Length, Meter)
	DeciMeter  = deci.makeUnit(Length, Meter)
	CentiMeter = centi.makeUnit(Length, Meter)
	MilliMeter = milli.makeUnit(Length, Meter)
	MicroMeter = micro.makeUnit(Length, Meter)
	NanoMeter  = nano.makeUnit(Length, Meter)
	PicoMeter  = pico.makeUnit(Length, Meter)
	FemtoMeter = femto.makeUnit(Length, Meter)
	AttoMeter  = atto.makeUnit(Length, Meter)

	Inch    = Length.NewUnit("inch", "in")
	Feet    = Length.NewUnit("feet", "ft")
	Yard    = Length.NewUnit("yard", "yd")
	Mile    = Length.NewUnit("mile", "ml")
	League  = Length.NewUnit("league", "lea")
	Furlong = Length.NewUnit("furlong", "fur")
)

func init() {
	Length.NewRatioConv(Inch, Meter, 0.0254)
	Length.NewRatioConv(Feet, Meter, 0.3048)
	Length.NewRatioConv(Yard, Meter, 0.9144)
	Length.NewRatioConv(Mile, Meter, 1609.344)
	Length.NewRatioConv(League, Meter, 4828.032)
	Length.NewRatioConv(Furlong, Meter, 201.168)
}
