package src

import (
	"github.com/c-bata/go-prompt"
)

// exported func
func Completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: "Store data 1"},
		{Text: "articles", Description: "Store data 2"},
		{Text: "comments", Description: "Store data 3"},
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
