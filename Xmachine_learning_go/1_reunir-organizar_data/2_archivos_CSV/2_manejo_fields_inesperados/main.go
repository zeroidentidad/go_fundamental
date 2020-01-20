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
	f, err := os.Open("./data/iris_unexpected_fields.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear un nuevo lector CSV que lea del archivo abierto.
	reader := csv.NewReader(f)

	// Deberíamos tener 5 campos por línea. Al establecer FieldsPerRecord en 5,
	// podemos validar que cada una de las filas de nuestro CSV
	// tenga el número correcto de campos.
	reader.FieldsPerRecord = 5

	// rawCSVData mantendrá filas analizadas con éxito.
	var rawCSVData [][]string

	// Leer en los registros uno por uno.
	for {

		// Leer en una fila. Comprueba si estamos al final del archivo.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Si tuvimos un error de análisis, registrar el error y continuar.
		if err != nil {
			log.Println(err)
			continue
		}

		// Agregar el registro al data set, si tiene el el número esperado de campos
		rawCSVData = append(rawCSVData, record)
	}

	fmt.Printf("analizadas %d líneas con éxito\n", len(rawCSVData))
}
