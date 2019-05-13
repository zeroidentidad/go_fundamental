package pkganimales

import "fmt"

type Gato struct {
	Nombre string
}

func (g Gato) Comunicarse() {
	fmt.Println("miauuu")
}
