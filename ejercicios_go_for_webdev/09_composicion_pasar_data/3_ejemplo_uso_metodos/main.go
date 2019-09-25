package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Nombre string
	Edad   int
}

func (p person) SomeProcessing() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Edad * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	p := person{
		Nombre: "Jesus Ferrer",
		Edad:   100,
	}

	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}
}

// Hablando en general, de prácticas:
// funciones de llamada en plantillas para formatear solamente; no procesamiento o acceso a datos.

// Las principales razones porque no se quiere hacer procesamiento de datos en plantilla:
// (1) separación de referencia
// (2) si estás usando una función más de una vez en una plantilla,
// el servidor necesita realizar el procesamiento más de una vez.
// (aunque la biblioteca estándar podría procesar en caché, indagar)
