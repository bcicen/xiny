package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	cmdRe = regexp.MustCompile("(-?[0-9.]+)\\s*(.+)\\s+in\\s+(.+)")
	optRe = regexp.MustCompile("-([a-z]+)")
)

type convertCmd struct {
	amount float64
	from   string
	to     string
}

func (c convertCmd) String() string {
	return fmt.Sprintf("%f %s in %s", c.amount, c.from, c.to)
}

func parseCmd(s string) (convertCmd, error) {
	var c convertCmd

	mg := cmdRe.FindStringSubmatch(s)
	if len(mg) != 4 {
		return c, fmt.Errorf("incomplete command")
	}
	c.from = mg[2]
	c.to = mg[3]

	q, err := strconv.ParseFloat(mg[1], 6)
	if err != nil {
		return c, fmt.Errorf("failed to parse value: %s", mg[1])
	}
	c.amount = q

	return c, nil
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
