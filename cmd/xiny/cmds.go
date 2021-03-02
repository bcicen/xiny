package main

import (
	"os"
	"strings"

	"github.com/bcicen/xiny/internal/command"
	"github.com/c-bata/go-prompt"
)

var interactiveUsageTxt = `
[n] [from-unit] in [to-unit]

conversion examples:
  20kg in lbs
  7.98 ounces in grams
  1.44MB in KB
`

func init() {
	command.Register(
		&command.Command{
			Names:       []string{"help", "?"},
			Description: "show help message",
			Fn: func(string) string {
				return interactiveUsageTxt + command.UsageText()
			},
		},
		&command.Command{
			Names:       []string{"version"},
			Description: "show full version info",
			Fn: func(string) string {
				return versionStr
			},
		},
		&command.Command{
			Names:       []string{"exit", "q"},
			Description: "quit xiny",
			Fn: func(string) (res string) {
				os.Exit(0)
				return
			},
		},
	)
}

func cmdCompleter(txt string) []prompt.Suggest {

	var a []*command.Command
	cmd, rtxt := command.Dig(txt)
	nextWord := strings.Split(rtxt, " ")[0]

	for _, subcmd := range cmd.SubCommands {
		if subcmd.FuzzyMatch(nextWord) {
			a = append(a, subcmd)
		}
	}

	return cmdSuggests(a...)
}

func cmdSuggests(cmds ...*command.Command) []prompt.Suggest {

	a := make([]prompt.Suggest, len(cmds))

	for n := range cmds {
		cmd := cmds[n]
		a[n] = prompt.Suggest{
			Text:        cmd.Names[0],
			Description: cmd.Description,
		}
	}

	return a
}
