package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func doble(x int) int {
	return x + x
}

func cuadrado(x int) float64 {
	return math.Pow(float64(x), 2)
}

func raizCuadrada(x float64) float64 {
	return math.Sqrt(x)
}

var fm = template.FuncMap{
	"fdbl":  doble,
	"fsq":   cuadrado,
	"fsqrt": raizCuadrada,
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
