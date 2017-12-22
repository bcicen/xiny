package units

import (
	"fmt"
)

type Value struct {
	val  float64
	unit Unit
}

func NewValue(val float64, unit Unit) *Value { return &Value{val, unit} }

func (v *Value) ToUnit(u Unit) {
	fmt.Printf("%.6f / (%.6f / %.6f)\n", v.val, u.Ratio, v.unit.Ratio)
	v.val = v.val / (u.Ratio / v.unit.Ratio)
	v.unit = u
}

func (v *Value) LongString() string {
	return fmt.Sprintf("%.6g %s", v.val, v.unit.Name)
}

func (v *Value) String() string {
	return fmt.Sprintf("%.6g%s", v.val, v.unit.Symbol)
}
