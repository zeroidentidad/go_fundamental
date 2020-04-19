package main

import "fmt"

type Usuario struct {
	nombre   string
	apellido string
}

func (u Usuario) getDatosPersonales() string {
	return fmt.Sprintf("%s, %s", u.nombre, u.apellido)
}

type Administrador struct {
	*Usuario
	sector string
}

func (a Administrador) getDatosCompletos() string {
	return fmt.Sprintf("%s - %s", a.getDatosPersonales(), a.sector)
}

func main() {
	var administrador = Administrador{&Usuario{"Jose", "Luis"}, "Computos"}

	fmt.Println(administrador.getDatosPersonales())
	fmt.Println(administrador.getDatosCompletos())
}
