package main

import (
	"encoding/csv"
	"fmt"
	"io"
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
	reader.FieldsPerRecord = -1

	// rawCSVData mantendrá filas analizadas con éxito.
	var rawCSVData [][]string

	// Leer en los registros uno por uno.
	for {

		// Leer en una fila. Comprueba si estamos al final del archivo.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Agregar el registro a nuestro conjunto de datos.
		rawCSVData = append(rawCSVData, record)
	}

	fmt.Println(rawCSVData)
}
