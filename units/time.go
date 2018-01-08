package units

var Time = NewQuantity("time")

func init() {
	second := Time.NewUnit("second", "s")
	MakeMagnitudeUnits(Time, second)

	minute := Time.NewUnit("minute", "min")
	hour := Time.NewUnit("hour", "hr")
	day := Time.NewUnit("day", "d")
	month := Time.NewUnit("month", "")
	year := Time.NewUnit("year", "yr")

	Time.NewRatioConv(minute, second, 60.0)
	Time.NewRatioConv(hour, second, 3600.0)
	Time.NewRatioConv(day, hour, 24.0)
	Time.NewRatioConv(month, day, 30.0)
	Time.NewRatioConv(year, day, 365.25)

	decade := Time.NewUnit("decade", "")
	century := Time.NewUnit("century", "")
	millennium := Time.NewUnit("millennium", "")

	Time.NewRatioConv(decade, year, 10.0)
	Time.NewRatioConv(century, year, 100.0)
	Time.NewRatioConv(millennium, year, 1000.0)

	// more esoteric time units
	planckTime := Time.NewUnit("planck time", "tP")
	Time.NewRatioConv(planckTime, second, 5.39e-44)

	fortnight := Time.NewUnit("fortnight", "")
	Time.NewRatioConv(fortnight, day, 14)

	score := Time.NewUnit("score", "")
	Time.NewRatioConv(score, year, 20.0)
}
