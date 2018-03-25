package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bcicen/go-units"
	"github.com/bcicen/xiny/log"
)

var (
	version  = "dev-build"
	build    = "unknown"
	usageStr = []string{
		"usage: xiny [options] [input]\n",
		"e.g xiny 20kg in lbs\n",
		"options",
	}
	versionStr = fmt.Sprintf("xiny version %s, build %s", version, build)
)

var opts = []Opt{
	{"i", "start xiny in interactive mode", func() { interactive() }},
	{"v", "enable verbose output", func() { log.Level = log.INFO }},
	{"vv", "enable debug output", func() { log.Level = log.DEBUG }},
	{"version", "print version info and exit", func() { fmt.Println(versionStr); os.Exit(0) }},
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
		spacing := "	"
		if len(opt.name) <= 5 {
			spacing = "		"
		}
		line := fmt.Sprintf(" -%s%s%s", opt.name, spacing, opt.desc)
		optUsage = append(optUsage, line)
	}

	fmt.Println(strings.Join(append(usageStr, optUsage...), "\n"))
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

	fromUnit, err := units.Find(convCmd.from)
	if err != nil {
		panic(err)
	}

	toUnit, err := units.Find(convCmd.to)
	if err != nil {
		panic(err)
	}

	if fromUnit.Quantity != toUnit.Quantity {
		e := fmt.Sprintf("%s -> %s", fromUnit.Quantity, toUnit.Quantity)
		panic(fmt.Errorf("unit mismatch: cannot convert %s", e))
	}

	path, err := units.ResolveConversion(fromUnit, toUnit)
	if err != nil {
		panic(err)
	}

	x := convCmd.amount
	var formula string
	for _, conv := range path {
		if formula != "" {
			formula = fmt.Sprintf("(%s)", strings.Replace(conv.Formula, "x", formula, 1))
		} else {
			formula = fmt.Sprintf("(%s)", conv.Formula)
		}
		log.Debugf("%s -> %s: %s", conv.From(), conv.To(), conv.Formula)
		x = conv.Fn(x)
	}
	log.Infof("%s -> %s: %s", fromUnit.Name, toUnit.Name, formula)

	return units.NewValue(x, toUnit).String()
}

func main() {
	defer recovery(true)

	if len(os.Args) <= 1 {
		usage()
	}

	cmd := strings.Join(os.Args[1:], " ")
	for _, optName := range parseOpts(&cmd) {
		handleOpt(optName)
	}

	fmt.Println(doConvert(cmd))
}
