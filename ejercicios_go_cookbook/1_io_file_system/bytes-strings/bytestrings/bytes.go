package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

// WorkWithBuffer utilizará el búfer creado por la
// función Buffer
func WorkWithBuffer() error {
	rawString := "es fácil codificar unicode en una matriz de bytes ❤️"

	b := Buffer(rawString)

	// podemos convertir rápidamente un búfer en bytes con
	// b.Bytes() o una cadena con b.String()
	fmt.Println(b.String())

	// porque este es un io Reader podemos hacer uso de genéricos
	// funciones del lector io tales como
	s, err := toString(b)
	if err != nil {
		return err
	}
	fmt.Println(s)

	// también podemos tomar nuestros bytes y crear un lector de bytes
	// estos lectores implementan interfaces io.Reader, io.ReaderAt,
	// io.WriterTo, io.Seeker, io.ByteScanner y io.RuneScanner
	reader := bytes.NewReader([]byte(rawString))

	// también podemos conectarlo a un escáner que permita
	// la lectura / tokenización almacenada en búfer
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	// iterar sobre todos los eventos de escaneo
	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}

	return nil
}
