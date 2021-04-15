package main

import (
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type cmdOptTest struct {
	input string
	cmd   string
	opts  []string
}

func TestParseCmdOpts(t *testing.T) {
	tests := []cmdOptTest{
		{
			"-v 20kg in lbs",
			"20kg in lbs",
			[]string{"v"},
		},
		{
			"-v -n 21 kg in lbs",
			"21 kg in lbs",
			[]string{"v", "n"},
		},
		{
			"-v -n -20 C in F",
			"-20 C in F",
			[]string{"v", "n"},
		},
		{
			"-21 C in F",
			"-21 C in F",
			nil,
		},
		{
			"-v C in F",
			"C in F",
			[]string{"v"},
		},
		{
			"-v",
			"",
			[]string{"v"},
		},
	}

	for _, x := range tests {
		cmd, opts := parseOpts(strings.Split(x.input, " "))
		assert.Equal(t, x.cmd, cmd, "unexpected parsed cmd")
		assert.Equal(t, x.opts, opts, "unexpected parsed opts")
	}
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
	_, _, _, err := parseCmd("20kg in")
	if err == nil {
		t.Errorf("missing expected parse error")
	}
}
