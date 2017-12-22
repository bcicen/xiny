package units

var Quantities = map[string]Quantity{}

type Quantity struct {
	Name    string
	RefName string // name of reference unit
}

func NewQuantity(name, refunit string) Quantity {
	if _, ok := Quantities[name]; !ok {
		Quantities[name] = Quantity{name, refunit}
	}
	return Quantities[name]
}

// reference unit for conversion ratio
func (q Quantity) RefUnit() (Unit, error) {
	return Find(q.RefName)
}
