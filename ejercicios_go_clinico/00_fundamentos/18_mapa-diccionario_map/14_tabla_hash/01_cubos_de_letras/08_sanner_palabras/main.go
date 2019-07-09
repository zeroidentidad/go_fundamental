package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get el libro moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// escanear la página
	// NewScanner toma un lector y res.Body implementa la interfaz del lector (por lo que es un lector)
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set la función de split para la operación de escaneo.
	scanner.Split(bufio.ScanWords)
	// Bucle sobre las palabras
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
