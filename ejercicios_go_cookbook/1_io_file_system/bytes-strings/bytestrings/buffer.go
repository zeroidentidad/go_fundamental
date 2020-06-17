package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Buffer muestra algunos trucos para inicializar bytes Buffers
// Estos buffers implementan una interfaz io.Reader
func Buffer(rawString string) *bytes.Buffer {

	// con una cadena codificada en bytes sin procesar
	rawBytes := []byte(rawString)

	// hay varias formas de crear un búfer a partir de
	// bytes sin procesar o de la cadena original
	var b = new(bytes.Buffer)
	b.Write(rawBytes)

	// alternativamente
	b = bytes.NewBuffer(rawBytes)

	// y evitando la matriz de bytes inicial por completo
	b = bytes.NewBufferString(rawString)

	return b
}

// ToString es un ejemplo de tomar un io.Reader y consumir
// todo, luego devolviendo una cadena
func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
