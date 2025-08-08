package cmd

import (
	"fmt"
	"github.com/c-bata/go-prompt"
)

func executor(input string) {
	fmt.Println(">>", input)
}

func completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: "add", Description: "add new event"},
		{Text: "list", Description: "list all events"},
		{Text: "remove", Description: "remove event"},
		{Text: "help", Description: "show help"},
		{Text: "exit", Description: "exit program"},
	}
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func Run() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("> "),
	)
	p.Run()
}
