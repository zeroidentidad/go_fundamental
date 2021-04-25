package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	// Ajuste y evaluación de modelo AR(2) parte 2 comparando con data original

	// Abrir archivo conjunto de datos diferenciados de registro (Log).
	transFile, err := os.Open("log_diff_series.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer transFile.Close()

	// Crear lector CSV que lea el archivo abierto.
	transReader := csv.NewReader(transFile)

	// Leer en todos los registros CSV
	transReader.FieldsPerRecord = 2
	transData, err := transReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Recorrer los datos que predicen las observaciones transformadas.
	var transPredictions []float64
	for i, _ := range transData {

		// Omitir encabezado y las dos primeras observaciones
		// (porque se necesitan dos retrasos para hacer una predicción).
		if i == 0 || i == 1 || i == 2 {
			continue
		}

		// Analizar primer retraso.
		lagOne, err := strconv.ParseFloat(transData[i-1][1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Analizar segundo retraso.
		lagTwo, err := strconv.ParseFloat(transData[i-2][1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Predecir variable transformada con modelo AR entrenado.
		transPredictions = append(transPredictions, 0.008159+0.234953*lagOne-0.173682*lagTwo)
	}

	// Abrir archivo original de conjunto de datos.
	origFile, err := os.Open("./data/AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer origFile.Close()

	// Crear lector CSV que lea desde archivo abierto.
	origReader := csv.NewReader(origFile)

	// Leer en todos los registros CSV
	origReader.FieldsPerRecord = 2
	origData, err := origReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// pts* mantendrá los valores para trazar.
	ptsObs := make(plotter.XYs, len(transPredictions))
	ptsPred := make(plotter.XYs, len(transPredictions))

	// Inviertir la transformación y calcular el MAE.
	var mAE float64
	var cumSum float64
	for i := 4; i <= len(origData)-1; i++ {

		// Analizar la observación original.
		observed, err := strconv.ParseFloat(origData[i][1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Analizar fecha original.
		date, err := strconv.ParseFloat(origData[i][0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Obtener suma acumulativa hasta el índice en las predicciones transformadas.
		cumSum += transPredictions[i-4]

		// Calcular la predicción transformada inversa.
		predicted := math.Exp(math.Log(observed) + cumSum)

		// Acumular el MAE.
		mAE += math.Abs(observed-predicted) / float64(len(transPredictions))

		// Llenar puntos para trazar.
		ptsObs[i-4].X = date
		ptsPred[i-4].X = date
		ptsObs[i-4].Y = observed
		ptsPred[i-4].Y = predicted
	}

	// Salida del MAE a la salida estándar.
	fmt.Printf("\nMAE = %0.2f\n\n", mAE)

	// Crear la trama.
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "time"
	p.Y.Label.Text = "passengers"
	p.Add(plotter.NewGrid())

	// Agregar puntos de trazado de línea para la serie de tiempo.
	lObs, err := plotter.NewLine(ptsObs)
	if err != nil {
		log.Fatal(err)
	}
	lObs.LineStyle.Width = vg.Points(1)

	lPred, err := plotter.NewLine(ptsPred)
	if err != nil {
		log.Fatal(err)
	}
	lPred.LineStyle.Width = vg.Points(1)
	lPred.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}

	// Guardar trama en archivo PNG.
	p.Add(lObs, lPred)
	p.Legend.Add("Observed", lObs)
	p.Legend.Add("Predicted", lPred)
	if err := p.Save(10*vg.Inch, 4*vg.Inch, "passengers_ts.png"); err != nil {
		log.Fatal(err)
	}
}
