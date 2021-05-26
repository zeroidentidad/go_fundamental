package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type myFloat float64

// Tiene un receptor de puntero.
func (f *myFloat) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", *f)), nil
}

func main() {
	f := myFloat(1.0 / 3.0)

	// Al codificar un valor, se utiliza el método MarshalJSON.
	js, err := json.Marshal(&f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", js)

	// Al codificar un valor, se ignora el método MarshalJSON.
	js, err = json.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", js)
}
