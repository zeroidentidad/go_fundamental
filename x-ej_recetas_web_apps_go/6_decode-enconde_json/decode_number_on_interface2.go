package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	_json := `10`

	var n interface{}

	dec := json.NewDecoder(strings.NewReader(_json))
	dec.UseNumber() // Llamar método UseNumber() en decoder antes de usarlo.
	err := dec.Decode(&n)
	if err != nil {
		log.Fatal(err)
	}

	// Aserción del tipo valor interface{} a un json.Number,
	// y luego llamar método Int64() para obtener el número como int64 en Go.
	nInt64, err := n.(json.Number).Int64()
	if err != nil {
		log.Fatal(err)
	}

	// Igual, puede usarse método String() para obtener número como string en Go.
	nString := n.(json.Number).String()

	fmt.Printf("type: %T; value: %v\n", n, n)
	fmt.Printf("type: %T; value: %v\n", nInt64, nInt64)
	fmt.Printf("type: %T; value: %v\n", nString, nString)
}
