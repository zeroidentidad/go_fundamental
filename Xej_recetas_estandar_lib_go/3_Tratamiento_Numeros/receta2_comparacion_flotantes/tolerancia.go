package main

import (
	"fmt"
	"math"
)

const da = 0.49999999999999998889776975374843459576368331909180
const db = 0.5

func main() {

	daStr := fmt.Sprintf("%.10f", da)
	dbStr := fmt.Sprintf("%.10f", db)

	// Mientras formatea el número a cadena
	// se redondea a 5.
	fmt.Printf("Strings %s = %s iguales: %v \n", daStr, dbStr, dbStr == daStr)

	// Los números no son iguales
	fmt.Printf("Números iguales: %v \n", db == da)

	// Como la precisión de la representación flotante
	// está limitada. Para la comparación de flotante es
	// mejor usar la comparación con cierta tolerancia.
	fmt.Printf("Números iguales con TOLERANCIA: %v \n", Equals(da, db))

}

const TOLERANCIA = 1e-8

// Equals compara los números de coma flotante
// con tolerancia 1e-8
func Equals(numA, numB float64) bool {
	delta := math.Abs(numA - numB)
	if delta < TOLERANCIA {
		return true
	}
	return false
}
