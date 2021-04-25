package main

import (
	"fmt"
	"strings"
)

const refString = "El hombre es el lobo del hombre"

func main() {
	replacer := strings.NewReplacer("lobo", "destructor", "hombre", "humano")
	out := replacer.Replace(refString)
	fmt.Println(out)
}
