package interfaces

import (
	"io"
	"os"
)

// PipeExample nos ayuda con varios ejemplos más de usos de
//interfaces io
func PipeExample() error {
	// el pipe reader y pipe writer implementan
	// io.Reader e io.Writer
	r, w := io.Pipe()

	// esto tiene que ser ejecutado en una rutina go separada
	// ya que se bloqueará esperando al reader
	// cerrar al final para limpieza
	go func() {
		// por ahora escribiremos algo básico,
		// esto también puede ser usado para codificar json
		// base64 encode, etc.
		w.Write([]byte("testn"))
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}
