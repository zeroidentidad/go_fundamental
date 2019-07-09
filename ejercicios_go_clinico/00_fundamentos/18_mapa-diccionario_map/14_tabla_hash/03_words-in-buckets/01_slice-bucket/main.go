package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get el libro las aventuras de sherlock holmes
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	// escanear la pagina
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set funcion split para la operacion de escaneo.
	scanner.Split(bufio.ScanWords)
	// Crear slice de slice, de string para obtener slices de palabras
	buckets := make([][]string, 12)
	// Crear slices obtener las palabras
	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}
	// Bucle sobre las palabras
	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}
	// Print len de cada cubo
	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}
	// Imprimir las palabras en uno de los cubos.
	// fmt.Println (buckets [6])
}

func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
	// comenta lo anterior, luego descomenta lo siguiente
	// una distribución más desigual
	// return len(word) % buckets
}
