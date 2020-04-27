package iterator

import (
	"fmt"
	"testing"
)

func TestPattern(t *testing.T) {
	radio := &Radio{}
	radio.Registrar("FM100")
	radio.Registrar("FM200")
	radio.Registrar("FM300")

	iterador := radio.CrearIterador()

	fmt.Printf("Escuhando la radio %s\n", iterador.Valor())

	iterador.Siguiente()
	fmt.Printf("Escuhando la radio %s\n", iterador.Valor())

	iterador.Siguiente()
	fmt.Printf("Escuhando la radio %s\n", iterador.Valor())

	iterador.Anterior()
	fmt.Printf("Escuhando nuevamente la radio %s\n", iterador.Valor())
}
