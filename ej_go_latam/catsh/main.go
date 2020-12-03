package main

import (
	"fmt"

	"catsh/src"

	"github.com/c-bata/go-prompt"
)

func main() {
	fmt.Println("Select table.")
	t := prompt.Input("-| ", src.Completer)
	fmt.Println("You selected", t)
}
