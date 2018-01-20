package units

var (
	Volume = NewQuantity("volume")

	// metric
	Liter      = Volume.NewUnit("liter", "l", UnitOptionAliases("litre"))
	ExaLiter   = exa.makeUnit(Volume, Liter)
	PetaLiter  = peta.makeUnit(Volume, Liter)
	TeraLiter  = tera.makeUnit(Volume, Liter)
	GigaLiter  = giga.makeUnit(Volume, Liter)
	MegaLiter  = mega.makeUnit(Volume, Liter)
	KiloLiter  = kilo.makeUnit(Volume, Liter)
	HectoLiter = hecto.makeUnit(Volume, Liter)
	DecaLiter  = deca.makeUnit(Volume, Liter)
	DeciLiter  = deci.makeUnit(Volume, Liter)
	CentiLiter = centi.makeUnit(Volume, Liter)
	MilliLiter = milli.makeUnit(Volume, Liter)
	MicroLiter = micro.makeUnit(Volume, Liter)
	NanoLiter  = nano.makeUnit(Volume, Liter)
	PicoLiter  = pico.makeUnit(Volume, Liter)
	FemtoLiter = femto.makeUnit(Volume, Liter)
	AttoLiter  = atto.makeUnit(Volume, Liter)

	// imperial
	Quart      = Volume.NewUnit("quart", "qt")
	Pint       = Volume.NewUnit("pint", "pt")
	Gallon     = Volume.NewUnit("gallon", "gal")
	FluidOunce = Volume.NewUnit("fluid ounce", "fl oz")

	// US
	FluidQuart          = Volume.NewUnit("fluid quart", "")
	FluidPint           = Volume.NewUnit("fluid pint", "")
	FluidGallon         = Volume.NewUnit("fluid gallon", "")
	CustomaryFluidOunce = Volume.NewUnit("customary fluid ounce", "")
)

func init() {
	Volume.NewRatioConv(Quart, Liter, 1.1365225)
	Volume.NewRatioConv(Pint, Liter, 0.56826125)
	Volume.NewRatioConv(Gallon, Liter, 4.54609)
	Volume.NewRatioConv(FluidOunce, MilliLiter, 28.4130625)

	Volume.NewRatioConv(FluidQuart, Liter, 0.946352946)
	Volume.NewRatioConv(FluidPint, Liter, 0.473176473)
	Volume.NewRatioConv(FluidGallon, Liter, 3.785411784)
	Volume.NewRatioConv(CustomaryFluidOunce, MilliLiter, 29.5735295625)
}
