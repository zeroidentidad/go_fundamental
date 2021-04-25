package main

import (
	"log"
	"math"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {

	// obtener autocorrelaciones parte 2 - grafica

	// Abrir archivo CSV
	passengersFile, err := os.Open("./data/AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// Crear un marco de datos a partir del archivo CSV.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// Obtener tiempos y pasajeros como slice de flotantes
	passengers := passengersDF.Col("AirPassengers").Float()

	// Crear nueva trama, para trazar autocorrelaciones.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = "Autocorrelations for AirPassengers"
	p.X.Label.Text = "Lag"
	p.Y.Label.Text = "ACF"
	p.Y.Min = 0
	p.Y.Max = 1

	w := vg.Points(3)

	// Crear los puntos para trazar.
	numLags := 20
	pts := make(plotter.Values, numLags)

	// Recorrer varios valores de retraso en la serie.
	for i := 1; i <= numLags; i++ {

		// Calcular autocorrelación.
		pts[i-1] = acf(passengers, i)
	}

	// Agregar los puntos a la trama.
	bars, err := plotter.NewBarChart(pts, w)
	if err != nil {
		log.Fatal(err)
	}
	bars.LineStyle.Width = vg.Length(0)
	bars.Color = plotutil.Color(1)

	// Guardar la trama en un archivo PNG.
	p.Add(bars)
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "acf.png"); err != nil {
		log.Fatal(err)
	}
}

// acf calcula la autocorrelación para una serie en el retraso dado.
func acf(x []float64, lag int) float64 {

	// Cambiar serie.
	xAdj := x[lag:len(x)]
	xLag := x[0 : len(x)-lag]

	// numerator tendrá el numerador acumulado,
	// y denominator tendrá el denominador acumulado.
	var numerator float64
	var denominator float64

	// Calcular la media valores x, que se utilizarán en cada término de la autocorrelación.
	xBar := stat.Mean(x, nil)

	// Calcular el numerador.
	for idx, xVal := range xAdj {
		numerator += ((xVal - xBar) * (xLag[idx] - xBar))
	}

	// Calcular el denominador.
	for _, xVal := range x {
		denominator += math.Pow(xVal-xBar, 2)
	}

	return numerator / denominator
}
