package main

import "fmt"

func main() {
	m := make(map[string]int)
	cambiaMe(m)
	fmt.Println(m["Jesus"]) // 26
}

func cambiaMe(z map[string]int) {
	z["Jesus"] = 26
	z["Veronica"] = 25
}

/*
Asignación con make
Vuelve a la asignación. La función incorporada make (T, args)
Sirve un propósito diferente de new(T). Crea slices, maps y channels solamente,
y devuelve un valor inicializado (no puesto a cero) de tipo T (no * T). La razón por
la distinción es que estos tres tipos representan, bajo las coberturas, referencias a estructuras de datos.
que deben ser inicializado antes de su uso. Una slice, por ejemplo, es un descriptor de tres elementos que contiene
un puntero a los datos (dentro de una matriz), la longitud y la capacidad, y hasta que esos elementos se inicialicen,
el slice es nulo. Para slices, maps y channels, make inicializa la estructura de datos interna y prepara
el valor para usar.
*/
// https://golang.org/doc/effective_go.html#allocation_make
