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
	// Crear map con una clave int
	// y un valor de otro map
	// con una key de string, la cual sera la palabra
	// y un valor de int, el cual sera el numero de veces que la palabra ocurre
	buckets := make(map[int]map[string]int)
	// Crear slices obtener las palabras
	for i := 0; i < 12; i++ {
		buckets[i] = make(map[string]int)
	}
	// Bucle sobre las palabras
	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(word, 12)
		buckets[n][word]++
	}
	// Print palabras en un cubo
	for k, v := range buckets[6] {
		fmt.Println(v, " \t- ", k)
	}
}

func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
