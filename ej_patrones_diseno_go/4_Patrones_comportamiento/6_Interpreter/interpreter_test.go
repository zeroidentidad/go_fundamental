package interpreter

import (
	"fmt"
	"testing"
)

// Test de uso del patrón
func TestPattern(t *testing.T) {
	expresionA := &ExpresionTerminal{"Perro"}
	expresionB := &ExpresionTerminal{"Gato"}
	expresionC := &ExpresionTerminal{"Perro"}

	contextoOR := Contexto{"Perro"}
	expresionOR := &ExpresionOR{expresionA, expresionB}
	fmt.Printf("La expresion OR contiene la palabra perro: %v\n", expresionOR.Interpretar(contextoOR))

	contextoAND := Contexto{"Perro"}
	expresionAND := &ExpresionAND{expresionA, expresionC}
	fmt.Printf("La expresion AND contiene dos palabras perro: %v\n", expresionAND.Interpretar(contextoAND))
}
