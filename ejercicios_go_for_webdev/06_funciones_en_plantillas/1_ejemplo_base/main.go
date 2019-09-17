package main

import (
	"log"
	"os"
	"strings"
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

// crea un FuncMap para registrar funciones.
// "uc" es como se llamará el func en la plantilla
// "uc" es la función ToUpper del paquete strings
// "ft" es una función que se ha declarado
// "ft" corta la cadena, devolviendo los primeros tres caracteres =>

var fm = template.FuncMap{ // <-
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
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
		Lema:   "El odio nunca termina con odio, pero solo con el amor se cura.",
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

	data := struct {
		Sabiduria  []sabio
		Transporte []carro
	}{
		sabios,
		carros,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
