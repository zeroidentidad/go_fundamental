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
Esto es fan out?
Respuesta:
Si.
-- SI
¿Estamos trabajando "fanning out"? Sí. Hemos lanzado varios goroutines que están publicando simultáneamente un mensaje en nuestro canal. El blog de Golang dice: "Fan out significa que tienes varias funciones que leen desde el mismo canal hasta que se cierra ese canal". Aquí tenemos múltiples funciones de lectura desde el mismo canal. Así que, bueno, estamos fanning out.


CHALLENGE #2
Esto es fan in?
No.
¿Qué se está "fanned in" aquí? Hemos lanzado varios goroutines de la misma función: workerProcess. ¿Qué hacen esos goroutines? Todos están leyendo de un canal sin búfer. Si hubiera una tremenda cantidad de procesamiento que cada función "workerProcess" ejecutara, entonces las tres funciones "workerProcess" podrían procesarse en paralelo: extraer valores del canal y procesarlos. No hay "fanning in" aunque aquí. Recuerda lo que describe el blog de golang en: "Una función puede leer desde varias entradas y continuar hasta que todas se cierren multiplexando los canales de entrada en un solo canal que se cierra cuando todas las entradas están cerradas". No tenemos muchos canales que convergen en un solo canal.
*/
