package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	// Abrir los datos de las observaciones y predicciones continuas.
	f, err := os.Open("./data/continuous_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear un nuevo lector CSV que lea del archivo abierto.
	reader := csv.NewReader(f)

	// observado y predicho contendrá los valores analizados y pronosticados analizados del archivo de datos continuo.
	var observed []float64
	var predicted []float64

	// La línea rastreará los números de fila para el registro.
	line := 1

	// Leer en los registros buscando tipos inesperados en las columnas.
	for {

		// Leer en una fila. Comprueba si estamos al final del archivo.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Saltar el encabezado
		if line == 1 {
			line++
			continue
		}

		// Leer en los valores observados y pronosticados.
		observedVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		// Agregar el registro al slice, si tiene el tipo esperado.
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	// Calcular el error absoluto medio y el error cuadrático medio.
	var mAE float64
	var mSE float64
	for idx, oVal := range observed {
		mAE += math.Abs(oVal-predicted[idx]) / float64(len(observed))
		mSE += math.Pow(oVal-predicted[idx], 2) / float64(len(observed))
	}

	// Salida del valor MAE y MSE a la salida estándar.
	fmt.Printf("\nMAE = %0.2f\n", mAE)
	fmt.Printf("\nMSE = %0.2f\n\n", mSE)

}
