package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// CSVRecord contiene una fila analizada correctamente del archivo CSV.
type CSVRecord struct {
	SepalLength float64
	SepalWidth  float64
	PetalLength float64
	PetalWidth  float64
	Species     string
	ParseError  error
}

func main() {

	// Abrir el archivo dataset de iris.
	f, err := os.Open("./data/iris_mixed_types.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear un nuevo lector CSV que lea del archivo abierto.
	reader := csv.NewReader(f)

	// Crear un valor de slice que contendrá todos los registros analizados correctamente del CSV.
	var csvData []CSVRecord

	// line, ayudará a realizar un seguimiento del número de línea para el registro.
	line := 1

	// Leer en los registros buscando tipos inesperados.
	for {

		// Leer en una fila. Comprueba si estamos al final del archivo.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Crear un valor CSVRecord para la fila.
		var csvRecord CSVRecord

		// Analizar cada uno de los valores en el registro en función de un tipo esperado.
		for idx, value := range record {

			// Analizar el valor en el registro como string para la columna de sring.
			if idx == 4 {

				// Validar que el valor no sea una cadena vacía. Si el valor es una cadena vacía,
				// romper el ciclo de análisis.
				if value == "" {
					log.Printf("Parsing line %d failed, unexpected type in column %d\n", line, idx)
					csvRecord.ParseError = fmt.Errorf("Empty string value")
					break
				}

				// Agregar el valor de cadena a CSVRecord.
				csvRecord.Species = value
				continue
			}

			// De lo contrario, analizar el valor en el registro como float64.
			// floatValue mantendrá el valor flotante analizado del registro
			// para las columnas numéricas
			var floatValue float64

			if floatValue, err = strconv.ParseFloat(value, 64); err != nil {
				log.Printf("Parsing line %d failed, unexpected type in column %d\n", line, idx)
				csvRecord.ParseError = fmt.Errorf("Could not parse float")
				break
			}

			// Agregar el valor flotante al campo respectivo en CSVRecord.
			switch idx {
			case 0:
				csvRecord.SepalLength = floatValue
			case 1:
				csvRecord.SepalWidth = floatValue
			case 2:
				csvRecord.PetalLength = floatValue
			case 3:
				csvRecord.PetalWidth = floatValue
			}
		}

		// Append successfully parsed records to the slice defined above.
		// Agregar registros analizados correctamente al segmento definido anteriormente.
		if csvRecord.ParseError == nil {
			csvData = append(csvData, csvRecord)
		}

		// Incrementar el contador de línea.
		line++
	}

	fmt.Printf("successfully parsed %d lines\n", len(csvData))
}
