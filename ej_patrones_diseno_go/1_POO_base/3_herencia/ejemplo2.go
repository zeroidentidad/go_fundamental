package main

import "fmt"

type Empleado struct {
	legajo  string
	ingreso string
}

func (e *Empleado) InformarIngreso() {
	fmt.Printf("Mi legajo fue el %s.\n", e.ingreso)
}

func (e *Empleado) Presentarse() {
	fmt.Printf("Mi legajo es %s.\n", e.legajo)
}

type Gerente struct {
	Empleado
}

func (g *Gerente) Presentarse() {
	g.Empleado.Presentarse()

	fmt.Println("Mi cargo es el de Gerente.\n")
}

func main() {
	gerente := Gerente{
		Empleado{"A0001", "01-01-2020"},
	}

	gerente.Presentarse()
	gerente.InformarIngreso()
}
