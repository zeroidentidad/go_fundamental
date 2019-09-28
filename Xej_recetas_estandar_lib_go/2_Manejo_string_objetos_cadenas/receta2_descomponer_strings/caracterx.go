package main

import (
	"fmt"
	"strings"
)

const refString = "Vale_más saber alguna_cosa de todo_, que_saberlo todo de una_sola cosa."

func main() {

	palabras := strings.Split(refString, "_")
	for idx, palabra := range palabras {
		fmt.Printf("Porción %d es: %s\n", idx, palabra)
	}

}
