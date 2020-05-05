package factorymethod

import (
	"fmt"
	"testing"
)

// Test de uso del patrón
func TestPattern(t *testing.T) {
	fmt.Println("El entrevisatador de IT pregunta:")
	recursosHumanosIT := &RecusosHumanosIT{&RecursosHumanos{}}
	recursosHumanosIT.TomarEntrevista(recursosHumanosIT)

	fmt.Println("\nEl entrevisatador de Finanzas pregunta:")
	recursosHumanosFinanzas := &RecusosHumanosFinanzas{&RecursosHumanos{}}
	recursosHumanosFinanzas.TomarEntrevista(recursosHumanosFinanzas)
}
