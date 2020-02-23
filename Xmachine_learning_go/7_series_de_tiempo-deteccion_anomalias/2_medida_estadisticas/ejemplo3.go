package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/kniren/gota/dataframe"
	"github.com/sajari/regression"
)

func main() {

	// autocorrelacion parcial

	// Abrir archivo CSV.
	passengersFile, err := os.Open("./data/AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// Crear marco de datos a partir de archivo CSV.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// Obtener tiempo y pasajeros como slice de flotantes.
	passengers := passengersDF.Col("AirPassengers").Float()

	// Recorrer varios valores de retraso en la serie.
	fmt.Println("Partial Autocorrelation:")
	for i := 1; i < 11; i++ {

		// Calcular la autocorrelación parcial.
		pac := pacf(passengers, i)
		fmt.Printf("Lag %d period: %0.2f\n", i, pac)
	}
}

// pacf calcula la autocorrelación parcial para una serie en el retraso dado.
func pacf(x []float64, lag int) float64 {

	// Crear valor regresssion.Regression necesario para
	// entrenar un modelo usando github.com/sajari/regression.
	var r regression.Regression
	r.SetObserved("x")

	// Definir retraso actual y todos los retrasos intermedios.
	for i := 0; i < lag; i++ {
		r.SetVar(i, "x"+strconv.Itoa(i))
	}

	// Cambiar serie.
	xAdj := x[lag:len(x)]

	// Recorrer la serie creando el conjunto de datos para la regresión.
	for i, xVal := range xAdj {

		// Recorrer los retrazos intermedios para acumular
		// variables independentes.
		laggedVariables := make([]float64, lag)
		for idx := 1; idx <= lag; idx++ {

			// Obtener las variables de series retrazadas.
			laggedVariables[idx-1] = x[lag+i-idx]
		}

		// Agregar estos puntos al valor de regresión.
		r.Train(regression.DataPoint(xVal, laggedVariables))
	}

	// Ajustar la regresión.
	r.Run()

	return r.Coeff(lag)
}
