package main

import (
	"fmt"
	"strings"
)

const refString = "Hola soy zeroidentidad, aprendo Go"

func main() {

	//se retornan boolean values
	lookFor := "zero"
	contain := strings.Contains(refString, lookFor)
	fmt.Printf("- \"%s\" contiene \"%s\": %t \n", refString, lookFor, contain)

	lookFor = "Java"
	contain = strings.Contains(refString, lookFor)
	fmt.Printf("- \"%s\" contiene \"%s\": %t \n", refString, lookFor, contain)

	startsWith := "Hola"
	starts := strings.HasPrefix(refString, startsWith)
	fmt.Printf("- \"%s\" inicia con \"%s\": %t \n", refString, startsWith, starts)

	endWith := "Go"
	ends := strings.HasSuffix(refString, endWith)
	fmt.Printf("- \"%s\" termina con \"%s\": %t \n", refString, endWith, ends)

}
