package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {

	// Pull del archivo CSV.
	irisFile, err := os.Open("./data/iris_labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Crear un dataframe a partir del archivo CSV.
	// Se inferirán los tipos de las columnas.
	irisDF := dataframe.ReadCSV(irisFile)

	// Crear un filtro para el dataframe.
	filter := dataframe.F{
		Colname:    "species",
		Comparator: "==",
		Comparando: "Iris-versicolor",
	}

	// Filtrar dataframe para ver solo las filas donde la especie de iris es "Iris-versicolor".
	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Err != nil {
		log.Fatal(versicolorDF.Err)
	}

	// Salida de los resultados a la salida estándar.
	fmt.Println(versicolorDF)

	// Filtrar dataframe nuevamente, pero solo seleccionar las columnas sepal_width y species.
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"})
	fmt.Println(versicolorDF)

	// Filtrar y seleccionar el marco de datos nuevamente, pero solo muestre los primeros tres resultados.
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"}).Subset([]int{0, 1, 2})
	fmt.Println(versicolorDF)

}
