package units

import (
	"fmt"
	"strconv"
)

type Value struct {
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
