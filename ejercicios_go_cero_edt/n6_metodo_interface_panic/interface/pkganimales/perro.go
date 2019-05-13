package pkganimales

import "fmt"

type Perro struct {
	Nombre string
}

func (p Perro) Comunicarse() { // los type no pueden ser punteross
	fmt.Println("wofff")
}
