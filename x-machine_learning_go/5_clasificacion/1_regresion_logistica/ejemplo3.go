package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Ejemplo simplificado de regresion logistica

	// Abrir archivo de dataset de préstamo.
	f, err := os.Open("./data/loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear un nuevo lector CSV que lea del archivo abierto.
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 2

	// Leer en todos los registros CSV
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Crear archivo de salida.
	f, err = os.Create("clean_loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear escritor CSV.
	w := csv.NewWriter(f)

	// Mover secuencialmente las filas escribiendo los valores analizados.
	for idx, record := range rawCSVData {

		// Saltar fila de encabezado.
		if idx == 0 {

			// Escribir encabezado en el archivo de salida.
			if err := w.Write([]string{"FICO_score", "class"}); err != nil {
				log.Fatal(err)
			}
			continue
		}

		// Inicializar slice para contener valores analizados.
		outRecord := make([]string, 2)

		// Analizar y normalizar el puntaje FICO.
		score, err := strconv.ParseFloat(strings.Split(record[0], "-")[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		outRecord[0] = strconv.FormatFloat((score-640.0)/(830.0-640.0), 'f', 4, 64)

		// Analizar la clase de tasa de interés.
		rate, err := strconv.ParseFloat(strings.TrimSuffix(record[1], "%"), 64)
		if err != nil {
			log.Fatal(err)
		}

		if rate <= 12.0 {
			outRecord[1] = "1.0"

			// Escribir registro en archivo de salida.
			if err := w.Write(outRecord); err != nil {
				log.Fatal(err)
			}
			continue
		}

		outRecord[1] = "0.0"

		// Escribir registro en archivo de salida.
		if err := w.Write(outRecord); err != nil {
			log.Fatal(err)
		}
	}

	// Escribir datos almacenados en el escritor subyacente (salida estándar).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
