package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	datos, errorLectura := ioutil.ReadFile("archivo.txt")

	mostrarError(errorLectura)

	fmt.Println(string(datos))
}

func mostrarError(e error) {
	if e != nil {
		panic(e)
	}
}
