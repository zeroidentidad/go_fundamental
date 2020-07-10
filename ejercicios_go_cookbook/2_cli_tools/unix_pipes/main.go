package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// WordCount toma un archivo y devuelve un mapa
// con cada palabra como clave y su número de
// apariencias como valor
func WordCount(f io.Reader) map[string]int {
	result := make(map[string]int)

	// crea un escáner para trabajar en el archivo
	// con la interfaz io.Reader
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		result[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return result
}

func main() {
	fmt.Printf("string: numero_de_ocurrencias\n\n")
	for key, value := range WordCount(os.Stdin) {
		fmt.Printf("%s: %d\n", key, value)
	}
}
