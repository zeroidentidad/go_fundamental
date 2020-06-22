package filedirs

import (
	"errors"
	"io"
	"os"
)

// Operate manipula archivos y directorios
func Operate() error {
	// en una línea de comando esto creará un director /tmp/example,
	// también puedes usar una ruta absoluta en lugar de una relativa
	if err := os.Mkdir("example_dir", os.FileMode(0755)); err != nil {
		return err
	}

	// go to /tmp dir
	if err := os.Chdir("example_dir"); err != nil {
		return err
	}

	// f es un objeto de archivo genérico
	// también implementa múltiples interfaces
	// y se puede usar como lector o escritor
	// si se establecen los bits correctos al abrir
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}

	// escribimos un valor de longitud conocida en el archivo y validamos
	// que escribió correctamente
	value := []byte("hello\n")
	count, err := f.Write(value)
	if err != nil {
		return err
	}
	if count != len(value) {
		return errors.New("incorrect length returned from write")
	}

	if err := f.Close(); err != nil {
		return err
	}

	// read file
	f, err = os.Open("test.txt")
	if err != nil {
		return err
	}

	io.Copy(os.Stdout, f)

	if err := f.Close(); err != nil {
		return err
	}

	// go to /tmp dir
	if err := os.Chdir(".."); err != nil {
		return err
	}

	// limpieza, os.RemoveAll puede ser peligroso si
	// apunta al directorio incorrecto, usa la entrada del usuario,
	// y especialmente si corres como root
	if err := os.RemoveAll("example_dir"); err != nil {
		return err
	}

	return nil
}
