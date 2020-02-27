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

	// Ajuste y evaluación de modelo AR(2) parte 1

	// Abrir archivo CSV.
	passengersFile, err := os.Open("log_diff_series.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// Crear marco de datos del archivo CSV.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// Obtener tiempos y pasajeros en slice de floats
	passengers := passengersDF.Col("log_differenced_passengers").Float()

	// Calcular los coeficientes para el retraso 1, 2 y error.
	coeffs, intercept := autoregressive(passengers, 2)

	// Salida del modelo AR(2) a stdout.
	fmt.Printf("\nlog(x(t)) - log(x(t-1)) = %0.6f + lag1*%0.6f + lag2*%0.6f\n\n", intercept, coeffs[0], coeffs[1])
}

// autoregressive calcula un modelo AR para una serie en un orden dado.
func autoregressive(x []float64, lag int) ([]float64, float64) {

	// Crear un valor regresssion.Regression necesario
	// para entrenar un modelo usando github.com/sajari/regression.
	var r regression.Regression
	r.SetObserved("x")

	// Definir retraso actual y todos los retrasos intermedios.
	for i := 0; i < lag; i++ {
		r.SetVar(i, "x"+strconv.Itoa(i))
	}

	// Cambiar la serie
	xAdj := x[lag:len(x)]

	// Recorrer la serie creando conjunto de datos para la regresión.
	for i, xVal := range xAdj {

		// Pasar sobre retrasos intermedios para construir variables independientes.
		laggedVariables := make([]float64, lag)
		for idx := 1; idx <= lag; idx++ {

			// Obtener variables de series retrasadas.
			laggedVariables[idx-1] = x[lag+i-idx]
		}

		// Agregar puntos al valor de regresión.
		r.Train(regression.DataPoint(xVal, laggedVariables))
	}

	// Ajustar regresión.
	r.Run()

	// coeff guardara los coeficientes para retrasos.
	var coeff []float64
	for i := 1; i <= lag; i++ {
		coeff = append(coeff, r.Coeff(i))
	}

	return coeff, r.Coeff(0)
}
