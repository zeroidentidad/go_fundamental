package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	// Abrir archivo dataset de prueba
	f, err := os.Open("./data/test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear un nuevo lector CSV que lea del archivo abierto.
	reader := csv.NewReader(f)

	// Leer en todos los registros CSV
	reader.FieldsPerRecord = 4
	testData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Recorrer los datos de reserva que predicen "y", y evaluar
	// la predicción con el error absoluto medio.
	var mAE float64
	for i, record := range testData {

		// Saltar el encabezado.
		if i == 0 {
			continue
		}

		// Analizar las ventas.
		yObserved, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Analizar el valor de TV.
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Analiza el valor de Radio.
		radioVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Analizar el valor de Newspaper.
		newspaperVal, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Predecir "y" con el modelo entrenado.
		yPredicted := predict(tvVal, radioVal, newspaperVal)

		// Agregar al error absoluto medio.
		mAE += math.Abs(yObserved-yPredicted) / float64(len(testData))
	}

	// Salida de MAE a la salida estándar
	fmt.Printf("\nMAE = %0.2f\n\n", mAE)
}

// predict utiliza el modelo de regresión entrenado para hacer una predicción
// basada en un valor de TV, Radio y Newspaper.
func predict(tv, radio, newspaper float64) float64 {
	return 3.038 + tv*0.047 + 0.177*radio + 0.001*newspaper
}
