package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

func main() {

	// Medidas de propagación o dispersión.

	// Abrir archivo CSV
	irisFile, err := os.Open("./data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Crear dataframe de archivo CSV.
	irisDF := dataframe.ReadCSV(irisFile)

	// Obtener los valores flotantes de la columna "sepal_length"
	// ya que veremos las medidas para esta variable.
	sepalLength := irisDF.Col("petal_length").Float()

	// Calcular la Minima de la variable.
	minVal := floats.Min(sepalLength)

	// Calcular la Maxima de la variable.
	maxVal := floats.Max(sepalLength)

	// Calcular la Mediana de la variable.
	rangeVal := maxVal - minVal

	// Calcular la varianza de la variable.
	varianceVal := stat.Variance(sepalLength, nil)

	// Calcular la desviacion estandar de la variable.
	stdDevVal := stat.StdDev(sepalLength, nil)

	// Ordenar los valores
	inds := make([]int, len(sepalLength))
	floats.Argsort(sepalLength, inds)

	// Obtener los Cuantiles.
	quant25 := stat.Quantile(0.25, stat.Empirical, sepalLength, nil)
	quant50 := stat.Quantile(0.50, stat.Empirical, sepalLength, nil)
	quant75 := stat.Quantile(0.75, stat.Empirical, sepalLength, nil)

	// Salida de los resultados a salida estándar.
	fmt.Printf("\nResumen estadístico Sepal Length:\n")
	fmt.Printf("Maxima valor: %0.2f\n", maxVal)
	fmt.Printf("Minima valor: %0.2f\n", minVal)
	fmt.Printf("Rango valor: %0.2f\n", rangeVal)
	fmt.Printf("Varianza valor: %0.2f\n", varianceVal)
	fmt.Printf("Desviacion estandar valor: %0.2f\n", stdDevVal)
	fmt.Printf("Cuantil 25: %0.2f\n", quant25)
	fmt.Printf("Cuantil 50: %0.2f\n", quant50)
	fmt.Printf("Cuantil 75: %0.2f\n\n", quant75)
}
