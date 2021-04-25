package main

import (
	"fmt"
	"strings"
)

const refString = "El hombre es el lobo del hombre"
const refStringTwo = "destructor destructor destructor destructor"

func main() {
	out := strings.Replace(refString, "lobo", "destructor", -1)
	fmt.Println(out)

	out = strings.Replace(refStringTwo, "lobo", "destructor", 2)
	fmt.Println(out)
}
