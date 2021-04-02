package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

/* Ej. uso: */
// 1.1 Run compilado -> echo "contar las palabras echo pipe" | ./binario
// 1.2 Run compilado -> ./binario archivo.txt
// 2.1 Run fuente -> echo "contar las palabras echo pipe" | go run .
// 2.2 Run fuente -> go run . archivo.txt
// Ayuda -> go run . -help

func main() {
	// Definición de una flag boolean -l para contar líneas en lugar de palabras
	lines := flag.Bool("l", false, "Contar líneas")
	flag.Parse()

	// Si se seteo argumento adicional, asumir que es nombre de archivo
	if fname := flag.Arg(0); fname != "" {
		// Abrir el archivo para leer
		fd, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer fd.Close()

		// Imprimir recuento palabras/líneas en el archivo
		fmt.Println(counter(fd, *lines))
		return
	}

	// En caso de no leer archivo sino entrada en pipe
	// Imprimir recuento palabras/líneas en el archivo
	fmt.Println(counter(os.Stdin, *lines))
}

func counter(r io.Reader, countLines bool) int {
	// Se utiliza scanner para leer texto de un Reader (como archivos)
	scanner := bufio.NewScanner(r)

	// Si flag contar líneas no está seteada, entonces contar palabras, así que
	// el scanner divida en palabras (valor predeterminado está dividido por líneas)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	// Definir un contador
	wc := 0
	// Por cada palabra o línea escaneada, agregar 1 al contador
	for scanner.Scan() {
		wc++
	}

	// Return total
	return wc

}
