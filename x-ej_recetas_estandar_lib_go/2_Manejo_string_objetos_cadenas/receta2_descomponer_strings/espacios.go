package main

import (
	"fmt"
	"strings"
)

const refString = "Lo que sabemos es una gota de agua; lo que ignoramos es el océano."

func main() {

	palabras := strings.Fields(refString)
	for idx, palabra := range palabras {
		fmt.Printf("Palabra %d es: %s\n", idx, palabra)
	}

}
