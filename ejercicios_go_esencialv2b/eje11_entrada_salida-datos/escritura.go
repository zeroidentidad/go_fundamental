package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	contenido := []byte("Un mensaje para ti daaaa")

	datos := ioutil.WriteFile("info.txt", contenido, 0755)

	mostrarError(datos)

	fmt.Println("info guardada")
}

func mostrarError(e error) {
	if e != nil {
		panic(e)
	}
}
