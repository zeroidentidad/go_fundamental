package main

import (
	"encoding/csv"
	"image/color"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	// transformar a series estacionarias - reajuste diff con log func

	// Abrir archivo CSV
	passengersFile, err := os.Open("./data/AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// Crear marco de datos a partir del archivo CSV.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// Extraer el número pasajeros y columnas de tiempo
	// como slice de flotantes
	passengerVals := passengersDF.Col("AirPassengers").Float()
	timeVals := passengersDF.Col("time").Float()

	// pts mantendrá los valores para trazar.
	pts := make(plotter.XYs, passengersDF.Nrow()-1)

	// differenced mantendrá valores diferenciados que se enviarán a nuevo archivo CSV.
	var differenced [][]string
	differenced = append(differenced, []string{"time", "log_differenced_passengers"})

	// Llenar pts con data.
	for i := 1; i < len(passengerVals); i++ {
		pts[i-1].X = timeVals[i]
		pts[i-1].Y = math.Log(passengerVals[i]) - math.Log(passengerVals[i-1])
		differenced = append(differenced, []string{
			strconv.FormatFloat(timeVals[i], 'f', -1, 64),
			strconv.FormatFloat(math.Log(passengerVals[i])-math.Log(passengerVals[i-1]), 'f', -1, 64), //diff ajust
		})
	}

	// Crear la trama
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "time"
	p.Y.Label.Text = "log(differenced passengers)"
	p.Add(plotter.NewGrid())

	// Agregar puntos de trazado de línea para la serie de tiempo.
	l, err := plotter.NewLine(pts)
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	// Guardar trama en archivo PNG.
	p.Add(l)
	if err := p.Save(10*vg.Inch, 4*vg.Inch, "log_diff_passengers_ts.png"); err != nil {
		log.Fatal(err)
	}

	// Guardar datos diferenciados en nuevo CSV.
	f, err := os.Create("log_diff_series.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.WriteAll(differenced)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
