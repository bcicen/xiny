package units

import (
	"fmt"
	"sort"
	"strings"
)

var (
	UnitMap = make(map[string]Unit)
)

// Return sorted list of all Unit names and symbols
func Names() (a []string) {
	for _, u := range UnitMap {
		a = append(a, u.Name)
		a = append(a, u.Symbol)
	}
	sort.Strings(a)
	return a
}

// Find Unit matching name or symbol provided, panicking on failure
func MustFind(s string) Unit {
	u, err := Find(s)
	if err != nil {
		panic(err)
	}
	return u
}

// Find Unit matching name or symbol provided
func Find(s string) (Unit, error) {

	// first try case-sensitive match
	for _, u := range UnitMap {
		if matchUnitName(s, u, true) {
			return u, nil
		}
	}

	// then case-insensitive
	for _, u := range UnitMap {
		if matchUnitName(s, u, false) {
			return u, nil
		}
	}

	// finally, try stripping plural suffix
	if strings.HasSuffix(s, "s") || strings.HasSuffix(s, "S") {
		s = strings.TrimSuffix(s, "s")
		s = strings.TrimSuffix(s, "S")
		for _, u := range UnitMap {
			if matchUnitName(s, u, false) {
				return u, nil
			}
		}
	}

	return Unit{}, fmt.Errorf("unit \"%s\" not found", s)
}

func matchUnitName(s string, u Unit, matchCase bool) bool {
	if u.Name == s || u.Symbol == s {
		return true
	}

	if !matchCase {
		if strings.EqualFold(s, u.Name) || strings.EqualFold(s, u.Symbol) {
			return true
		}
	}

	return false
}
