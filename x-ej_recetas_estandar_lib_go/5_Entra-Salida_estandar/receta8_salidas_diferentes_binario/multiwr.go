package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {

	buf := bytes.NewBuffer([]byte{})
	// archivos de salida en texto plano
	f, err := os.OpenFile("sample.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	f2, err := os.OpenFile("sample.json", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// escritura en buffer y archivos de ejemplo
	wr := io.MultiWriter(buf, f, f2)
	_, err = io.WriteString(wr, "Hello, Golang is working!\n")
	if err != nil {
		panic(err)
	}

	// ejemplo de salida en buffer de ejecución
	fmt.Println("Contenido de buffer: " + buf.String())
}
