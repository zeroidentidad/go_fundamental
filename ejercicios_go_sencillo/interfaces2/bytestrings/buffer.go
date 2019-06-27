package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Buffer nos muestra algunos trucos para inicializar búferes
// de bytes
// Estos búferes implementan una interface io.Reader
func Buffer(rawString string) *bytes.Buffer {

	// Empezaremos con un string codificado en bytes sin procesar
	rawBytes := []byte(rawString)

	// existen varios modos para crear un buffer a partir de
	// los bytes sin procesar o desde el string original
	var b = new(bytes.Buffer)
	b.Write(rawBytes)

	// alternativamente
	b = bytes.NewBuffer(rawBytes)

	// y evitando por completo el conjunto de bytes inicial
	b = bytes.NewBufferString(rawString)

	return b
}

// ToString es un ejemplo de tomar un io.Reader y consumirlo,
// devolviendo luego un string
func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
