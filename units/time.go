package units

var (
	Time = NewQuantity("time")

	Second      = Time.NewUnit("second", "s")
	ExaSecond   = exa.makeUnit(Time, Second)
	PetaSecond  = peta.makeUnit(Time, Second)
	TeraSecond  = tera.makeUnit(Time, Second)
	GigaSecond  = giga.makeUnit(Time, Second)
	MegaSecond  = mega.makeUnit(Time, Second)
	KiloSecond  = kilo.makeUnit(Time, Second)
	HectoSecond = hecto.makeUnit(Time, Second)
	DecaSecond  = deca.makeUnit(Time, Second)
	DeciSecond  = deci.makeUnit(Time, Second)
	CentiSecond = centi.makeUnit(Time, Second)
	MilliSecond = milli.makeUnit(Time, Second)
	MicroSecond = micro.makeUnit(Time, Second)
	NanoSecond  = nano.makeUnit(Time, Second)
	PicoSecond  = pico.makeUnit(Time, Second)
	FemtoSecond = femto.makeUnit(Time, Second)
	AttoSecond  = atto.makeUnit(Time, Second)

	Minute = Time.NewUnit("minute", "min")
	Hour   = Time.NewUnit("hour", "hr")
	Day    = Time.NewUnit("day", "d")
	Month  = Time.NewUnit("month", "")
	Year   = Time.NewUnit("year", "yr")

	Decade     = Time.NewUnit("decade", "")
	Century    = Time.NewUnit("century", "")
	Millennium = Time.NewUnit("millennium", "")

	// more esoteric time units
	PlanckTime = Time.NewUnit("planck time", "𝑡ₚ")
	Fortnight  = Time.NewUnit("fortnight", "")
	Score      = Time.NewUnit("score", "")
)

func init() {
	Time.NewRatioConv(Minute, Second, 60.0)
	Time.NewRatioConv(Hour, Second, 3600.0)
	Time.NewRatioConv(Day, Hour, 24.0)
	Time.NewRatioConv(Month, Day, 30.0)
	Time.NewRatioConv(Year, Day, 365.25)

	Time.NewRatioConv(Decade, Year, 10.0)
	Time.NewRatioConv(Century, Year, 100.0)
	Time.NewRatioConv(Millennium, Year, 1000.0)

	Time.NewRatioConv(PlanckTime, Second, 5.39e-44)
	Time.NewRatioConv(Fortnight, Day, 14)
	Time.NewRatioConv(Score, Year, 20.0)
}
