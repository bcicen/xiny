package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bcicen/go-units"
	"github.com/c-bata/go-prompt"
)

type Command struct {
	Names       []string
	Description string
	Fn          func(string) string
	CompleteFn  func(string) []prompt.Suggest
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
			Names:       []string{"info"},
			Description: "show short info and stats",
			Fn: func(string) string {
				var sb strings.Builder
				sb.WriteString(versionStr)
				sb.WriteRune('\n')
				sb.WriteString(strconv.Itoa(len(units.All())))
				sb.WriteString(" units")
				sb.WriteRune('\n')
				return sb.String()
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

func cmdCompleter(txt string) []prompt.Suggest {
	parts := strings.SplitN(txt, " ", 2)
	//if len(parts) <= 1 {
	//return suggestCmd(txt)
	//}

	//cmd := getCmd(parts[0])
	//if cmd != nil && cmd.CompleteFn != nil {
	//return cmd.CompleteFn(parts[1])
	//}

	//return nil

	a := []prompt.Suggest{
		{
			Text:        "text",
			Description: txt,
		},
	}

	for n, s := range parts {
		a = append(a, prompt.Suggest{
			Text:        fmt.Sprintf("part%d", n),
			Description: s,
		})
	}
	return a
}

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
		names := strings.Join(cmd.Names, ", ")
		sb.WriteString(fmt.Sprintf("%-15s", names))
		sb.WriteString(cmd.Description)
		sb.WriteString("\n")
	}
	interactiveUsage += sb.String()
}
