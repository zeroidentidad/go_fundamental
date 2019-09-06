package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sabio struct {
	Nombre string
	Lema   string
}

type carro struct {
	Fabricante string
	Modelo     string
	Puertas    int
}

type items struct {
	Sabiduria  []sabio
	Transporte []carro
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	b := sabio{
		Nombre: "Buddha",
		Lema:   "La creencia de no creer",
	}

	g := sabio{
		Nombre: "Gandhi",
		Lema:   "Sé el cambio",
	}

	m := sabio{
		Nombre: "Martin Luther King",
		Lema:   "El odio nunca cesa con el odio, pero solo con el amor se cura.",
	}

	f := carro{
		Fabricante: "Ford",
		Modelo:     "F150",
		Puertas:    2,
	}

	c := carro{
		Fabricante: "Toyota",
		Modelo:     "Corolla",
		Puertas:    4,
	}

	sabios := []sabio{b, g, m}
	carros := []carro{f, c}

	data := items{
		Sabiduria:  sabios,
		Transporte: carros,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
