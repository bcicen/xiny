package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bcicen/xiny/log"
	"github.com/bcicen/xiny/units"
)

func handleOpt(opt string) {
	switch opt {
	case "v":
		log.Level = log.INFO
	case "vv":
		log.Level = log.DEBUG
	case "list":
		listUnits()
		os.Exit(0)
	default:
		exitErr(fmt.Errorf("unknown option: -%s", opt))
	}
}

func usage() {
	fmt.Println("usage: xiny [options] [input]\n")
	fmt.Println("e.g xiny 20kg in lbs\n")
	os.Exit(1)
}

func listUnits() { fmt.Println(strings.Join(units.Names(), "\n")) }

func exitErr(err error) {
	log.Error(err)
	os.Exit(1)
}

func main() {
	if len(os.Args) == 1 {
		usage()
	}

	cmd := strings.Join(os.Args[1:], " ")
	for _, opt := range parseOpts(&cmd) {
		handleOpt(opt)
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

	fmtOpts := units.FmtOptions{true, 3}
	fmt.Println(newVal.Fmt(fmtOpts))
}
