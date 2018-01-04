package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bcicen/xiny/log"
	"github.com/bcicen/xiny/units"
)

type Opt struct {
	name    string
	desc    string
	handler func()
}

func handleOpt(s string) {
	if s == "h" || s == "help" {
		usage()
	}
	for _, opt := range opts {
		if opt.name == s {
			opt.handler()
			return
		}
	}
	exitErr(fmt.Errorf("unknown option: -%s", s))
}

var opts = []Opt{
	{"v", "enable verbose output", func() { log.Level = log.INFO }},
	{"vv", "enable debug output", func() { log.Level = log.DEBUG }},
	{"list", "list all potential unit names and exit", listUnits},
}

func usage() {
	s := []string{
		"usage: xiny [options] [input]\n",
		"e.g xiny 20kg in lbs\n",
		"options",
	}

	for _, opt := range opts {
		s = append(s, fmt.Sprintf(" -%s	%s", opt.name, opt.desc))
	}

	fmt.Println(strings.Join(s, "\n"))
	os.Exit(1)
}

func listUnits() {
	fmt.Println(strings.Join(units.Names(), "\n"))
	os.Exit(0)
}

func exitErr(err error) {
	log.Error(err)
	os.Exit(1)
}

func main() {
	if len(os.Args) == 1 {
		usage()
	}

	cmd := strings.Join(os.Args[1:], " ")
	for _, optName := range parseOpts(&cmd) {
		handleOpt(optName)
	}

	q, u1, u2 := parseCmd(cmd)

	fromUnit, err := units.Find(u1)
	if err != nil {
		exitErr(err)
	}

	toUnit, err := units.Find(u2)
	if err != nil {
		exitErr(err)
	}

	if fromUnit.Quantity != toUnit.Quantity {
		e := fmt.Sprintf("%s -> %s", fromUnit.Quantity.Name, toUnit.Quantity.Name)
		exitErr(fmt.Errorf("unit mismatch: cannot convert %s", e))
	}

	val := fromUnit.MakeValue(q)

	newVal, err := val.Convert(toUnit)
	if err != nil {
		exitErr(err)
	}

	fmtOpts := units.FmtOptions{false, 6}
	fmt.Println(newVal.Fmt(fmtOpts))
}
