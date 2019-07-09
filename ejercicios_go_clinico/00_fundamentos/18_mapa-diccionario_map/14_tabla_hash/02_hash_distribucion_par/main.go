package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get del libro Las aventuras de Sherlock Holmes.
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	// escanear la pagina
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set la funcion split para la operación de escaneo.
	scanner.Split(bufio.ScanWords)
	// Crear slice para obtener cuentas
	cubos := make([]int, 12)
	// Bucle sobre las palabras
	for scanner.Scan() {
		n := hashCubo(scanner.Text(), 12)
		cubos[n]++
	}
	fmt.Println(cubos)
}

func hashCubo(palabra string, cubos int) int {
	var sum int
	for _, v := range palabra {
		sum += int(v)
	}
	return sum % cubos
}
