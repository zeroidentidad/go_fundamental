package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe" //github.com/go-gota/gota
)

func main() {

	// Abrir el archivo CSV.
	irisFile, err := os.Open("./data/iris_labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Crear un dataframe a partir del archivo CSV.
	// Se inferirán los tipos de las columnas.
	irisDF := dataframe.ReadCSV(irisFile)

	// Como verificación de sanidad, muestrar los registros para stdout.
	// Gota formateará el marco de datos para impresión.
	fmt.Println(irisDF)
}
