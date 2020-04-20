package main

import "fmt"

type Boton interface {
	OnClick()
	OnDobleClick()
}

type BotonLogin struct{}

func (b BotonLogin) OnClick() {
	fmt.Println("Evento onClick")
}

func (b BotonLogin) OnDobleClick() {
	fmt.Println("No requiero de este evento pero debo implementarlo")
}

func main() {
	boton := BotonLogin{}

	boton.OnClick()
	boton.OnDobleClick()
}
