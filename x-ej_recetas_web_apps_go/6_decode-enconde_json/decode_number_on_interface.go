package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	_json := `10` // Este number JSON es un número entero.

	var n interface{}

	err := json.NewDecoder(strings.NewReader(_json)).Decode(&n)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("type: %T; value: %v\n", n, n)
}
