package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	// Digamos que lo único que nos interesa es procesar
	// el array "genres" del siguiente objeto JSON
	_json := `{"title": "Top Gun", "genres": ["action", "romance"], "year": 1986}`

	// Decodificar el objeto JSON a un tipo map[string]json.RawMessage
	// los valores json.RawMessage en el mapa conservarán sus valores
	// JSON originales sin decodificar.
	var m map[string]json.RawMessage

	err := json.NewDecoder(strings.NewReader(_json)).Decode(&m)
	if err != nil {
		log.Fatal(err)
	}

	// Luego se puede acceder al valor JSON "genres" desde el map y decodificarlo
	// normalmente usando la función json.Unmarshal().
	var genres []string

	err = json.Unmarshal(m["genres"], &genres)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("genres: %v\n", genres)
}
