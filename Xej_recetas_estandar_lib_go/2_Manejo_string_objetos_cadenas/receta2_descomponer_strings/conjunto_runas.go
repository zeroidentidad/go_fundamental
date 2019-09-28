package main

import (
	"fmt"
	"strings"
)

const refString = "Daría*todo lo%que sé, por_la mitad de,lo que ignoro."

func main() {

	// Se llama a splitFunc para cada runa en una cadena. Si la runa
	// es igual a cualquiera de los caracteres en del conjunto "*%, _"
	// la refString será dividida.
	splitFunc := func(r rune) bool {
		return strings.ContainsRune("*%,_", r)
	}

	palabras := strings.FieldsFunc(refString, splitFunc)
	for idx, palabra := range palabras {
		fmt.Printf("Palabra %d es: %s\n", idx, palabra)
	}

}
