package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sajari/regression"
)

func main() {

	// Abrir archivo de dataset de entrenamiento
	f, err := os.Open("./data/training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear un nuevo lector CSV que lea del archivo abierto.
	reader := csv.NewReader(f)

	// Leer en todos los registros CSV
	reader.FieldsPerRecord = 4
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// En este caso, se intentará modelar las Ventas por las
	// características de TV y Radio más una intercepción.
	var r regression.Regression
	r.SetObserved("Sales")
	r.SetVar(0, "TV")
	r.SetVar(1, "Radio")

	// Recorrer los registros CSV agregando los datos de entrenamiento.
	for i, record := range trainingData {

		// Saltar encabezado
		if i == 0 {
			continue
		}

		// Analizar las Ventas
		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Analizar el valor de TV.
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Analizar el valor de Radio.
		radioVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Agregar estos puntos al valor de regresión.
		r.Train(regression.DataPoint(yVal, []float64{tvVal, radioVal}))
	}

	// Entrenar/ajustar el modelo de regresión.
	r.Run()

	// Salida de los parámetros del modelo entrenado.
	fmt.Printf("\nRegression Formula:\n%v\n\n", r.Formula)
}
