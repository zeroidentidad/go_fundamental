package nulls

import (
	"encoding/json"
	"fmt"
)

// json que tiene name pero no age
const (
	jsonBlob     = `{"name": "Zero"}`
	fulljsonBlob = `{"name":"Zero", "age":0}`
)

// Example es una estructura básica con campos de age y name.
type Example struct {
	Age  int    `json:"age,omitempty"`
	Name string `json:"name"`
}

// BaseEncoding muestra codificación y decodificación con tipos normales
func BaseEncoding() error {
	e := Example{}

	// considerar que no age = 0 age
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Regular Marshal, no age:", string(value))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshal, age = 0: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Regular Marshal, age = 0:", string(value))

	return nil
}
