package interfaces

import (
	"io"
	"os"
)

// PipeExample ayuda a dar algunos ejemplos más del uso
// de interfaces io

func PipeExample() error {
	// pipe reader y pipe writer implementan
	// io.Reader y io.Writer
	r, w := io.Pipe()

	// esto debe ejecutarse en una rutina separada
	// ya que bloqueará la espera del lector
	// cierre al final para la limpieza
	go func() {

		// por ahora algo básico,
		// esto también podría usarse para codificar json
		// codificación base64, etc.
		w.Write([]byte("testn"))
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}

	return nil
}
