package units

var Time = NewQuantity("time", DefaultFormatter)

func init() {
	second := Time.NewUnit("second", "s")
	MakeMagnitudeUnits(Time, second)

	minute := Time.NewUnit("minute", "min")
	hour := Time.NewUnit("hour", "hr")
	day := Time.NewUnit("day", "d")
	year := Time.NewUnit("year", "yr")

	Time.NewRatioConv(minute, second, 60.0)
	Time.NewRatioConv(hour, second, 3600.0)
	Time.NewRatioConv(day, hour, 24.0)
	Time.NewRatioConv(year, day, 365.25)
}
