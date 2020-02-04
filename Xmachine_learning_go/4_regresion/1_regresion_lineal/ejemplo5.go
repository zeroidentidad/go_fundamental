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

	// Entrenar modelo con conjunto de datos de entrenamiento (parte 4)

	// Abrir archivo de dataset de entrenamiento.
	f, err := os.Open("./newData/training.csv")
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

	// En este caso, se intentará modelar las Ventas (y) por la caracteristica de TV más una intercepción.
	// Como tal, se crea la estructura necesaria para entrenar un modelo usando github.com/sajari/regression.
	var r regression.Regression
	r.SetObserved("Sales")
	r.SetVar(0, "TV")

	// Bucle de registros en el CSV, agregando los datos de entrenamiento al valor de regresión.
	for i, record := range trainingData {

		// Saltar el encabezado
		if i == 0 {
			continue
		}

		// Analizar la medida de regresión de Ventas, o "y".
		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Analizar el valor para TV.
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Agregar estos puntos al valor de regresión.
		r.Train(regression.DataPoint(yVal, []float64{tvVal}))
	}

	// Entrenar/ajustar el modelo de regresión.
	r.Run()

	// Salida de los parámetros del modelo entrenado.
	fmt.Printf("\nRegression Formula:\n%v\n\n", r.Formula)
}
