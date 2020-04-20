package main

import "fmt"

type Respuesta interface {
	Imprimir()
}

type RespuestaJSON struct{}

func (r RespuestaJSON) Imprimir() {
	fmt.Println("Impresión de Respuesta JSON")
}

type RespuestaHTML struct{}

func (r RespuestaHTML) Imprimir() {
	fmt.Println("Impresión de Respuesta HTML")
}

type Emision struct{}

func (e Emision) Emitir(r Respuesta) {
	r.Imprimir()
}

func main() {
	respJSON := RespuestaJSON{}
	respHTML := RespuestaHTML{}

	emision := Emision{}

	emision.Emitir(respJSON)
	emision.Emitir(respHTML)
}
