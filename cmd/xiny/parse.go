package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	cmdRe   = regexp.MustCompile("(-?[0-9.]+)\\s*(.+)\\s+in\\s+(.+)")
	optRe   = regexp.MustCompile("-([a-z]+)")
	stripRe = regexp.MustCompile(",")
)

func parseCmd(s string) (amount float64, fromStr, toStr string, err error) {
	s = stripRe.ReplaceAllString(s, "")
	mg := cmdRe.FindStringSubmatch(s)
	if len(mg) != 4 {
		return amount, fromStr, toStr, fmt.Errorf("incomplete command")
	}
	fromStr = mg[2]
	toStr = mg[3]

	q, err := strconv.ParseFloat(mg[1], 6)
	if err != nil {
		return amount, fromStr, toStr, fmt.Errorf("failed to parse value: %s", mg[1])
	}
	amount = q

	return amount, fromStr, toStr, nil
}

// parse out options from a given string args
func parseOpts(a []string) (cmd string, opts []string) {

	var n int
	var s string
	for n, s = range a {
		if !optRe.MatchString(s) {
			break
		}
		opts = append(opts, strings.TrimPrefix(s, "-"))
	}

	if n < len(a)-1 {
		cmd = strings.TrimSpace(strings.Join(a[n:], " "))
	}

	return
}
