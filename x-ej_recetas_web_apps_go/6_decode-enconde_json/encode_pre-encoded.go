package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	m := struct {
		Person string
	}{
		Person: `{"name": "Alice", "age": 21}`,
	}

	b, err := json.Marshal(m)
	if err != nil {
		log.Println(err.Error())
	}

	os.Stdout.Write(b)
	fmt.Println()

	m2 := struct {
		Person json.RawMessage
	}{
		Person: json.RawMessage(`{"name": "Alice", "age": 21}`),
	}

	b2, err := json.Marshal(m2)
	if err != nil {
		log.Println(err.Error())
	}

	os.Stdout.Write(b2)
	fmt.Println()
}
