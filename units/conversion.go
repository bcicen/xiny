package units

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	valuate "github.com/Knetic/govaluate"
)

type ConversionFn func(float64) float64

type Conversion struct {
	from    Unit
	to      Unit
	Fn      ConversionFn
	Formula string
}

// Conversion implements bfstree.Edge interface
func (c Conversion) To() string   { return c.to.Name }
func (c Conversion) From() string { return c.from.Name }

func (c Conversion) String() string { return c.Formula }

func newConversion(from, to Unit, formula string) Conversion {
	expr, err := valuate.NewEvaluableExpression(formula)
	if err != nil {
		panic(err)
	}

	// create conversion function
	fn := func(x float64) float64 {
		params := make(map[string]interface{})
		params["x"] = x

		res, err := expr.Evaluate(params)
		if err != nil {
			panic(err)
		}
		return res.(float64)
	}

	return Conversion{from, to, fn, fmtFormula(formula)}
}

// Replace float in formula string with scientific notation where necessary
func fmtFormula(s string) string {
	re := regexp.MustCompile("(-?[0-9.]+)")
	for _, match := range re.FindAllString(s, -1) {
		f, err := strconv.ParseFloat(match, 64)
		if err != nil {
			return s
		}
		s = strings.Replace(s, match, fmt.Sprintf("%g", f), 1)
	}
	return s
}
