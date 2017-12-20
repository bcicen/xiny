package units

var All = massUnits()

func massUnits() (all []Unit) {
	all = append(all, makeMagnitudeUnits("gram", "g", Mass)...)
	all = append(all, imperialMassUnits...)

	return all
}

var imperialMassUnits = []Unit{
	{"grain", "gr", Mass, 0.06479891},
	{"drachm", "dr", Mass, 1.7718451953125},
	{"ounce", "oz", Mass, 28.349523125},
	{"pound", "lb", Mass, 453.59237},
	{"stone", "st", Mass, 6350.29318},
	{"ton", "t", Mass, 1016046.9088},
	{"slug", "slug", Mass, 14593.90294},
}
