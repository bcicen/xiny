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

// ValueFormatter creates human-readable strings for a Unit value
type ValueFormatter func(Value, FmtOptions) string

func DefaultFormatter(v Value, opts FmtOptions) string {
	var label string

	if opts.Short {
		label = v.Unit.Symbol
	} else {
		label = v.Unit.Name
		// make label plural if needed
		if v.Val > 1.0 {
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
