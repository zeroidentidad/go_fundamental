package abstractfactory

import (
	"fmt"
	"testing"
)

// Test de uso del patrón
func TestPattern(t *testing.T) {
	fabricaPuertaMadera := &FabricaPuertaMadera{}
	puertaMadera := fabricaPuertaMadera.ConstruirPuerta()
	fmt.Printf("Se construyo puerta de: %s\n", puertaMadera.VerMaterial())

	fabricaPuertaMetal := &FabricaPuertaMetal{}
	puertaMetal := fabricaPuertaMetal.ConstruirPuerta()
	fmt.Printf("Se construyo puerta de: %s\n", puertaMetal.VerMaterial())
}
