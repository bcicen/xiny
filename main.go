package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bcicen/xiny/log"
	"github.com/bcicen/xiny/units"
)

var (
	usageStr = []string{
		"usage: xiny [options] [input]\n",
		"e.g xiny 20kg in lbs\n",
		"options",
	}
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

func panicExit() {
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
		exitErr(err)
	}
}

func main() {
	defer panicExit()

	if len(os.Args) == 1 {
		usage()
	}

	cmd := strings.Join(os.Args[1:], " ")
	for _, optName := range parseOpts(&cmd) {
		handleOpt(optName)
	}

	convCmd, err := parseCmd(cmd)
	failOnErr(err, "parse error")

	fromUnit := units.MustFind(convCmd.from)
	toUnit := units.MustFind(convCmd.to)

	if fromUnit.Quantity != toUnit.Quantity {
		e := fmt.Sprintf("%s -> %s", fromUnit.Quantity.Name, toUnit.Quantity.Name)
		exitErr(fmt.Errorf("unit mismatch: cannot convert %s", e))
	}

	val := fromUnit.MakeValue(convCmd.amount)

	val, err = val.Convert(toUnit)
	failOnErr(err)

	fmtOpts := units.FmtOptions{false, 6}
	fmt.Println(val.Fmt(fmtOpts))
}
