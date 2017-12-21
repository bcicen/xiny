package mass

import (
	"github.com/bcicen/xiny/units"
)

var Mass = units.NewQuantity("mass", "gram")

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
