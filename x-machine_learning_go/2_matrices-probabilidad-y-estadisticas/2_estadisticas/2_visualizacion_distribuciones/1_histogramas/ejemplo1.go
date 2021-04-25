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

	// Abrir archivo CSV.
	irisFile, err := os.Open("./data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Crear dataframe de archivo CSV.
	irisDF := dataframe.ReadCSV(irisFile)

	// Crear un histograma para cada una de las columnas de características del dataset
	for _, colName := range irisDF.Names() {

		// Si la columna es una de las columnas de características, crear un histograma de los valores.
		if colName != "species" {

			// Crear valor plotter.Values y llenarlo con los
			// valores de la columna respectiva del dataframe.
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}

			// Hacer una trama y establecer su título.
			p, err := plot.New()
			if err != nil {
				log.Fatal(err)
			}
			p.Title.Text = fmt.Sprintf("Histograma de %s", colName)

			// Crear histograma de valores extraídos del estándar normal.
			h, err := plotter.NewHist(v, 16)
			if err != nil {
				log.Fatal(err)
			}

			// Normalizar el histograma.
			h.Normalize(1)

			// Agregar histograma al plot.
			p.Add(h)

			// Guardar plot en un archivo PNG.
			if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
				log.Fatal(err)
			}
		}
	}
}
