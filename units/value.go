package units

import (
	"fmt"
	"strconv"
	"strings"
)

type FmtOptions struct {
	Short     bool // if false, use unit shortname or symbol
	Precision int  // precision to truncate value
}

type Value struct {
	Val  float64
	Unit Unit
}

// Convert this Value to another Unit, returning the new Value
func (v Value) Convert(to Unit) (newVal Value, err error) {
	// allow converting to same unit
	if v.Unit.Name == to.Name {
		return v, nil
	}

	fns, err := v.Unit.Quantity.Resolve(v.Unit, to)
	if err != nil {
		return newVal, err
	}

	fVal := v.Val
	for _, fn := range fns {
		fVal = fn(fVal)
	}

	return Value{fVal, to}, nil
}

func (v Value) Fmt(opts FmtOptions) string {
	var label string

	if opts.Short {
		label = v.Unit.Symbol
	} else {
		label = v.Unit.Name
		// make label plural if needed
		if v.Unit.plural && v.Val > 1.0 {
			label = fmt.Sprintf("%ss", label)
		}
	}

	vstr := strconv.FormatFloat(v.Val, 'f', opts.Precision, 64)
	vstr = trimTrailing(vstr)

	return fmt.Sprintf("%s %s", vstr, label)
}

// Trim trailing zeros from string
func trimTrailing(s string) string {
	s = strings.TrimRight(s, "0")
	if s[len(s)-1] == '.' {
		s = strings.TrimSuffix(s, ".")
	}
	return s
}
