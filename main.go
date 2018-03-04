package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/bcicen/xiny/log"
	"github.com/bcicen/xiny/units"
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
	{"list", "list all potential unit names and exit", listUnits},
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

func listUnits() {
	var a []string
	for name, u := range units.UnitMap {
		s := name
		if u.Symbol != "" {
			s += fmt.Sprintf(" (%s)", u.Symbol)
		}
		a = append(a, s)
	}
	sort.Strings(a)
	fmt.Println(strings.Join(a, "\n"))
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

	return val.String()
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
