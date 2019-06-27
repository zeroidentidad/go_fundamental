package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

// WorkWithBuffer hará uso del buffer creado por la función Buffer
// Buffer function
func WorkWithBuffer() error {
	rawString := "Es sencillo codificar unicode en un array byte"

	b := Buffer(rawString)

	// Podemos rápidamente convertir un buffer de nuevo en bytes con
	// b.Bytes() o en un string con b.String()
	fmt.Println(b.String())

	// como esto es un io Reader podemos hacer uso de
	// funciones genéricas como
	s, err := toString(b)
	if err != nil {
		return err
	}
	fmt.Println(s)

	// podemos también tomar nuestros bytes y crear un bytes reader
	// estos readers implementan las interfaces io.Reader, io.ReaderAt,
	// io.WriterTo, io.Seeker, io.ByteScanner, e
	// io.RuneScanner
	reader := bytes.NewReader([]byte(rawString))

	// Podemos también conectarlo a un escaner que nos permite
	// buffered lectura y tokeización
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	// nos iteramos sobre todos los eventos del escaneo
	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}

	return nil
}



