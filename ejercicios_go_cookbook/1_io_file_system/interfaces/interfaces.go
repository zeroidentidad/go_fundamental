package interfaces

import (
	"fmt"
	"io"
	"os"
)

// Copy copia datos de adentro hacia afuera primero directamente,
// luego usando un buffer. También escribe a stdout
func Copy(in io.ReadSeeker, out io.Writer) error {

	// write out, Stdout
	w := io.MultiWriter(out, os.Stdout)

	// una copia estándar, esto puede ser peligroso
	// si hay muchos datos en in
	if _, err := io.Copy(w, in); err != nil {
		return err
	}
	in.Seek(0, 0)

	// escritura en búfer utilizando fragmentos de 64 bytes
	buf := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}

	// imprimir una nueva línea
	fmt.Println()
	return nil
}
