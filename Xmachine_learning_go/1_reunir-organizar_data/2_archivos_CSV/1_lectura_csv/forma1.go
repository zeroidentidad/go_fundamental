package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	// Abrir el archivo dataset de iris.
	f, err := os.Open("./data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear un nuevo lector CSV que lea del archivo abierto.
	reader := csv.NewReader(f)

	// Supongamos que no sabemos la cantidad de campos por línea.
	// Al establecer FieldsPerRecord negativo,
	// cada fila puede tener un número variable de campos.
	reader.FieldsPerRecord = -1

	// Leer en todos los registros CSV.
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rawCSVData)
}
