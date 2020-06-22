package tempfiles

import (
	"fmt"
	"io/ioutil"
	"os"
)

// WorkWithTemp dará algunos patrones básicos para trabajar
// con archivos y directorios temporales
func WorkWithTemp() error {
	// Si necesita un lugar temporal para almacenar archivos con el
	// mismo nombre, es decir. template1-10.html un directorio temporal es un buen
	// forma de abordarlo, el primer argumento en blanco significa que
	// usará crear el directorio en la ubicación devuelta por
	// os.TempDir ()
	t, err := ioutil.TempDir("", "tmp")
	if err != nil {
		return err
	}

	// Esto eliminará todo el contenido del archivo temporal cuando esto
	// la función se cierra si desea hacer esto más tarde, asegúrese de regresar
	// el nombre del directorio a la función de llamada
	defer os.RemoveAll(t)

	// el directorio debe existir para crear el archivo temporal
	// creado. t es un objeto * os.File.
	tf, err := ioutil.TempFile(t, "tmp")
	if err != nil {
		return err
	}

	fmt.Println(tf.Name())

	// normalmente eliminaríamos el archivo temporal aquí, pero porque
	// lo estamos colocando en un directorio temporal, se limpia
	// por el aplazamiento anterior

	return nil
}
