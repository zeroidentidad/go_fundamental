package main

import "fmt"

type Primate interface {
	Alimentar(string)
}

type Antropoide struct{}

func (t Antropoide) Alimentar(fruta string) {
	fmt.Printf("Comiendo %s", fruta)
}

type Gorila struct {
	Antropoide
}

func DarDeComer(primate Primate) {
	primate.Alimentar("banana")
}

func main() {
	kong := Gorila{}
	DarDeComer(kong)
}
