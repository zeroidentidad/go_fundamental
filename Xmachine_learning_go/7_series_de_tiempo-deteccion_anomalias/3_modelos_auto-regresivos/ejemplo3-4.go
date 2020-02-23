package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"github.com/kniren/gota/dataframe"
	"github.com/sajari/regression"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {

	// analisis autocorrelaciones
	diagrama_ACF()

	diagrama_PACF()

}

// Metodos diagramas para graficar calculos

func diagrama_ACF() {
	// Abrir archivo CSV.
	passengersFile, err := os.Open("./log_diff_series.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// Crear marco de datos a partir del archivo CSV.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// Obtener tiempos y pasajeros como slice de flotantes
	passengers := passengersDF.Col("log_differenced_passengers").Float()

	// Crear nueva trama, para trazar autocorrelaciones.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = "Autocorrelations for log(differenced passengers)"
	p.X.Label.Text = "Lag"
	p.Y.Label.Text = "ACF"
	p.Y.Min = 0
	p.Y.Max = 1

	w := vg.Points(3)

	// Crear puntos para trazar.
	numLags := 20
	pts := make(plotter.Values, numLags)

	// Recorrer valores de retraso en la serie.
	for i := 1; i <= numLags; i++ {

		// Calcular autocorrelación.
		pts[i-1] = acf(passengers, i)
	}

	// Agregar puntos a la trama.
	bars, err := plotter.NewBarChart(pts, w)
	if err != nil {
		log.Fatal(err)
	}
	bars.LineStyle.Width = vg.Length(0)
	bars.Color = plotutil.Color(1)

	// Guardar la trama en archivo PNG.
	p.Add(bars)
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "acf.png"); err != nil {
		log.Fatal(err)
	}
}

func diagrama_PACF() {
	// Abrir archivo CSV.
	passengersFile, err := os.Open("./log_diff_series.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// Crear marco de datos a partir del archivo CSV.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// Obtener tiempos y pasajeros como slice de flotantes
	passengers := passengersDF.Col("log_differenced_passengers").Float()

	// Crear nueva trama, para trazar autocorrelaciones.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = "Partial Autocorrelations for log(differenced passengers)"
	p.X.Label.Text = "Lag"
	p.Y.Label.Text = "PACF"
	p.Y.Min = 15
	p.Y.Max = -1

	w := vg.Points(3)

	// Crear puntos para trazar
	numLags := 20
	pts := make(plotter.Values, numLags)

	// Recorrer valores de retraso en la serie.
	for i := 1; i <= numLags; i++ {

		// Calcular autocorrelación parcial
		pts[i-1] = pacf(passengers, i)
	}

	// Agregar puntos a la trama.
	bars, err := plotter.NewBarChart(pts, w)
	if err != nil {
		log.Fatal(err)
	}
	bars.LineStyle.Width = vg.Length(0)
	bars.Color = plotutil.Color(1)

	// Guardar la trama en archivo PNG.
	p.Add(bars)
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "pacf.png"); err != nil {
		log.Fatal(err)
	}
}

// Funciones calculo atucorrelacion

// acf calcula la autocorrelación para una serie en el retraso dado.
func acf(x []float64, lag int) float64 {

	// Cambiar series
	xAdj := x[lag:len(x)]
	xLag := x[0 : len(x)-lag]

	// numerator mantendrá numerador acumulado, y
	// denominator mantendrá denominador acumulado.
	var numerator float64
	var denominator float64

	// Calcular la media de valores x,
	// que se utilizará en cada término de la autocorrelación.
	xBar := stat.Mean(x, nil)

	// Calcular numerador
	for idx, xVal := range xAdj {
		numerator += ((xVal - xBar) * (xLag[idx] - xBar))
	}

	//Calcular denomindador
	for _, xVal := range x {
		denominator += math.Pow(xVal-xBar, 2)
	}

	return numerator / denominator
}

// pacf calcula la autocorrelación parcial para una serie en el retraso dado.
func pacf(x []float64, lag int) float64 {

	// Crear un valor regresssion.Regression necesario para entrenar
	// un modelo usando github.com/sajari/regression.
	var r regression.Regression
	r.SetObserved("x")

	// Definir retraso actual y todos los retrasos intermedios.
	for i := 0; i < lag; i++ {
		r.SetVar(i, "x"+strconv.Itoa(i))
	}

	//Cambiar la serie
	xAdj := x[lag:len(x)]

	// Recorrer serie creando onjunto de datos para la regresión.
	for i, xVal := range xAdj {

		// Loop over the intermediate lags to build up
		// our independent variables.
		laggedVariables := make([]float64, lag)
		for idx := 1; idx <= lag; idx++ {

			// Obtener variables de series retrazadas.
			laggedVariables[idx-1] = x[lag+i-idx]
		}

		// Agregar puntos al valor de regresión.
		r.Train(regression.DataPoint(xVal, laggedVariables))
	}

	// Ajustar la regresión.
	r.Run()

	return r.Coeff(lag)
}
