package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	cmdRe = regexp.MustCompile("(-?[0-9.]+)\\s*(\\w+)\\s+in\\s+(\\w+)")
	optRe = regexp.MustCompile("-([a-z]+)")
)

func parseCmd(s string) (float64, string, string) {
	mg := cmdRe.FindStringSubmatch(s)
	if len(mg) != 4 {
		panic(fmt.Errorf("parse error"))
	}
	q, err := strconv.ParseFloat(mg[1], 6)
	if err != nil {
		panic(fmt.Errorf("parse error"))
	}
	return q, mg[2], mg[3]
}

// parse out and return options
func parseOpts(s *string) (opts []string) {
	ns := *s

	matches := optRe.FindAllString(ns, -1)
	for _, m := range matches {
		opts = append(opts, strings.TrimPrefix(m, "-"))
		ns = strings.Replace(ns, m, "", 1)
	}

	s = &ns
	return opts
}
