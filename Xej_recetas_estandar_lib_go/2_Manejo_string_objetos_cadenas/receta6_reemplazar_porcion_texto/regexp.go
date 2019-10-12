package main

import (
	"fmt"
	"regexp"
)

const refString = "El hombre es el lobo del hombre"

func main() {
	regex := regexp.MustCompile("h[a-z]+")
	out := regex.ReplaceAllString(refString, "reemplazo")
	fmt.Println(out)
}
