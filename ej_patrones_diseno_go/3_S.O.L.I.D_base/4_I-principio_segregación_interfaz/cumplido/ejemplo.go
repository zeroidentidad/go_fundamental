package main

import "fmt"

type OnClickListener interface {
	OnClick()
}

type OnDobleClickListener interface {
	OnDobleClick()
}

type BotonLogin struct{}

func (b BotonLogin) OnClick() {
	fmt.Println("Evento onClick en login")
}

type BotonIcono struct{}

func (b BotonIcono) OnDobleClick() {
	fmt.Println("Evento onDobleClick en ícono")
}

type BotonComplejo struct{}

func (b BotonComplejo) OnClick() {
	fmt.Println("Evento onClick en boton complejo")
}

func (b BotonComplejo) OnDobleClick() {
	fmt.Println("Evento onDobleClick en boton complejo")
}

func main() {
	botonLogin := BotonLogin{}
	botonLogin.OnClick()

	botonIcono := BotonIcono{}
	botonIcono.OnDobleClick()

	botonComplejo := BotonComplejo{}
	botonComplejo.OnClick()
	botonComplejo.OnDobleClick()
}
