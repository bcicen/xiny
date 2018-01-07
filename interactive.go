package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/bcicen/xiny/units"
	"github.com/c-bata/go-prompt"
)

var (
	ColorTheme = []prompt.Option{
		prompt.OptionPrefixTextColor(prompt.Blue),
		prompt.OptionPrefixBackgroundColor(prompt.DefaultColor),
		prompt.OptionInputTextColor(prompt.DefaultColor),
		prompt.OptionInputBGColor(prompt.DefaultColor),
		prompt.OptionPreviewSuggestionTextColor(prompt.Cyan),
		prompt.OptionPreviewSuggestionBGColor(prompt.DefaultColor),
		prompt.OptionSuggestionTextColor(prompt.White),
		prompt.OptionSuggestionBGColor(prompt.DarkGray),
		prompt.OptionSelectedSuggestionTextColor(prompt.Black),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionDescriptionTextColor(prompt.White),
		prompt.OptionDescriptionBGColor(prompt.DarkGray),
		prompt.OptionSelectedDescriptionTextColor(prompt.Black),
		prompt.OptionSelectedDescriptionBGColor(prompt.LightGray),
	}
	MonoTheme = []prompt.Option{
		prompt.OptionPrefixTextColor(prompt.Blue),
		prompt.OptionPrefixBackgroundColor(prompt.DefaultColor),
		prompt.OptionInputTextColor(prompt.DefaultColor),
		prompt.OptionInputBGColor(prompt.DefaultColor),
		prompt.OptionPreviewSuggestionTextColor(prompt.Cyan),
		prompt.OptionPreviewSuggestionBGColor(prompt.DefaultColor),
		prompt.OptionSuggestionTextColor(prompt.White),
		prompt.OptionSuggestionBGColor(prompt.DefaultColor),
		prompt.OptionSelectedSuggestionTextColor(prompt.Black),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionDescriptionTextColor(prompt.White),
		prompt.OptionDescriptionBGColor(prompt.DefaultColor),
		prompt.OptionSelectedDescriptionTextColor(prompt.Black),
		prompt.OptionSelectedDescriptionBGColor(prompt.DefaultColor),
	}
	unitSuggestions  = buildSuggest()
	emptySuggestions = []prompt.Suggest{}
	alphaRe          = regexp.MustCompile("([a-zA-Z]+)")
)

func buildSuggest() (a []prompt.Suggest) {
	for _, u := range units.UnitMap {
		a = append(a, prompt.Suggest{Text: u.Symbol})
		name := u.Name
		if u.Plural() {
			name += "s"
		}
		a = append(a, prompt.Suggest{Text: name})
	}
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

func stripNumbers(s string) string { return alphaRe.FindString(s) }

func Completer(d prompt.Document) []prompt.Suggest {
	//if !unicode.IsLetter(runeBeforeCursor(d)) {
	//return emptySuggestions
	//}

	args := strings.Split(d.TextBeforeCursor(), " ")
	argLen := len(args)

	if argLen == 0 {
		return emptySuggestions
	}

	// if first arg is both quantity and unit, treat it is two args
	if stripNumbers(args[0]) != "" {
		argLen += 1
	}

	if argLen == 3 {
		return []prompt.Suggest{{Text: "in"}}
	}

	w := d.GetWordBeforeCursor()
	if w == "" {
		return emptySuggestions
	}

	return prompt.FilterHasPrefix(unitSuggestions, w, true)
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
		append(opts, MonoTheme...)...,
	)
	p.Run()

	os.Exit(0)
}
