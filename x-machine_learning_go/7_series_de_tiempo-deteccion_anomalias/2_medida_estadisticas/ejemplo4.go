package main

import (
	"log"
	"os"
	"strconv"

	"github.com/kniren/gota/dataframe"
	"github.com/sajari/regression"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {

	// diagrama autocorrelacion parcial

	// Abrir archivo CSV.
	passengersFile, err := os.Open("./data/AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// Crear marco de datos a partir de archivo CSV.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// Obtener tiempo y pasajeros como slice de flotantes.
	passengers := passengersDF.Col("AirPassengers").Float()

	// Crear nueva trama, para trazar autocorrelaciones.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = "Partial Autocorrelations for AirPassengers"
	p.X.Label.Text = "Lag"
	p.Y.Label.Text = "PACF"
	p.Y.Min = 15
	p.Y.Max = -1

	w := vg.Points(3)

	// Crear los puntos para trazar.
	numLags := 20
	pts := make(plotter.Values, numLags)

	// Recorrer varios valores de retraso en la serie.
	for i := 1; i <= numLags; i++ {

		// Calcular la autocorrelación parcial.
		pts[i-1] = pacf(passengers, i)
	}

	// Agrega los puntos a la trama.
	bars, err := plotter.NewBarChart(pts, w)
	if err != nil {
		log.Fatal(err)
	}
	bars.LineStyle.Width = vg.Length(0)
	bars.Color = plotutil.Color(1)

	// Guarde trama en archivo PNG.
	p.Add(bars)
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "pacf.png"); err != nil {
		log.Fatal(err)
	}
}

// pacf calcula la autocorrelación parcial para una serie en el retraso dado.
func pacf(x []float64, lag int) float64 {

	// Crear valor regresssion.Regression necesario para
	// entrenar un modelo usando github.com/sajari/regression.
	var r regression.Regression
	r.SetObserved("x")

	// Definir retraso actual y todos los retrasos intermedios.
	for i := 0; i < lag; i++ {
		r.SetVar(i, "x"+strconv.Itoa(i))
	}

	// Cambiar serie.
	xAdj := x[lag:len(x)]

	// Recorrer la serie creando el conjunto de datos para la regresión.
	for i, xVal := range xAdj {

		// Recorrer los retrazos intermedios para acumular
		// variables independentes.
		laggedVariables := make([]float64, lag)
		for idx := 1; idx <= lag; idx++ {

			// Obtener las variables de series retrazadas.
			laggedVariables[idx-1] = x[lag+i-idx]
		}

		// Agregar estos puntos al valor de regresión.
		r.Train(regression.DataPoint(xVal, laggedVariables))
	}

	// Ajustar la regresión.
	r.Run()

	return r.Coeff(lag)
}
