package main

import "fmt"

type Felino interface {
	Caminar()
	Saltar()
}

type Leon struct {
	Nombre string
}

func (t Leon) Caminar() {
	fmt.Println("CAMINAR")
}

func (t Leon) Saltar() {
	fmt.Println("SALTAR")
}

func SaltarYCaminar(felino Felino) {
	felino.Saltar()
	felino.Caminar()
}

func main() {
	leon := Leon{"Simba"}
	SaltarYCaminar(leon)
}
