package main

import "fmt"

type InterfaceMetodosAbstractos interface {
	MetodoAbstracto()
}

type ClaseAbstracta struct{}

func (ca *ClaseAbstracta) MetodoConcreto(self InterfaceMetodosAbstractos) {
	self.MetodoAbstracto()
}

type ClaseHija struct {
	*ClaseAbstracta
}

func (ch *ClaseHija) MetodoAbstracto() {
	fmt.Println("Soy metodo abstracto")
}

func main() {
	claseHija := &ClaseHija{&ClaseAbstracta{}}
	claseHija.MetodoConcreto(claseHija)
}
