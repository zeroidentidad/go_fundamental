package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe" // https://github.com/go-gota/gota/wiki
	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/stat"
)

func main() {

	// Medidas de tendencia central:

	// Abrir archivo CSV
	irisFile, err := os.Open("./data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Crear un dataframe a partir del archivo CSV.
	irisDF := dataframe.ReadCSV(irisFile)

	// Obtener los valores flotantes de la columna "sepal_length",
	// ya que se analizará las medidas para esta variable.
	sepalLength := irisDF.Col("petal_length").Float()

	// Calcular la Media de la variable.
	meanVal := stat.Mean(sepalLength, nil)

	// Calcular la Moda de la variable.
	modeVal, modeCount := stat.Mode(sepalLength, nil)

	// Calcular Mediana de la variable.
	medianVal, err := stats.Median(sepalLength)
	if err != nil {
		log.Fatal(err)
	}

	// Salida de resultados a la salida estándar.
	fmt.Printf("\nResumen estadístico: Sepal Length:\n")
	fmt.Printf("Media, valor: %0.2f\n", meanVal)
	fmt.Printf("Moda, valor: %0.2f\n", modeVal)
	fmt.Printf("Moda, cuenta: %d\n", int(modeCount))
	fmt.Printf("Mediana, valor: %0.2f\n\n", medianVal)
}
