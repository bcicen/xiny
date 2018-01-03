package units

var Length = NewQuantity("length", DefaultFormatter)

func init() {
	metre := Length.NewUnit("metre", "m")
	MakeMagnitudeUnits(Length, metre)

	inch := Length.NewUnit("inch", "in")
	feet := Length.NewUnit("feet", "ft")
	yard := Length.NewUnit("yard", "yd")
	mile := Length.NewUnit("mile", "ml")
	league := Length.NewUnit("league", "lea")
	furlong := Length.NewUnit("furlong", "fur")

	Length.NewRatioConv(inch, metre, 0.0254)
	Length.NewRatioConv(feet, metre, 0.3048)
	Length.NewRatioConv(yard, metre, 0.9144)
	Length.NewRatioConv(mile, metre, 1609.344)
	Length.NewRatioConv(league, metre, 4828.032)
	Length.NewRatioConv(furlong, metre, 201.168)
}
