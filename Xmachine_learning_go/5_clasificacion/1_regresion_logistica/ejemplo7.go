package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	// Prueba de prediccion calculada en el modelo

	// Abrir ejemplos de prueba
	f, err := os.Open("./newData/test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Crear un nuevo lector CSV que lea del archivo abierto.
	reader := csv.NewReader(f)

	// observed y predicted contendrán los valores observados y pronosticados
	// analizados del archivo de datos etiquetado.
	var observed []float64
	var predicted []float64

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

		// Leer en el valor observado.
		observedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		// Hacer la predicción correspondiente.
		score, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal := predict(score)

		// Agregar el registro al slice, si tiene el tipo esperado.
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	// Esta variable mantendrá recuento de valores verdaderos positivos y verdaderos negativos.
	var truePosNeg int

	// Acumular recuento verdadero positivo/negativo.
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

// predic hace una predicción basada en el
// modelo de regresión logística entrenado.
func predict(score float64) float64 {

	// Calcular la probabilidad pronosticada.
	p := 1 / (1 + math.Exp(-13.65*score+4.89))

	// Salida de la clase correspondiente.
	if p >= 0.5 {
		return 1.0
	}

	return 0.0
}
