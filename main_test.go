package main

import (
	"math"
	"testing"

	"github.com/bcicen/xiny/units"
)

func TestParseCmd(t *testing.T) {
	cmds := []string{
		"20kg in lbs",
		"20 kg in lbs",
		"20 kilograms in pounds",
		"-20 C in F",
	}

	for _, cmd := range cmds {
		convCmd, err := parseCmd(cmd)
		if err != nil {
			t.Errorf("unexpected parse error: %s", err)
			continue
		}
		if math.Abs(convCmd.amount) != 20 {
			t.Errorf("parsed unexpected value: %v", convCmd.amount)
			continue
		}
		t.Logf("parsed conversion: %s", convCmd)
	}
}

func TestParseCmdFailure(t *testing.T) {
	_, err := parseCmd("20kg in")
	if err == nil {
		t.Errorf("missing expected parse error")
	}
}

func TestUnitLookup(t *testing.T) {
	for _, name := range units.Names() {
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
	for _, u := range units.UnitMap {
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
