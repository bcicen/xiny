package units

var (
	Mass = NewQuantity("mass")

	// metric
	Gram      = Mass.NewUnit("gram", "g")
	ExaGram   = exa.makeUnit(Mass, Gram)
	PetaGram  = peta.makeUnit(Mass, Gram)
	TeraGram  = tera.makeUnit(Mass, Gram)
	GigaGram  = giga.makeUnit(Mass, Gram)
	MegaGram  = mega.makeUnit(Mass, Gram)
	KiloGram  = kilo.makeUnit(Mass, Gram)
	HectoGram = hecto.makeUnit(Mass, Gram)
	DecaGram  = deca.makeUnit(Mass, Gram)
	DeciGram  = deci.makeUnit(Mass, Gram)
	CentiGram = centi.makeUnit(Mass, Gram)
	MilliGram = milli.makeUnit(Mass, Gram)
	MicroGram = micro.makeUnit(Mass, Gram)
	NanoGram  = nano.makeUnit(Mass, Gram)
	PicoGram  = pico.makeUnit(Mass, Gram)
	FemtoGram = femto.makeUnit(Mass, Gram)
	AttoGram  = atto.makeUnit(Mass, Gram)

	// imperial
	Grain  = Mass.NewUnit("grain", "gr", BI)
	Drachm = Mass.NewUnit("drachm", "dr", BI)
	Ounce  = Mass.NewUnit("ounce", "oz", BI)
	Pound  = Mass.NewUnit("pound", "lb", BI)
	Stone  = Mass.NewUnit("stone", "st", BI)
	Ton    = Mass.NewUnit("ton", "t", BI)
	Slug   = Mass.NewUnit("slug", "", BI)
)

func init() {
	Mass.NewRatioConv(Grain, Gram, 0.06479891)
	Mass.NewRatioConv(Drachm, Gram, 1.7718451953125)
	Mass.NewRatioConv(Ounce, Gram, 28.349523125)
	Mass.NewRatioConv(Pound, Gram, 453.59237)
	Mass.NewRatioConv(Stone, Gram, 6350.29318)
	Mass.NewRatioConv(Ton, Gram, 1016046.9088)
	Mass.NewRatioConv(Slug, Gram, 14593.90294)
}
