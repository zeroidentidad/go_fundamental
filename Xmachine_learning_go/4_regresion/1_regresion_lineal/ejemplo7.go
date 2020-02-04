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

	// Evaluar mostrando diagrama de dispersion con formula de regresión predefinida (parte 6)

	// Abrir archivo de dataset fuente
	f, err := os.Open("./data/Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear un marco de datos a partir del archivo CSV.
	advertDF := dataframe.ReadCSV(f)

	// Extraer la columna de destino.
	yVals := advertDF.Col("Sales").Float()

	// pts mantendrá los valores para trazar.
	pts := make(plotter.XYs, advertDF.Nrow())

	// ptsPred mantendrá los valores predecidos para el trazado.
	ptsPred := make(plotter.XYs, advertDF.Nrow())

	// Rellenar pts con datos.
	for i, floatVal := range advertDF.Col("TV").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
		ptsPred[i].X = floatVal
		ptsPred[i].Y = predict(floatVal)
	}

	// Crea la trama.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "TV"
	p.Y.Label.Text = "Sales"
	p.Add(plotter.NewGrid())

	// Agregue los puntos del diagrama de dispersión para las observaciones.
	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(3)

	// Agregar los puntos de trazado de línea para las predicciones.
	l, err := plotter.NewLine(ptsPred)
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	// Guardar la trama en un archivo PNG.
	p.Add(s, l)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "regression_line.png"); err != nil {
		log.Fatal(err)
	}
}

// predict utiliza el modelo de regresión entrenado para hacer una predicción.
func predict(tv float64) float64 {
	return 7.07 + tv*0.05
}
