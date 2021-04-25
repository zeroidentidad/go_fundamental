package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/gonum/stat"
)

func main() {

	// R-Cuadratica ó, coeficiente de determinación

	// Abrir los datos de las observaciones y predicciones continuas.
	f, err := os.Open("./data/continuous_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear un nuevo lector CSV que lea del archivo abierto.
	reader := csv.NewReader(f)

	// observed y predicted contendrá los valores analizados y pronosticados analizados del archivo de datos continuo.
	var observed []float64
	var predicted []float64

	// La línea rastreará los números de fila para el registro.
	line := 1

	// Leer en los registros buscando tipos inesperados en las columnas.
	for {

		// Leer en una fila. Comprobar si estamos al final del archivo.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Saltar el encabezado
		if line == 1 {
			line++
			continue
		}

		// Leer en los valores observados y pronosticados.
		observedVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		// Agregar el registro al slice, si tiene el tipo esperado.
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	// Calcular el valor R^2.
	rSquared := stat.RSquaredFrom(observed, predicted, nil)

	// Salida del valor R^2 a la salida estándar.
	fmt.Printf("\nR^2 = %0.2f\n\n", rSquared)
}
