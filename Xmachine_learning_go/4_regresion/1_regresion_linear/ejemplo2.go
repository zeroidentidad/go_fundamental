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

	// Abrir archivo de dataset.
	f, err := os.Open("./data/Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear un marco de datos a partir del archivo CSV.
	advertDF := dataframe.ReadCSV(f)

	// Crear un histograma para cada una de las columnas del conjunto de datos.
	for _, colName := range advertDF.Names() {

		// Crear un valor de plotter.Values y llenar con los valores de la columna
		// respectiva del marco de datos.
		plotVals := make(plotter.Values, advertDF.Nrow())
		for i, floatVal := range advertDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		// Hacer una trama y establecer su título.
		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histograma de %s", colName)

		// Crear un histograma de valores extraídos del estándar normal.
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
