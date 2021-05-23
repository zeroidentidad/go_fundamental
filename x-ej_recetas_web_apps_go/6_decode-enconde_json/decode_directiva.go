package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	_json := `{"name": "alice", "age": 21}`

	var person struct {
		Name string `json:"name"`
		Age  int32  `json:"-"`
	}

	err := json.NewDecoder(strings.NewReader(_json)).Decode(&person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", person)
}
