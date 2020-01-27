package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64" //deprecado sin uso
)

func main() {

	// Crear un nuevo valor vectorial.
	myvector := mat64.NewVector(2, []float64{11.0, 5.2})

	fmt.Println(myvector)
}
