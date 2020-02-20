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

	// representacion grafica y conversion de valores

	// Abrir archivo CSV
	passengersFile, err := os.Open("./data/AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// Crear marco de datos a partir del archivo CSV.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// Extraer columna de número de pasajeros.
	yVals := passengersDF.Col("AirPassengers").Float()

	// pts mantendrá valores para trazar.
	pts := make(plotter.XYs, passengersDF.Nrow())

	// Llenar pts con data.
	for i, floatVal := range passengersDF.Col("time").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
	}

	// Crear la trama.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "time"
	p.Y.Label.Text = "passengers"
	p.Add(plotter.NewGrid())

	// Agregar los puntos de trazado de línea para la serie de tiempo.
	l, err := plotter.NewLine(pts)
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	// Guardar la trama en un archivo PNG.
	p.Add(l)
	if err := p.Save(10*vg.Inch, 4*vg.Inch, "passengers_ts.png"); err != nil {
		log.Fatal(err)
	}
}
