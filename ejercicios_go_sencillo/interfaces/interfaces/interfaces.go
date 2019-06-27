package interfaces

import (
	"fmt"
	"io"
	"os"
)

// Copy copia los datos desde in a out primero directamente,
// luego usa un buffer. También escribe a stdout
func Copy(in io.ReadSeeker, out io.Writer) error {
	// escribimos a out, pero también Stdout
	w := io.MultiWriter(out, os.Stdout)

	// una copia estándar, esto puede ser peligroso si hay
	// muchos datos en in
	if _, err := io.Copy(w, in); err != nil {
		return err
	}

	in.Seek(0, 0)

	// escritura en buffer utilizando fragmentos de 64 bytes
	buf := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}

	// imprimimos una nueva línea
	fmt.Println()

	return nil
}
