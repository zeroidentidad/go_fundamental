package bytestrings

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// SearchString muestra varios métodos
// para buscar una cadena
func SearchString() {
	s := "esto es una prueba perro"

	// devuelve verdadero porque s contiene
	// la palabra esto
	fmt.Println(strings.Contains(s, "esto"))

	// devuelve verdadero porque s contiene la letra a
	// también coincidiría si contuviera b o c
	fmt.Println(strings.ContainsAny(s, "abc"))

	// devuelve verdadero porque s comienza con esto
	fmt.Println(strings.HasPrefix(s, "esto"))

	// devuelve verdadero porque s termina con prueba
	fmt.Println(strings.HasSuffix(s, "prueba"))
}

// ModifyString modifica una cadena de varias maneras
func ModifyString() {
	s := "simple string"

	// prints [simple string]
	fmt.Println(strings.Split(s, " "))

	// prints "Simple String"
	fmt.Println(strings.Title(s))

	// prints "simple string"; todo al final y
	// se elimina el espacio en blanco inicial
	s = " simple string "
	fmt.Println(strings.TrimSpace(s))
}

// StringReader demuestra cómo crear
// una interfaz io.Reader rápidamente con una cadena
func StringReader() {
	s := "simple string\n"
	r := strings.NewReader(s)

	// prints s en Stdout
	io.Copy(os.Stdout, r)
}
