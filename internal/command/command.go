package command

import (
	"fmt"
	"strings"
)

var root = &Command{} // top-level placeholder root command

// UsageText returns a concatenated usage string for all
// registered commands
func UsageText() string {
	var sb strings.Builder
	sb.WriteRune('\n')
	sb.WriteString("commands:")
	sb.WriteRune('\n')
	for _, cmd := range root.SubCommands {
		sb.WriteString("  ")
		names := strings.Join(cmd.Names, ", ")
		sb.WriteString(fmt.Sprintf("%-15s", names))
		sb.WriteString(cmd.Description)
		sb.WriteRune('\n')
	}
	return sb.String()
}

// Register registers a top-level Command to the app
// root, panicking on any name conflict
func Register(cmds ...*Command) {
	for _, cmd := range cmds {
		for _, other := range cmd.SubCommands {
			for _, name := range cmd.Names {
				if other.Match(name) {
					panic(fmt.Errorf("duplicate cmd name: %s", name))
				}
			}
		}
		root.SubCommands = append(root.SubCommands, cmd)
	}
}

// Dig recurses registered commands and subcommands, returning the
// last subcommand exactly matched in the given space-delimited string,
// and all remaining unmatched text
func Dig(s string) (*Command, string) {

	var (
		n     int
		cmd   = root
		parts = strings.Split(s, " ")
	)

loop:
	for n = range parts {
		for _, subcmd := range cmd.SubCommands {
			if subcmd.Match(parts[n]) {
				cmd = subcmd
				continue loop
			}
		}
		// at last match
		break
	}

	return cmd, strings.Join(parts[n:], " ")
}

// Get returns a top-level registered Command exactly matching
// a given name
func Get(s string) *Command {
	for _, cmd := range root.SubCommands {
		for _, name := range cmd.Names {
			if s == name {
				return cmd
			}
		}
	}
	return nil
}
