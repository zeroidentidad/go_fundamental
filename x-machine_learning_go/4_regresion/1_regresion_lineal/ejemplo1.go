package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {

	// Perfilando los datos (parte 1)

	// Abrir archivo CSV
	advertFile, err := os.Open("./data/Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer advertFile.Close()

	// Cree un marco de datos a partir del archivo CSV.
	advertDF := dataframe.ReadCSV(advertFile)

	// Usar el método Describe para calcular estadísticas de resumen para todas las columnas de una sola vez.
	advertSummary := advertDF.Describe()

	// Enviar resumen de estadísticas a stdout.
	fmt.Println(advertSummary)
}
