package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	// Abrir las observaciones y predicciones etiquetadas.
	f, err := os.Open("./data/labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear un nuevo lector CSV que lea del archivo abierto.
	reader := csv.NewReader(f)

	// observed y predicted contendrá los valores analizados y pronosticados analizados del archivo de datos etiquetado.
	var observed []int
	var predicted []int

	// La línea rastreará los números de fila para el registro.
	line := 1

	// Leer en los registros buscando tipos inesperados en las columnas.
	for {

		// Leer en una fila. Comprueba si estamos al final del archivo.
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
		observedVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal, err := strconv.Atoi(record[1])
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		// Agregar el registro al slice, si tiene el tipo esperado.
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	// classes contiene las tres clases posibles en los datos etiquetados.
	classes := []int{0, 1, 2}

	// Recorre cada clase.
	for _, class := range classes {

		// Estas variables mantendrán recuento de positivos verdaderos y recuento de falsos positivos.
		var truePos int
		var falsePos int
		var falseNeg int

		// Acumular los recuentos positivo verdadero y falso positivo.
		for idx, oVal := range observed {

			switch oVal {

			// Si el valor observado es la clase relevante, se deberia verificar para predecir esa clase.
			case class:
				if predicted[idx] == class {
					truePos++
					continue
				}

				falseNeg++

			// Si el valor observado es una clase diferente, deberíamos verificar si predecimos un falso positivo.
			default:
				if predicted[idx] == class {
					falsePos++
				}
			}
		}

		// Calcular la precision.
		precision := float64(truePos) / float64(truePos+falsePos)

		// Calcule el retiro.
		recall := float64(truePos) / float64(truePos+falseNeg)

		// Salida del valor de precisión a la salida estándar.
		fmt.Printf("\nPrecision (class %d) = %0.2f", class, precision)
		fmt.Printf("\nRecall (class %d) = %0.2f\n\n", class, recall)
	}
}
