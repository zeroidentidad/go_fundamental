package main

import "fmt"

type Documento struct {
	Nombre string
}

func (d *Documento) GuardarEnArchivo() {
	fmt.Println("Documento Guardado en Archivo")
}

func (d *Documento) GuardarEnBaseDatos() {
	fmt.Println("Documento Guardado en Base de Datos")
}

func main() {
	doc := &Documento{"documento.doc"}

	doc.GuardarEnArchivo()
	doc.GuardarEnBaseDatos()
}
