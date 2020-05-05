package prototype

import (
	"fmt"
	"testing"
)

// Test de uso del patrón
func TestPattern(t *testing.T) {
	elementoA := &Elemento{"Azufre", 1}

	elementoB := elementoA.Clonar()
	elementoB.Material = elementoB.Material + " (fortificado)"

	elementoC := elementoB.Clonar()

	fmt.Printf("El elemento A es de %s y se clono %d veces\n", elementoA.Material, elementoA.Copias)
	fmt.Printf("El elemento B es de %s y se clono %d veces\n", elementoB.Material, elementoB.Copias)
	fmt.Printf("El elemento C es de %s y se clono %d veces\n", elementoC.Material, elementoC.Copias)
}
