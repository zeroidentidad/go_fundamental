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

	// Resumen estadisticas e histograma ejemplo anterior

	// Open the CSV file.
	loanDataFile, err := os.Open("clean_loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer loanDataFile.Close()

	// Crear un marco de datos a partir del archivo CSV.
	loanDF := dataframe.ReadCSV(loanDataFile)

	// Usar el método Describe para calcular estadísticas de resumen
	// para todas las columnas de un golpe.
	loanSummary := loanDF.Describe()

	// Enviar las estadísticas de resumen a stdout.
	fmt.Println(loanSummary)

	// Crear un histograma para cada una de las columnas en el dataset.
	for _, colName := range loanDF.Names() {

		// Crear un valor de plotter.Values y llenarlo con los valores
		// de la columna respectiva del marco de datos.
		plotVals := make(plotter.Values, loanDF.Nrow())
		for i, floatVal := range loanDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		// Hacer una trama y establecer su título.
		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histograma de %s", colName)

		// Crear histograma de valores.
		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}

		// Normalizar histograma.
		h.Normalize(1)

		// Agregar histograma a la trama.
		p.Add(h)

		// Guardar la trama en un archivo PNG.
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}
}
