package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	tFile, err := ioutil.TempFile("", "gotmpfile")
	if err != nil {
		panic(err)
	}
	// Llamada diferida responsable de manejar la limpieza.
	defer os.Remove(tFile.Name())

	fmt.Println(tFile.Name())

	// TempDir devuelve la ruta en cadena de texto.
	tDir, err := ioutil.TempDir("", "gotmpdir")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tDir)

	fmt.Println(tDir)

}
