package mass

import (
	"github.com/bcicen/xiny/units"
)

var Mass = units.NewQuantity("mass", "gram")

type MassValue struct {
	val  float64
	unit Unit
}

func NewValue(val float64, unit Unit) *Value { return &Value{val, unit} }

func (v *Value) String() string { return v.FString(false, 3) }

func (v *Value) FString(long bool, precision int) string {
	label := v.unit.Symbol
	if long {
		label = v.unit.Name
	}

	if plural(v.val) {
		label = fmt.Sprintf("%ss", label)
	}

	vstr := strconv.FormatFloat(v.val, 'f', precision, 64)

	return fmt.Sprintf("%s %s", vstr, label)
}

func (v *Value) ToUnit(u Unit) {
	fmt.Printf("%.6f / (%.6f / %.6f)\n", v.val, u.Ratio, v.unit.Ratio)
	v.val = v.val / (u.Ratio / v.unit.Ratio)
	v.unit = u
}

func plural(v float64) bool { return v > 1.0 }

func init() {
	units.NewMagnitudeUnits("gram", "g", Mass)

	// imperial units
	units.New("grain", "gr", Mass, 0.06479891)
	units.New("drachm", "dr", Mass, 1.7718451953125)
	units.New("ounce", "oz", Mass, 28.349523125)
	units.New("pound", "lb", Mass, 453.59237)
	units.New("stone", "st", Mass, 6350.29318)
	units.New("ton", "t", Mass, 1016046.9088)
	units.New("slug", "slug", Mass, 14593.90294)
}
