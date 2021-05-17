package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "?ola mundo"

	fmt.Println(Solution(text))
}

func Solution(s string) string {
	var c uint8 = s[0]
	upper, _ := regexp.Match("[A-Z]", []byte(string(c)))
	lower, _ := regexp.Match("[a-z]", []byte(string(c)))
	digit, _ := regexp.Match("[0-9]", []byte(string(c)))

	if upper { // please fix condition
		return "upper"
	} else if lower { // please fix condition
		return "lower"
	} else if digit { // please fix condition
		return "digit"
	} else {
		return "other"
	}
}
