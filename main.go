package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/bcicen/xiny/units"
	_ "github.com/bcicen/xiny/units/mass"
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
	fmt.Printf("%g %s -> %s\n", q, u1, u2)

	//for _, mu := range units.All {
	//fmt.Printf("%s %s %.6g\n", mu.Name, mu.Symbol, mu.Ratio)
	//}

	fromUnit, err := units.Find(u1)
	if err != nil {
		panic(err)
	}

	toUnit, err := units.Find(u2)
	if err != nil {
		panic(err)
	}

	val := units.NewValue(q, fromUnit)
	fmt.Println(val.String())
	fmt.Println(val.FString(true))

	val.ToUnit(toUnit)
	fmt.Println(val.String())
	fmt.Println(val.FString(true))
}
