package main

import (
	"math"
	"testing"
)

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
	_, _, _, err := parseCmd("20kg in")
	if err == nil {
		t.Errorf("missing expected parse error")
	}
}
