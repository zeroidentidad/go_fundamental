//R1: Recuperando la versión de Golang

package main

import (
	"log"
	"runtime"
)

const info = `
Aplicacion %s iniciada.
Binario construido en GO: %s`

// igual que en js funcionan los template literals

func main() {
	log.Printf(info, "Receta 1", runtime.Version())

}
