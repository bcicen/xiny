package main

import (
	"math"
	"testing"

	"github.com/bcicen/go-units"
)

// aggregate all unit names, aliases, etc
func aggrNames() (a []string) {
	for _, u := range units.All() {
		for _, name := range u.Names() {
			a = append(a, name)
		}
	}
	return a
}

func TestParseCmd(t *testing.T) {
	cmds := []string{
		"20kg in lbs",
		"20 kg in lbs",
		"20 kilograms in pounds",
		"-20 C in F",
	}

	for _, cmd := range cmds {
		amount, fromStr, toStr, err := parseCmd(cmd)
		if err != nil {
			t.Errorf("unexpected parse error: %s", err)
			continue
		}
		if math.Abs(amount) != 20 {
			t.Errorf("parsed unexpected value: %v", amount)
			continue
		}
		t.Logf("parsed conversion: %s to %s", fromStr, toStr)
	}
}

func TestParseCmdFailure(t *testing.T) {
	amount, fromStr, toStr, err := parseCmd("20kg in")
	if err == nil {
		t.Errorf("missing expected parse error")
	}
}

func TestUnitLookup(t *testing.T) {
	for _, name := range aggrNames() {
		u, err := units.Find(name)
		if err != nil {
			t.Errorf(err.Error())
			continue
		}
		t.Logf("found unit by name: %s (%s)", name, u.Name)
	}
}

func TestUnitNameOverlap(t *testing.T) {
	nameMap := make(map[string]units.Unit)

	var total, failed int
	for _, u := range nameMap {
		for _, name := range u.Names() {
			if existing, ok := nameMap[name]; ok {
				t.Errorf("overlap in unit names: %s, %s (%s)", u.Name, existing.Name, name)
				failed++
			} else {
				nameMap[name] = u
			}
			total++
		}
	}
	t.Logf("tested %d unit names, %d overlaps", total, failed)
}

// ensure all units within the same quantity resolve
// a conversion path
func TestPathResolve(t *testing.T) {
	for _, q := range units.QuantityMap {
		units := q.Units()
		for _, u1 := range units {
			v1 := u1.MakeValue(1.0)
			for _, u2 := range units {
				if u1.Name == u2.Name {
					continue
				}
				_, err := v1.Convert(u2)
				if err != nil {
					t.Errorf("failed to resolve path: %s -> %s", u1.Name, u2.Name)
				}
			}
		}
	}
}
