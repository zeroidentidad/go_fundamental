package main

import "fmt"

type RespuestaJSON struct{}

func (r RespuestaJSON) Imprimir() {
	fmt.Println("Impresión de Respuesta JSON")
}

type RespuestaHTML struct{}

func (r RespuestaHTML) Imprimir() {
	fmt.Println("Impresión de Respuesta HTML")
}

type Emision struct{}

func (e Emision) EmitirJSON(r RespuestaJSON) {
	r.Imprimir()
}

func (e Emision) EmitirHTML(r RespuestaHTML) {
	r.Imprimir()
}

func main() {
	respJSON := RespuestaJSON{}
	respHTML := RespuestaHTML{}

	emision := Emision{}

	emision.EmitirJSON(respJSON)
	emision.EmitirHTML(respHTML)
}
