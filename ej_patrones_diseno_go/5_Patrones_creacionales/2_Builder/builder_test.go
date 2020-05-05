package builder

import (
	"fmt"
	"testing"
)

// Test de uso del patrón
func TestPattern(t *testing.T) {
	localComida := &LocalComida{}

	hamburguesaA := localComida.ConstruirHamburguesa(&ConstructorHamburguesaSimple{})
	hamburguesaB := localComida.ConstruirHamburguesa(&ConstructorHamburguesaDoble{})

	fmt.Printf("Se solicito una hamburguesa: %s\n", hamburguesaA.Tamanio)
	fmt.Printf("Se solicito una hamburguesa: %s\n", hamburguesaB.Tamanio)
}
