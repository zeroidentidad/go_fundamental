package main

import (
	"image/color"
	"log"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	// Crear una nueva trama.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.Title.Text = "Logistic Function"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "f(x)"

	// Crear la función del trazador.
	logisticPlotter := plotter.NewFunction(func(x float64) float64 { return logistic(x) })
	logisticPlotter.Color = color.RGBA{B: 255, A: 255}

	// Agregar la función del trazador al diagrama.
	p.Add(logisticPlotter)

	// Establecer los rangos de eje. A diferencia de otros conjuntos de datos,
	// las funciones no establecen los rangos de eje automáticamente
	// ya que las funciones no necesariamente tienen un rango finito de valores x e y.
	p.X.Min = -10
	p.X.Max = 10
	p.Y.Min = -0.1
	p.Y.Max = 1.1

	// Save the plot to a PNG file.
	// Guardar la trama en un archivo PNG.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "logistic.png"); err != nil {
		log.Fatal(err)
	}
}

// Logistic implementa la función logistic,
// que se utiliza en la regresión logística.
func logistic(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}
