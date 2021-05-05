package main

import (
	"strings"
)

func main() {
}

// ok
func reverseStringArrayFields(s string) string {
	// para evitar espacios extras
	values := strings.Fields(s)

	// 2 indices uno con valor 0 y otro con el largo del array
	// intercambiar a la hora de devolver array
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		// swap words string,
		values[i], values[j] = values[j], values[i]
	}

	return strings.Join(values, " ")
}

// aproximado, espacios extras al final
func reverseStringFieldsBuilder(s string) string {
	values := strings.Fields(s)
	builder := strings.Builder{}

	for i, j := 0, len(values)-1; j >= 0; i, j = i+1, j-1 {
		// swap words string
		builder.WriteString(values[j])
		builder.WriteString(" ")
	}

	return builder.String()
}

func mirrorStringBytes(s string) string {
	a := []byte(s) // to byte

	// Reverse en a
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	return string(a)
}

func mirrorStringRunes(s string) string {
	rns := []rune(s) // to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		// swap letters string,
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func mirrorStringLikeRunes(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
