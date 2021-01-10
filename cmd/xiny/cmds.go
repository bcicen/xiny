package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bcicen/go-units"
	"github.com/c-bata/go-prompt"
)

var Root = &Command{}

// AddCommand registers a top-level Command to the app
// root, panicing on any name conflict
func AddCommand(cmd *Command) {
	for _, other := range cmd.SubCommands {
		for _, name := range cmd.Names {
			if other.Match(name) {
				panic(fmt.Errorf("duplicate cmd name: %s", name))
			}
		}
	}
	Root.SubCommands = append(Root.SubCommands, cmd)
}

type Command struct {
	Names       []string
	Description string
	Fn          func(string) string
	SubCommands []*Command
}

// Suggest returns this command represented as one or more
// prompt completion suggestions
func (cmd *Command) Suggest() []prompt.Suggest {
	a := make([]prompt.Suggest, len(cmd.Names))
	for n := range cmd.Names {
		a[n] = prompt.Suggest{
			Text:        cmd.Names[n],
			Description: cmd.Description,
		}
	}
	return a
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

var (
	interactiveUsage = `
[n] [from-unit] in [to-unit]

conversion examples:
  20kg in lbs
  7.98 ounces in grams
  1.44MB in KB
`

	configSet = &Command{
		Names:       []string{"set"},
		Description: "set config value",
		Fn: func(s string) string {
			return fmt.Sprintf("SET %s", s)
		},
	}

	configGet = &Command{
		Names:       []string{"get"},
		Description: "get config value",
		Fn: func(s string) string {
			return fmt.Sprintf("GET %s", s)
		},
	}
)

func init() {
	AddCommand(&Command{
		Names:       []string{"help", "?"},
		Description: "show help message",
		Fn: func(string) string {
			return interactiveUsage
		},
	})

	AddCommand(&Command{
		Names:       []string{"config"},
		Description: "get and set configuration values",
		SubCommands: []*Command{
			configSet,
			configGet,
		},
	})

	AddCommand(&Command{
		Names:       []string{"info"},
		Description: "show short info and stats",
		Fn: func(string) string {
			var sb strings.Builder
			sb.WriteString(versionStr)
			sb.WriteRune('\n')
			sb.WriteRune('\n')
			sb.WriteString(strconv.Itoa(len(units.All())))
			sb.WriteString(" units")
			sb.WriteRune('\n')
			return sb.String()
		},
	})

	AddCommand(&Command{
		Names:       []string{"version"},
		Description: "show full version info",
		Fn: func(string) string {
			return versionStr
		},
	})

	AddCommand(&Command{
		Names:       []string{"exit", "q"},
		Description: "quit xiny",
		Fn: func(string) (res string) {
			os.Exit(0)
			return
		},
	})
}

func cmdCompleter(txt string) []prompt.Suggest {

	var (
		a []prompt.Suggest

		cmd   = Root
		parts = strings.Split(txt, " ")
	)

	for _, p := range parts {
		for _, subcmd := range cmd.SubCommands {
			if subcmd.Match(p) {
				cmd = subcmd
				break
			}
		}
	}

	for _, subcmd := range cmd.SubCommands {
		if subcmd.FuzzyMatch(parts[len(parts)-1]) {
			a = append(a, subcmd.Suggest()...)
		}
	}

	return a

	//a := []prompt.Suggest{
	//{
	//Text:        "text",
	//Description: txt,
	//},
	//}

	//for n, s := range parts {
	//a = append(a, prompt.Suggest{
	//Text:        fmt.Sprintf("part%d", n),
	//Description: s,
	//})
	//}
	//return a
}

func getCmd(s string) *Command {
	for _, cmd := range Root.SubCommands {
		for _, name := range cmd.Names {
			if s == name {
				return cmd
			}
		}
	}
	return nil
}

func init() {
	var sb strings.Builder
	sb.WriteString("\ncommands:\n")
	for _, cmd := range Root.SubCommands {
		sb.WriteString("  ")
		names := strings.Join(cmd.Names, ", ")
		sb.WriteString(fmt.Sprintf("%-15s", names))
		sb.WriteString(cmd.Description)
		sb.WriteString("\n")
	}
	interactiveUsage += sb.String()
}
