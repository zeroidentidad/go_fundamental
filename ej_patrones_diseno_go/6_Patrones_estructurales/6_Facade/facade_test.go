package facade

import (
	"fmt"
	"testing"
)

// Test de uso del patrón
func TestPattern(t *testing.T) {
	agencia := &AgenciaViaje{}

	fmt.Printf("%s\n", agencia.BuscarPaquete("2018-12-01", "2018-12-08"))
}
