package main

import (
	"fmt"
	"net/http"
)

func main() {

	header := http.Header{}

	// Usando el encabezado como slice
	header.Set("Auth-X", "abcdef1234")
	header.Add("Auth-X", "defghijkl")
	fmt.Println(header)

	// recuperar slice de valores en el encabezado
	resSlice := header["Auth-X"]
	fmt.Println(resSlice)

	// obtener primer valor
	resFirst := header.Get("Auth-X")
	fmt.Println(resFirst)

	// reemplazar todos los valores existentes con este
	header.Set("Auth-X", "newvalue")
	fmt.Println(header)

	// Remover header
	header.Del("Auth-X")
	fmt.Println(header)

}
