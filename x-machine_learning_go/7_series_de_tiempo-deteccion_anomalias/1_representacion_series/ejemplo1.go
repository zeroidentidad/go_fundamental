package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {

	// Abrir archivo CSV
	passengersFile, err := os.Open("./data/AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// Crear marco de datos a partir del archivo CSV.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// Como verificación de integridad, mostrar los registros para stdout.
	// Gota formateará marco de datos para una buena impresión.
	fmt.Println(passengersDF)
}
