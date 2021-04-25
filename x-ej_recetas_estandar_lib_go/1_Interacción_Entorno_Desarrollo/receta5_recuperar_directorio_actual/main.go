package main

import (
	"fmt"
	"os"
	"path/filepath" // <-
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
		//detener la ejecución al haber un error y mostrarlo
	}

	// Ruta al archivo ejecutable
	fmt.Println(ex)

	// Resolver el directorio
	// del ejecutable
	exPath := filepath.Dir(ex)
	fmt.Println("Ruta del ejecutable:" + exPath)

	// Se usa EvalSymlinks para obtener la ruta real.
	realPath, err := filepath.EvalSymlinks(exPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Symlink evaludado:" + realPath)
}
