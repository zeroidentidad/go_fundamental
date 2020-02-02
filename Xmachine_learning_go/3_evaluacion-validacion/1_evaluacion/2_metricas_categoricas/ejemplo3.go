package main

import (
	"fmt"

	"gonum.org/v1/gonum/integrate"
	"gonum.org/v1/gonum/stat"
)

func main() {

	// AUC: Area bajo la curva

	// Definir puntajes y clases.
	scores := []float64{0.1, 0.35, 0.4, 0.8}
	classes := []bool{true, false, true, false}

	// Calcular las tasas positivas verdaderas (retiros) y las tasas positivas falsas.
	tpr, fpr, _ := stat.ROC(nil, scores, classes, nil)

	// Calcule el área bajo la curva.
	auc := integrate.Trapezoidal(fpr, tpr)

	// Salida de los resultados a la salida estándar.
	fmt.Printf("true  positive rate: %v\n", tpr)
	fmt.Printf("false positive rate: %v\n", fpr)
	fmt.Printf("auc: %v\n", auc)
}
