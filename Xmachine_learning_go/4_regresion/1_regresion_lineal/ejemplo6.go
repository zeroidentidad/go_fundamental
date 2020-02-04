package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/sajari/regression"
)

func main() {

	// Evaluar modelo entrenado con conjunto de datos de prueba (parte 5)

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

	// Evaluar modelo entrenado con conjunto de datos de prueba (parte 5.1)
	/* ===================================================== */

	// Abrir archivo de conjunto de datos de prueba
	f, err = os.Open("./newData/test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear un lector CSV que lea desde el archivo abierto.
	reader = csv.NewReader(f)

	// Leer en todos los registros CSV
	reader.FieldsPerRecord = 4
	testData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Recorrer los datos de la prueba prediciendo "y"
	// y evaluando la predicción con el error absoluto medio.
	var mAE float64
	for i, record := range testData {

		// Saltar el encabezado
		if i == 0 {
			continue
		}

		// Analizar la medida de progresión de diabetes observada, o "y".
		yObserved, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Analizar el valor de bmi.
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Predecir con nuestro modelo entrenado
		yPredicted, err := r.Predict([]float64{tvVal})

		// Agregar al error absoluto medio.
		mAE += math.Abs(yObserved-yPredicted) / float64(len(testData))
	}

	// Salida de MAE a la salida estándar.
	fmt.Printf("MAE = %0.2f\n\n", mAE)
}
