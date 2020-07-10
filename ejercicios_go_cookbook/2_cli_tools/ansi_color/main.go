package main

import (
	"fmt"

	"./ansi"
)

func main() {
	r := ansi.ColorText{
		TextColor: ansi.Red,
		Text:      "I'm red!",
	}

	fmt.Println(r.String())

	r.TextColor = ansi.Green
	r.Text = "Now I'm green!"

	fmt.Println(r.String())

	r.TextColor = ansi.ColorNone
	r.Text = "Back normal..."

	fmt.Println(r.String())
}
