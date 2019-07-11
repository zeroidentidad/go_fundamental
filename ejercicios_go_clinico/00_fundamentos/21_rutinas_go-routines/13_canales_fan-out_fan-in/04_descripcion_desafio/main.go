package main

import (
	"fmt"
	"time"
)

var workerID int
var publisherID int

func main() {
	entrada := make(chan string)
	go workerProcess(entrada)
	go workerProcess(entrada)
	go workerProcess(entrada)
	go publisher(entrada)
	go publisher(entrada)
	go publisher(entrada)
	go publisher(entrada)
	time.Sleep(1 * time.Millisecond)
}

// publisher envia datos en un canal
func publisher(salida chan string) {
	publisherID++
	thisID := publisherID
	dataID := 0
	for {
		dataID++
		fmt.Printf("publisher %d está enviando datos\n", thisID)
		data := fmt.Sprintf("Data del publisher %d. Data %d", thisID, dataID)
		salida <- data
	}
}

func workerProcess(in <-chan string) {
	workerID++
	thisID := workerID
	for {
		fmt.Printf("%d: esperado entrada...\n", thisID)
		entrada := <-in
		fmt.Printf("%d: entrada es: %s\n", thisID, entrada)
	}
}

/*
CHALLENGE #1
Es este fan out?

CHALLENGE #2
Es este fan in?
*/
