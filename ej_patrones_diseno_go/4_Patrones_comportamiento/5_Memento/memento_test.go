package memento

import (
	"fmt"
	"testing"
)

// Test de uso del patrón
func TestPattern(t *testing.T) {
	editor := &Editor{}
	editor.Escribir("TextoA")
	editor.Escribir("TextoB")
	fmt.Printf("El editor contiene:%s\n", editor.VerContenido())

	fmt.Println("Se guarda el estado actual")
	memento := editor.Guardar()

	fmt.Println("Se escribe unnuevo contenido")
	editor.Escribir("TextoC")

	fmt.Printf("El editor ahora contiene:%s\n", editor.VerContenido())

	fmt.Println("Se restaura el contenido guardado")
	editor.Restaurar(memento)

	fmt.Printf("El editor nuevamente contiene:%s\n", editor.VerContenido())
}
