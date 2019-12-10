package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const num = 100000.5678

func main() {
	fmt.Println("-Ingles:")
	p := message.NewPrinter(language.English)
	p.Printf(" %.2f \n", num)

	fmt.Println("-Aleman:")
	p = message.NewPrinter(language.German)
	p.Printf(" %.2f \n", num)

	fmt.Println("-Español:")
	p = message.NewPrinter(language.Spanish)
	p.Printf(" %.2f \n", num)
}
