package pkganimales

import "fmt"

type Gato struct {
	Nombre string
}

func (g Gato) Comunicarse() { // los type no pueden ser punteros para implementarse en interface
	fmt.Println("miauuu")
}
