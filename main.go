package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/bcicen/xiny/log"
	"github.com/bcicen/xiny/units"
)

var re = regexp.MustCompile("([0-9.]+)\\s*(\\w+)\\s+in\\s+(\\w+)")

func parse(s string) (float64, string, string) {
	mg := re.FindStringSubmatch(s)
	if len(mg) != 4 {
		panic(fmt.Errorf("parse error"))
	}
	q, err := strconv.ParseFloat(mg[1], 6)
	if err != nil {
		panic(fmt.Errorf("parse error"))
	}
	return q, mg[2], mg[3]
}

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

// parse out and handle options, returning the remainder of the command line
func parseOpts(s string) string {
	ore := regexp.MustCompile("-([a-z]+)")
	matches := ore.FindAllString(s, -1)

	for _, m := range matches {
		opt := strings.TrimPrefix(m, "-")
		handleOpt(opt)
		s = strings.Replace(s, m, "", 1)
	}

	return s
}

func listUnits() { fmt.Println(strings.Join(units.Names(), "\n")) }

func exitErr(err error) {
	log.Error(err)
	os.Exit(1)
}

func main() {

	cmd := parseOpts(strings.Join(os.Args[1:], " "))

	q, u1, u2 := parse(cmd)

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
