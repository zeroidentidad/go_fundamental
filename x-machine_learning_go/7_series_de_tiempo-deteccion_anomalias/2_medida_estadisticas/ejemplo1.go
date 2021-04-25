package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/gonum/stat"
)

func main() {

	// obtener autocorrelaciones parte 1

	// abrir archivo CSV
	passengersFile, err := os.Open("./data/AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// Crear marco de datos a partir del archivo CSV.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// Obtener tiempos y pasajeros como slice de flotantes
	passengers := passengersDF.Col("AirPassengers").Float()

	// Recorrer sobre varios valores de retraso en la serie.
	fmt.Println("Autocorrelation:")
	for i := 1; i < 11; i++ {

		// Calcular la autocorrelacion
		ac := acf(passengers, i)
		fmt.Printf("Lag %d period: %0.2f\n", i, ac)
	}
}

// acf calcula la autocorrelación para una serie en el retraso dado.
func acf(x []float64, lag int) float64 {

	// Cambiar serie.
	xAdj := x[lag:len(x)]
	xLag := x[0 : len(x)-lag]

	// numerator tendrá el numerador acumulado,
	// y denominator tendrá el denominador acumulado.
	var numerator float64
	var denominator float64

	// Calcular la media valores x, que se utilizarán en cada término de la autocorrelación.
	xBar := stat.Mean(x, nil)

	// Calcular el numerador.
	for idx, xVal := range xAdj {
		numerator += ((xVal - xBar) * (xLag[idx] - xBar))
	}

	// Calcular el denominador.
	for _, xVal := range x {
		denominator += math.Pow(xVal-xBar, 2)
	}

	return numerator / denominator
}
