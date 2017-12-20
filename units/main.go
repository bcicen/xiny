package units

type Quantity struct {
	Name string
}

// reference unit for conversion ratio
func (q Quantity) RefUnit() Unit { return baseUnits[q.Name] }

type Unit struct {
	Name     string
	Symbol   string
	Quantity Quantity
	Ratio    float64 // ratio of one unit to quantity reference unit
}

var (
	Time        = Quantity{"time"}
	Length      = Quantity{"length"}
	Mass        = Quantity{"mass"}
	Temperature = Quantity{"temperature"}
	Frequency   = Quantity{"frequency"}
	Energy      = Quantity{"energy"}
	Power       = Quantity{"power"}
)

var baseUnits = map[string]Unit{
	"time":        Unit{"second", "s", Time, 1.0},
	"length":      Unit{"meter", "m", Length, 1.0},
	"mass":        Unit{"gram", "g", Mass, 1.0},
	"temperature": Unit{"celsius", "c", Temperature, 1.0},
	"frequency":   Unit{"hertz", "hz", Frequency, 1.0},
	"energy":      Unit{"joule", "j", Energy, 1.0},
	"power":       Unit{"watt", "w", Power, 1.0},
}
