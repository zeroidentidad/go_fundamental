package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	// Esta llamada imprimirá
	// los argumentos de la línea de comando.
	fmt.Println(args)

	// El primer argumento, elemento cero de la matriz,
	// es el nombre del binario llamado.
	programName := args[0]
	fmt.Printf("El nombre binario es: %s \n", programName)

	// El resto de los argumentos podrían obtenerse
	// omitiendo el primer argumento.
	otherArgs := args[1:]
	fmt.Println(otherArgs)

	for idx, arg := range otherArgs {
		fmt.Printf("Arg %d = %s \n", idx, arg)
	}

}
