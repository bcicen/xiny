package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/bcicen/xiny/units"
	"github.com/c-bata/go-prompt"
)

var (
	Theme = []prompt.Option{
		prompt.OptionPrefixTextColor(prompt.Blue),
		prompt.OptionPrefixBackgroundColor(prompt.DefaultColor),
		prompt.OptionInputTextColor(prompt.DefaultColor),
		prompt.OptionInputBGColor(prompt.DefaultColor),
		prompt.OptionPreviewSuggestionTextColor(prompt.Cyan),
		prompt.OptionPreviewSuggestionBGColor(prompt.DefaultColor),
		prompt.OptionSuggestionTextColor(prompt.White),
		prompt.OptionSuggestionBGColor(prompt.DefaultColor),
		prompt.OptionSelectedSuggestionTextColor(prompt.Black),
		prompt.OptionSelectedSuggestionBGColor(prompt.White),
		prompt.OptionDescriptionTextColor(prompt.LightGray),
		prompt.OptionDescriptionBGColor(prompt.DefaultColor),
		prompt.OptionSelectedDescriptionTextColor(prompt.LightGray),
		prompt.OptionSelectedDescriptionBGColor(prompt.DefaultColor),
	}
	quantityFilterStr string
	unitSuggestions   = buildSuggest(false)
	emptySuggestions  = []prompt.Suggest{}
	progress1Re       = regexp.MustCompile("-?[0-9.]+\\s+")
	progress2Re       = regexp.MustCompile("(-?[0-9.]+)\\s*([a-zA-Z|\\s]+)\\s+")
	progress3Re       = regexp.MustCompile("(-?[0-9.]+)\\s*(.+)\\s+in\\s+")
)

type UnitSuggests []prompt.Suggest

func (a UnitSuggests) Len() int           { return len(a) }
func (a UnitSuggests) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a UnitSuggests) Less(i, j int) bool { return a[i].Text < a[j].Text }

func buildSuggest(includeSymbols bool) (a UnitSuggests) {
	for _, u := range units.UnitMap {
		if includeSymbols {
			a = append(a, prompt.Suggest{Text: u.Symbol, Description: u.Quantity.Name})
		}
		name := u.Name
		if u.Plural() {
			name += "s"
		}
		a = append(a, prompt.Suggest{Text: name, Description: u.Quantity.Name})
	}
	sort.Sort(a)
	return a
}

func Executor(s string) {
	s = strings.TrimSpace(s)
	if s == "exit" {
		os.Exit(0)
	}
	defer recovery(false)
	fmt.Println(doConvert(s))
}

func runeBeforeCursor(d prompt.Document) rune {
	empty := ' '
	pos := d.CursorPosition - 1

	if pos < 0 {
		return empty
	}

	r := []rune(d.Text)

	if len(r) > 0 {
		for pos >= 0 {
			if r[pos] != empty {
				fmt.Println(string(r[pos]))
				fmt.Println(unicode.IsNumber(r[pos]))
				return r[pos]
			}
			pos--
		}
	}

	return empty
}

func filterContains(suggests []prompt.Suggest, sub string) []prompt.Suggest {
	sub = strings.ToLower(sub)

	var filtered []prompt.Suggest
	for _, s := range suggests {
		if strings.Contains(s.Text, sub) {
			filtered = append(filtered, s)
		}
	}

	return filtered
}

func filterQuantity() []prompt.Suggest {
	if quantityFilterStr == "" {
		return unitSuggestions
	}

	var filtered []prompt.Suggest
	for _, us := range unitSuggestions {
		if us.Description == quantityFilterStr {
			filtered = append(filtered, us)
		}
	}

	return filtered
}

func Completer(d prompt.Document) []prompt.Suggest {
	cmd := d.TextBeforeCursor()
	w := d.GetWordBeforeCursor()

	if cmd == "" {
		return emptySuggestions
	}

	if progress3Re.FindString(cmd) != "" {
		return filterContains(filterQuantity(), w)
	}

	mg := progress2Re.FindStringSubmatch(cmd)
	if mg != nil {
		if quantityFilterStr == "" {
			fromName := strings.Trim(mg[2], " ")
			unit, err := units.Find(fromName)
			if err == nil {
				quantityFilterStr = unit.Quantity.Name
			}
		}
		if quantityFilterStr != "" {
			return []prompt.Suggest{{Text: "in", Description: "keyword"}}
		}
	}

	if progress1Re.FindString(cmd) != "" {
		return filterContains(unitSuggestions, w)
	}

	quantityFilterStr = ""
	return emptySuggestions
}

func interactive() {
	fmt.Printf("xiny version %s\n", version)
	fmt.Println("use `exit` or `ctrl-d` to exit")
	defer fmt.Println("bye!")
	opts := []prompt.Option{
		prompt.OptionTitle("xiny interactive mode"),
		prompt.OptionPrefix(">>> "),
	}
	p := prompt.New(
		Executor,
		Completer,
		append(opts, Theme...)...,
	)
	p.Run()

	os.Exit(0)
}
