package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bcicen/go-units"
	"github.com/bcicen/xiny/log"
)

var (
	version = "dev-build"
	build   = "unknown"
	//examples = []string{
	//"xiny 20kg in lbs",
	//"xiny 7.98 ounces in grams",
	//"xiny 1.44MB in KB",
	//"xiny version",
	//}
	usageStr = []string{
		"usage: xiny [options] [input]\n",
		"e.g", "  xiny 20kg in lbs",
		"  xiny 7.98 ounces in grams",
		"  xiny 1.44MB in KB",
		"  xiny version\n",
		"options",
	}
	versionStr = fmt.Sprintf("xiny v%s-%s", version, build)

	fmtOpts = units.DefaultFmtOptions
)

var opts = []Opt{
	{"n", "display only numeric output (exclude units)", func() { fmtOpts.Label = false }},
	{"v", "enable more verbose output (twice for debug)", func() { log.Level += 1 }},
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
	x, fromStr, toStr, err := parseCmd(cmd)
	if err != nil {
		panic(fmt.Errorf("parse error: %s", err))
	}

	fromUnit, err := units.Find(fromStr)
	if err != nil {
		panic(err)
	}

	toUnit, err := units.Find(toStr)
	if err != nil {
		panic(err)
	}

	// allow convert to same unit
	if fromUnit.Name == toUnit.Name {
		return units.NewValue(x, fromUnit).Fmt(fmtOpts)
	}

	if fromUnit.Quantity != toUnit.Quantity {
		e := fmt.Sprintf("%s -> %s", fromUnit.Quantity, toUnit.Quantity)
		panic(fmt.Errorf("unit mismatch: cannot convert %s", e))
	}

	path, err := units.ResolveConversion(fromUnit, toUnit)
	if err != nil {
		panic(err)
	}

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

	return units.NewValue(x, toUnit).Fmt(fmtOpts)
}

func main() {
	defer recovery(true)

	var cmd string
	var opts []string

	if len(os.Args) > 1 {
		cmd, opts = parseOpts(os.Args[1:])
	}

	for _, optName := range opts {
		handleOpt(optName)
	}

	switch cmd {
	case "":
		interactive()
	case "version":
		fmt.Println(versionStr)
	default:
		fmt.Println(doConvert(cmd))
	}
}
