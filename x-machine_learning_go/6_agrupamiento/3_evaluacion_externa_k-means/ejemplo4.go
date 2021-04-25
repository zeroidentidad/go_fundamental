package main

import (
	"image/color"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	// evaluar agrupamientos generados visualmente parte 1

	// Abrir archivo conjunto de datos conductores.
	f, err := os.Open("./data/fleet_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear marco de datos a partir del archivo CSV.
	driverDF := dataframe.ReadCSV(f)

	// Extraer la columna de distancia.
	yVals := driverDF.Col("Distance_Feature").Float()

	// clusterOne y clusterTwo mantendrán los valores para el trazado.
	var clusterOne [][]float64
	var clusterTwo [][]float64

	// Rellenar los grupos con datos.
	for i, xVal := range driverDF.Col("Speeding_Feature").Float() {
		distanceOne := floats.Distance([]float64{yVals[i], xVal}, []float64{50.05, 8.83}, 2)
		distanceTwo := floats.Distance([]float64{yVals[i], xVal}, []float64{180.02, 18.29}, 2)
		if distanceOne < distanceTwo {
			clusterOne = append(clusterOne, []float64{xVal, yVals[i]})
			continue
		}
		clusterTwo = append(clusterTwo, []float64{xVal, yVals[i]})
	}

	// pts * mantendrá los valores para trazar
	ptsOne := make(plotter.XYs, len(clusterOne))
	ptsTwo := make(plotter.XYs, len(clusterTwo))

	// Rellenar pts con datos.
	for i, point := range clusterOne {
		ptsOne[i].X = point[0]
		ptsOne[i].Y = point[1]
	}

	for i, point := range clusterTwo {
		ptsTwo[i].X = point[0]
		ptsTwo[i].Y = point[1]
	}

	// Crear la trama
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "Speeding"
	p.Y.Label.Text = "Distance"
	p.Add(plotter.NewGrid())

	sOne, err := plotter.NewScatter(ptsOne)
	if err != nil {
		log.Fatal(err)
	}
	sOne.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	sOne.GlyphStyle.Radius = vg.Points(3)

	sTwo, err := plotter.NewScatter(ptsTwo)
	if err != nil {
		log.Fatal(err)
	}
	sTwo.GlyphStyle.Color = color.RGBA{B: 255, A: 255}
	sTwo.GlyphStyle.Radius = vg.Points(3)

	// Guardar la trama en un archivo PNG.
	p.Add(sOne, sTwo)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "fleet_data_clusters.png"); err != nil {
		log.Fatal(err)
	}
}
