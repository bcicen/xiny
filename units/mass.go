package units

var Mass = NewQuantity("mass")

func init() {
	gram := Mass.NewUnit("gram", "g")
	MakeMagnitudeUnits(Mass, gram)

	// imperial units
	grain := Mass.NewUnit("grain", "gr")
	drachm := Mass.NewUnit("drachm", "dr")
	ounce := Mass.NewUnit("ounce", "oz")
	pound := Mass.NewUnit("pound", "lb")
	stone := Mass.NewUnit("stone", "st")
	ton := Mass.NewUnit("ton", "t")
	slug := Mass.NewUnit("slug", "slug")

	Mass.NewRatioConv(grain, gram, 0.06479891)
	Mass.NewRatioConv(drachm, gram, 1.7718451953125)
	Mass.NewRatioConv(ounce, gram, 28.349523125)
	Mass.NewRatioConv(pound, gram, 453.59237)
	Mass.NewRatioConv(stone, gram, 6350.29318)
	Mass.NewRatioConv(ton, gram, 1016046.9088)
	Mass.NewRatioConv(slug, gram, 14593.90294)
}
