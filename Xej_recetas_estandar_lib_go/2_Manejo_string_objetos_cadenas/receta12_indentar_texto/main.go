package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	text := "Oh! Bueno, Go es super."
	text = Indent(text, 6)
	fmt.Println(text)

	text = Unindent(text, 3)
	fmt.Println(text)

	text = Unindent(text, 10)
	fmt.Println(text)

	text = IndentByRune(text, 10, '.')
	fmt.Println(text)
}

// Sangrado de la entrada por indentación dada y runa .
func IndentByRune(input string, indent int, r rune) string {
	return strings.Repeat(string(r), indent) + input
}

// Sangrado de la entrada por indentación dada
func Indent(input string, indent int) string {
	padding := indent + len(input)
	return fmt.Sprintf("% "+strconv.Itoa(padding)+"s", input)
}

// Unindent desangra la cadena de entrada. En caso de que
// la entrada está sangrada por menos de espacios de "sangría"
// se elimina el mínimo de ambos.
func Unindent(input string, indent int) string {

	count := 0
	for _, val := range input {
		if unicode.IsSpace(val) {
			count++
		}
		if count == indent || !unicode.IsSpace(val) {
			break
		}
	}

	return input[count:]
}
