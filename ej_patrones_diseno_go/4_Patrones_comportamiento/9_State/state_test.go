package state

import (
	"fmt"
	"testing"
)

func TestPattern(t *testing.T) {
	editor := &EditorMarkdown{}
	fmt.Printf("Texto redactado sin estado: %s\n", editor.Redactar("Lorem ipsum"))

	editor.SetEstado(&EstadoNegrita{})
	fmt.Printf("Texto redactado en negrita: %s\n", editor.Redactar("Lorem ipsum"))

	editor.SetEstado(&EstadoCursiva{})
	fmt.Printf("Texto redactado en cursiva: %s\n", editor.Redactar("Lorem ipsum"))
}
