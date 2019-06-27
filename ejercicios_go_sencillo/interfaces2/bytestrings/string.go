package bytestrings

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// SearchString nos muestra varios métodos
// para buscar en un string
func SearchString() {
	s := "esto es una prueba"

	// devuelve true porque s contiene
	// la palabra esto
	fmt.Println(strings.Contains(s, "esto"))

	// devuelve true porque s contiene la letra a
	// aunque no contenga las otras dos letras
	fmt.Println(strings.ContainsAny(s, "aic"))

	// devuelve true porque s empieza con esto
	fmt.Println(strings.HasPrefix(s, "esto"))

	// devuelve true porque s finaliza con esto
	fmt.Println(strings.HasSuffix(s, "prueba"))
}

// ModifyString modifica un string de distintos modos
func ModifyString() {
	s := "simple string"

	// prints [simple string]
	fmt.Println(strings.Split(s, " "))

	// imprime "Simple String"
	fmt.Println(strings.Title(s))

	// imprime "simple string"; Se elimina todos los espacios en blanco
	// al principio y al final
	s = " simple string "
	fmt.Println(strings.TrimSpace(s))
}

// StringReader demuestra como crear
// una interface io.Reader rápidamente con un string
func StringReader() {
	s := "simple stringn"
	r := strings.NewReader(s)

	// prints s on Stdout
	io.Copy(os.Stdout, r)
}
