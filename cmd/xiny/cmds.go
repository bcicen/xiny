package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

type Command struct {
	Names       []string
	Description string
	Fn          func(string) string
	CompleteFn  func(string) []prompt.Suggest
}

func (cmd *Command) name() string {
	return strings.Join(cmd.Names, ", ")
}

var (
	interactiveUsage = `
[n] [from-unit] in [to-unit]

conversion examples:
  20kg in lbs
  7.98 ounces in grams
  1.44MB in KB
`

	cmds = []*Command{
		{
			Names:       []string{"help", "?"},
			Description: "show help message",
			Fn: func(string) string {
				return interactiveUsage
			},
		},
		{
			Names:       []string{"version"},
			Description: "show full version info",
			Fn: func(string) string {
				return fmt.Sprintf("xiny v%s-%s", version, build)
			},
		},
		{
			Names:       []string{"exit", "q"},
			Description: "quit xiny",
			Fn: func(string) (res string) {
				os.Exit(0)
				return
			},
		},
	}
)

func getCmd(s string) *Command {
	for _, cmd := range cmds {
		for _, name := range cmd.Names {
			if s == name {
				return cmd
			}
		}
	}
	return nil
}

func suggestCmd(s string) (a []prompt.Suggest) {
	for _, cmd := range cmds {
		for _, name := range cmd.Names {
			if !strings.HasPrefix(name, s) {
				continue
			}
			a = append(a, prompt.Suggest{
				Text:        name,
				Description: cmd.Description,
			})
		}
	}
	return a
}

func init() {
	var sb strings.Builder
	sb.WriteString("\ncommands:\n")
	for _, cmd := range cmds {
		sb.WriteString("  ")
		sb.WriteString(fmt.Sprintf("%-15s", cmd.name()))
		sb.WriteString(cmd.Description)
		sb.WriteString("\n")
	}
	interactiveUsage += sb.String()
}
