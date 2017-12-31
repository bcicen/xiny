package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/bcicen/xiny/units"
)

var re = regexp.MustCompile("^([0-9.]+)(\\w+) in (\\w+)")

func parse(s string) (float64, string, string) {
	fmt.Println(s)
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

func main() {
	q, u1, u2 := parse(strings.Join(os.Args[1:], " "))

	fromQ, fromUnit, err := units.Find(u1)
	if err != nil {
		panic(err)
	}

	val := fromUnit.MakeValue(q)

	_, toUnit, err := units.Find(u2)
	if err != nil {
		panic(err)
	}

	newVal, err := fromQ.Convert(val, toUnit)
	if err != nil {
		panic(err)
	}

	fmtOpts := units.FmtOptions{true, 3}

	fmt.Println(fromQ.FmtValue(newVal, fmtOpts))
}
