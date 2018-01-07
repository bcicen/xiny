package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bcicen/xiny/log"
	"github.com/bcicen/xiny/units"
)

var (
	version  = "dev-build"
	usageStr = []string{
		"usage: xiny [options] [input]\n",
		"e.g xiny 20kg in lbs\n",
		"options",
	}
)

var opts = []Opt{
	{"i", "use interactive mode", func() { interactive() }},
	{"v", "enable verbose output", func() { log.Level = log.INFO }},
	{"vv", "enable debug output", func() { log.Level = log.DEBUG }},
	{"list", "list all potential unit names and exit", listUnits},
}

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

func usage() {
	var optUsage []string
	for _, opt := range opts {
		optUsage = append(optUsage, fmt.Sprintf(" -%s	%s", opt.name, opt.desc))
	}

	fmt.Println(strings.Join(append(usageStr, optUsage...), "\n"))
	os.Exit(0)
}

func listUnits() {
	fmt.Println(strings.Join(units.Names(), "\n"))
	os.Exit(0)
}

func failOnErr(err error, pfix ...string) {
	if err != nil {
		if len(pfix) != 0 {
			err = fmt.Errorf("%s: %s", pfix[0], err)
		}
		exitErr(err)
	}
}

func exitErr(err error) {
	log.Error(err)
	os.Exit(1)
}

func recovery(exit bool) {
	var err error
	if r := recover(); r != nil {
		switch x := r.(type) {
		case string:
			err = fmt.Errorf(r.(string))
		case error:
			err = x
		default:
			panic(r)
		}
		log.Error(err)
		if exit {
			os.Exit(1)
		}
	}
	if exit {
		os.Exit(0)
	}
}

func doConvert(cmd string) string {
	convCmd, err := parseCmd(cmd)
	if err != nil {
		panic(fmt.Errorf("parse error: %s", err))
	}

	fromUnit := units.MustFind(convCmd.from)
	toUnit := units.MustFind(convCmd.to)

	if fromUnit.Quantity != toUnit.Quantity {
		e := fmt.Sprintf("%s -> %s", fromUnit.Quantity.Name, toUnit.Quantity.Name)
		panic(fmt.Errorf("unit mismatch: cannot convert %s", e))
	}

	val := fromUnit.MakeValue(convCmd.amount)

	val, err = val.Convert(toUnit)
	if err != nil {
		panic(err)
	}

	fmtOpts := units.FmtOptions{false, 6}
	return val.Fmt(fmtOpts)
}

func main() {
	defer recovery(true)

	if len(os.Args) == 1 {
		usage()
	}

	cmd := strings.Join(os.Args[1:], " ")
	for _, optName := range parseOpts(&cmd) {
		handleOpt(optName)
	}

	fmt.Println(doConvert(cmd))
}
