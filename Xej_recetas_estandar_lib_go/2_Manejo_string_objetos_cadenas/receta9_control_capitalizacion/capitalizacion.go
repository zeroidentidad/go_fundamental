package main

import (
	"fmt"
	"strings"
	"unicode"
)

const email = "EjemPlo@dominio.com"
const name = "steve jobs"
const upc = "upc"
const i = "i"

const snakeCase = "primer_nombre"

func main() {

	// Para comparar la entrada del usuario
	// a veces es mejor
	// comparar la entrada en un mismo
	// caso.
	input := "Ejemplo@dominio.com"
	input = strings.ToLower(input)
	emailToCompare := strings.ToLower(email)
	matches := input == emailToCompare
	fmt.Printf("Email compatible: %t\n", matches)

	upcCode := strings.ToUpper(upc)
	fmt.Println("UPPER case: " + upcCode)

	// Este dígrafo tiene diferentes mayúsculas y
	// titulo del caso.
	str := "dz"
	fmt.Printf("%s en mayusculas: %s y titulo: %s \n",
		str,
		strings.ToUpper(str),
		strings.ToTitle(str))

	// Uso de la función XXXSpecial
	title := strings.ToTitle(i)
	titleTurk := strings.ToTitleSpecial(unicode.TurkishCase, i)
	if title != titleTurk {
		fmt.Printf("ToTitle es diferente: %#U vs. %#U \n",
			title[0],
			[]rune(titleTurk)[0])
	}

	// En algunos casos la entrada
	// necesita ser corregido por si acaso.
	correctNameCase := strings.Title(name)
	fmt.Println("Nombre corregido: " + correctNameCase)

	// Convirtiendo el snake case
	// a camel case con el uso de
	// las funciones Title y ToLower
	firstNameCamel := toCamelCase(snakeCase)
	fmt.Println("Camel case: " + firstNameCamel)

}

func toCamelCase(input string) string {
	titleSpace := strings.Title(strings.Replace(input, "_", " ", -1))
	camel := strings.Replace(titleSpace, " ", "", -1)
	return strings.ToLower(camel[:1]) + camel[1:]
}
