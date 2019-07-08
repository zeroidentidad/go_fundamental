// Package stringutil contiene funciones de utilidad para trabajar con cadenas.
package stringutil

// Reverso devuelve su cadena de argumento invertida en sentido de runas de izquierda a derecha.
func Reverso(s string) string {
	return reverseDos(s)
}

/*
go build
	Va construir reverso.go reverseDos.go
	no producirá un archivo de salida. Tomar encuenta desde el main

go install
 	colocará el paquete dentro del directorio pkg del espacio de trabajo. (gopath)
*/
