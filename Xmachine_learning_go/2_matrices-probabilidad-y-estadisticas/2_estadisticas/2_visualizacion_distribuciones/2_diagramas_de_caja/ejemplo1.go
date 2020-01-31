package main

import (
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	// Abrir archivo CSV
	irisFile, err := os.Open("./data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Crear dataframe de archivo CSV
	irisDF := dataframe.ReadCSV(irisFile)

	// Crear el diagrama y establecer su título y etiqueta de eje.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = "Diagramas de caja"
	p.Y.Label.Text = "Valores"

	// Crear el cuadro para los datos.
	w := vg.Points(50)

	// Cree un diagrama de caja para cada una de las columnas de características del dataset.
	for idx, colName := range irisDF.Names() {

		// Si es una de las columnas de características, crear un diagrama de valores.
		if colName != "species" {

			// Crear valor plotter.Values y llenarlo con los
			// valores de la columna respectiva del dataframe.
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}

			// Agregar los datos al plot.
			b, err := plotter.NewBoxPlot(w, float64(idx), v)
			if err != nil {
				log.Fatal(err)
			}
			p.Add(b)
		}
	}

	// Establecer el eje X de la gráfica en nominal con los nombres dados para x=0, x=1, etc.
	p.NominalX("sepal_length", "sepal_width", "petal_length", "petal_width")

	if err := p.Save(6*vg.Inch, 8*vg.Inch, "boxplots.png"); err != nil {
		log.Fatal(err)
	}
}
