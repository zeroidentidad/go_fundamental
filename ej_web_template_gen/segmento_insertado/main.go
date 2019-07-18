package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, 5*10)
	if err != nil {
		log.Fatalln(err)
	}
}

/*
En ingeniería de software, un pipeline consiste en una cadena de elementos de procesamiento.
(procesos, hilos, goroutines, funciones, etc.), dispuestos de forma que
la salida de cada elemento es la entrada del siguiente;
*/
