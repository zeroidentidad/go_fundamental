package observer

import (
	"fmt"
	"testing"
)

func TestPattern(t *testing.T) {
	observadorA := &ObservadorEmpleo{"Juan"}
	observadorB := &ObservadorEmpleo{"Maria"}

	agencia := &AgenciaEmpleo{}
	agencia.Adquirir(observadorA)
	agencia.Adquirir(observadorB)

	agencia.IngresarOfertaLaboral("Programador JAVA")
	fmt.Printf("\n")
	agencia.IngresarOfertaLaboral("Programador C#")
}
