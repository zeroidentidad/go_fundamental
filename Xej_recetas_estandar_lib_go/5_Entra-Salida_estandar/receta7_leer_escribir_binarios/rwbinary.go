package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {

	// Escribir valores binarios
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, 1.004); err != nil {
		panic(err)
	}

	vstring := []byte("Hola mundo X")
	if err := binary.Write(buf, binary.BigEndian, vstring); err != nil {
		panic(err)
	}

	// Leer los valores escritos
	var num float64
	if err := binary.Read(buf, binary.BigEndian, &num); err != nil {
		panic(err)
	}
	fmt.Printf("float64: %.3f\n", num)

	greeting := make([]byte, len(vstring))
	if err := binary.Read(buf, binary.BigEndian, &greeting); err != nil {
		panic(err)
	}
	fmt.Printf("string: %s\n", string(greeting))
}
