package main

import (
	"image/color"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	// perfilar data parte 2 - diagrama dispersion

	// Abrir archivo de conjunto de datos choferes.
	f, err := os.Open("./data/fleet_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear marco de datos a partir del archivo CSV.
	driverDF := dataframe.ReadCSV(f)

	// Extraer la columna de distancia.
	yVals := driverDF.Col("Distance_Feature").Float()

	// pts mantendrá los valores para trazar
	pts := make(plotter.XYs, driverDF.Nrow())

	// Llenar pts con data.
	for i, floatVal := range driverDF.Col("Speeding_Feature").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
	}

	// Crear la trama.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "Speeding"
	p.Y.Label.Text = "Distance"
	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(3)

	// Guardar la trama en un archivo PNG.
	p.Add(s)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "fleet_data_scatter.png"); err != nil {
		log.Fatal(err)
	}
}
