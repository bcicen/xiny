package command

import (
	"strings"
)

type Command struct {
	Names       []string
	Description string
	Fn          func(string) string
	SubCommands []*Command
}

// FuzzyMatch returns whether a given string matches prefixes
// any of Commands Names
func (cmd *Command) FuzzyMatch(s string) bool {
	for _, name := range cmd.Names {
		if strings.HasPrefix(name, s) {
			return true
		}
	}
	return false
}

// Match returns whether a given string matches any of
// Commands Names
func (cmd *Command) Match(s string) bool {
	for _, name := range cmd.Names {
		if s == name {
			return true
		}
	}
	return false
}
