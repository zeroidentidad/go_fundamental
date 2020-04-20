package main

import (
	"fmt"

	model "./model"
)

type GuardadorDocumentoArchivo struct {
	Path string
}

type GuardadorDocumentoBaseDatos struct {
	StringConnection string
}

type GuardadorDocumento interface {
	Guardar(d model.Documento)
}

func (gda GuardadorDocumentoArchivo) Guardar(d model.Documento) {
	fmt.Println("Documento Guardado en Archivo")
}

func (gdbd GuardadorDocumentoBaseDatos) Guardar(d model.Documento) {
	fmt.Println("Documento Guardado en Base de Datos")
}

func main() {
	doc := model.Documento{Nombre: "documento.doc"}

	gda := GuardadorDocumentoArchivo{"/var/tmp/"}
	gda.Guardar(doc)

	gdbd := GuardadorDocumentoBaseDatos{"root@123"}
	gdbd.Guardar(doc)
}
