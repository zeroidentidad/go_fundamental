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

	// Abrir las observaciones binarias y las predicciones.
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

		// Saltar el encabezado.
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

	// Esta variable mantendrá recuento de valores verdaderos positivos y verdaderos negativos.
	var truePosNeg int

	// Acumular el recuento de verdadero positivo/negativo.
	for idx, oVal := range observed {
		if oVal == predicted[idx] {
			truePosNeg++
		}
	}

	// Calcular la precisión (precisión del subconjunto).
	accuracy := float64(truePosNeg) / float64(len(observed))

	// Salida del valor de precisión a la salida estándar.
	fmt.Printf("\nAccuracy = %0.2f\n\n", accuracy)
}
