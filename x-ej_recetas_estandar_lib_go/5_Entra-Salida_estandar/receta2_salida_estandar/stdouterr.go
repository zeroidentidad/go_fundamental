package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// Simplemente escribir una cadena
	io.WriteString(os.Stdout,
		"Cadena de texto en salida estándar.\n")

	io.WriteString(os.Stderr,
		"Cadena de texto en salida de error estándar.\n")

	// Stdout/err implementa interfaz de writer
	buf := []byte{0xAF, 0xFF, 0xFE}
	for i := 0; i < 200; i++ {
		if _, e := os.Stdout.Write(buf); e != nil {
			panic(e)
		}
	}

	// El paquete fmt también puede usarse
	fmt.Fprintln(os.Stdout, "\n")
}
