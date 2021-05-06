package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func getInput() string {
	fmt.Print("Ingresar una oración: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func isNotLetter(c rune) bool {
	return !unicode.IsLetter(c)
}

func isPalindromic(s string) bool {
	// split, remove non-alphabetic
	words := strings.FieldsFunc(s, isNotLetter)

	l := len(words)
	for i := 0; i < l/2; i++ {
		fw := words[i]     // palabra delantera
		bw := words[l-i-1] // palabra trasera
		if !strings.EqualFold(fw, bw) {
			return false
		}
	}

	return true
}

func main() {
	for l := getInput(); l != ""; l = getInput() {
		if isPalindromic(l) {
			fmt.Println("... es palindrómica!")
		} else {
			fmt.Println("... no es palindrómica.")
		}
	}
}
