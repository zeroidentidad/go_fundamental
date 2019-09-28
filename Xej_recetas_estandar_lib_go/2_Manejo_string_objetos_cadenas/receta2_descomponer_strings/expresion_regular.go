package main

import (
	"fmt"
	"regexp"
)

const refString = "Cada día*sabemos, más_y entendemos%menos."

func main() {

	//cada coinciendia del conjunto indicado, web de utilidad: regexr.com
	palabras := regexp.MustCompile("[*,%_]{1}").Split(refString, -1)
	for idx, palabra := range palabras {
		fmt.Printf("Palabra %d es: %s\n", idx, palabra)
	}

}
