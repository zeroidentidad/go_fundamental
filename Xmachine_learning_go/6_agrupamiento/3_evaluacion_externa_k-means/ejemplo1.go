package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	// perfilar data parte 1

	// Abrir archivo CSV.
	driverDataFile, err := os.Open("./data/fleet_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer driverDataFile.Close()

	// Crear marco de datos a partir del archivo CSV.
	driverDF := dataframe.ReadCSV(driverDataFile)

	// Usar método Describe para calcular estadísticas de resumen para todas las columnas de una sola vez.
	driverSummary := driverDF.Describe()

	// Enviar las estadísticas de resumen a stdout.
	fmt.Println(driverSummary)

	// Crear un histograma para cada una de las columnas del conjunto de datos.
	for _, colName := range driverDF.Names() {

		// Crear valor de plotter.Values y llenarlo con
		// los valores de la columna respectiva del marco de datos.
		plotVals := make(plotter.Values, driverDF.Nrow())
		for i, floatVal := range driverDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		// Hacer una trama y establecer su título.
		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histograma de %s", colName)

		// Crear un histograma de los valores.
		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}

		// Normalizar el histograma.
		h.Normalize(1)

		// Agregar el histograma a la trama.
		p.Add(h)

		// Guardar la trama en un archivo PNG.
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}
}
