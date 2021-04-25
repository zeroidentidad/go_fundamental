package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	// Abrir el CSV.
	f, err := os.Open("file.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Leer en los registros CSV.
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Obtener el valor máximo en la columna de num int.
	var intMax int
	for _, record := range records {
		// Analizar conversion valor entero.
		intVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}
		// Reemplazar valor máximo si es apropiado.
		if intVal > intMax {
			intMax = intVal
		}
	}

	// imprimir el valor máximo.
	fmt.Println(intMax)
}
