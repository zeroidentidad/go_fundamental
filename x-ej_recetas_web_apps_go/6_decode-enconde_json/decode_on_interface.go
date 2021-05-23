package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	// El array JSON contiene string, number, array y tipos boolean JSON
	_json := `["foo", true, 3.14, ["bar", false]]`

	// Decodificar JSON en un slice []interface{}
	var s []interface{}

	err := json.NewDecoder(strings.NewReader(_json)).Decode(&s)
	if err != nil {
		log.Fatal(err)
	}

	// El primer valor en el slice tendrá el tipo string de Go implicito,
	// el segundo tendrá el tipo bool de Go implicito, etc.
	// Se puede afirmar o ver si se reflejan los tipos, e imprimirlos
	fmt.Printf("item: 0; type: %T; value: %v\n", s[0], s[0].(string))
	fmt.Printf("item: 1; type: %T; value: %v\n", s[1], s[1].(bool))
	fmt.Printf("item: 2; type: %T; value: %v\n", s[2], s[2].(float64))
	fmt.Printf("item: 3; type: %T; value: %v\n", s[3], s[3].([]interface{}))
}
