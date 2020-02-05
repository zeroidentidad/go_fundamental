package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

func main() {

	// Abrir archivo dataset de entrenamiento
	f, err := os.Open("./data/training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Cree un nuevo lector CSV que lea del archivo abierto.
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 4

	// Leer en todos los registros CSV
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// featureData contendrá los valores flotantes que se usarán para formar la matriz de características.
	featureData := make([]float64, 4*len(rawCSVData))
	yData := make([]float64, len(rawCSVData))

	// featureIndex e yIndex rastrearán el índice actual de los valores de la matriz.
	var featureIndex int
	var yIndex int

	// Mover secuencialmente las filas en un slice de flotantes.
	for idx, record := range rawCSVData {

		// Omitir la fila de encabezado.
		if idx == 0 {
			continue
		}

		// Bucle sobre las columnas flotantes.
		for i, val := range record {

			// Convertir el valor en un flotante.
			valParsed, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal("Could not parse float value")
			}

			if i < 3 {

				// Agregar una intersección al modelo.
				if i == 0 {
					featureData[featureIndex] = 1
					featureIndex++
				}

				// Agregar el valor flotante al slice de caracteristicas en flotantes.
				featureData[featureIndex] = valParsed
				featureIndex++
			}

			if i == 3 {

				// Agregar el valor flotante al slices de flotantes y.
				yData[yIndex] = valParsed
				yIndex++
			}

		}
	}

	// Formar las matrices que serán entrada a la regresión.
	features := mat.NewDense(len(rawCSVData), 4, featureData)
	y := mat.NewVecDense(len(rawCSVData), yData)

	if features != nil && y != nil {
		fmt.Println("Matrices formadas para ridge regression")
	}
}
