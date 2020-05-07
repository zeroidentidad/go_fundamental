package adapter

import (
	"fmt"
	"testing"
)

func TestPattern(t *testing.T) {
	jugadorA := &Jugador{&Elfo{}}
	fmt.Printf("Jugador A: %s\n", jugadorA.Atacar())

	jugadorB := &Jugador{&MagoAdaptador{&Mago{}}}
	fmt.Printf("Jugador B: %s\n", jugadorB.Atacar())
}
