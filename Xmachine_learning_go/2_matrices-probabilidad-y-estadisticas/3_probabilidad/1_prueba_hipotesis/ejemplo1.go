package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func main() {

	// Definir las frecuencias observadas.
	observed := []float64{
		260.0, // Este número es el número de observados sin ejercicio regular.
		135.0, // Este número es el número de observados con ejercicio esporádico.
		105.0, // Este número es el número de observados con ejercicio regular.
	}

	// Definir total observado
	totalObserved := 500.0

	// Calcular la frecuencia esperada (nuevamente asumiendo la hypotesis nula)
	expected := []float64{
		totalObserved * 0.60,
		totalObserved * 0.25,
		totalObserved * 0.15,
	}

	// Calcular la estadística de prueba de Chi cuadrado.
	chiSquare := stat.ChiSquare(observed, expected)

	// Salida de la estadística de prueba a la salida estándar.
	fmt.Printf("\nChi-cuadrado: %0.2f\n", chiSquare)

	// Crear una distribución Chi-cuadrado con K grados de libertad.
	// En este caso tenemos K = 3-1 = 2, porque los grados de libertad
	// para una distribución de Chi-cuadrado es el número de categorías posibles menos uno.
	chiDist := distuv.ChiSquared{
		K:   2.0,
		Src: nil,
	}

	// Calcular el valor p para nuestra estadística de prueba específica.
	pValue := chiDist.Prob(chiSquare)

	// Salida del valor p a la salida estándar.
	fmt.Printf("p-value: %0.4f\n\n", pValue)
}
