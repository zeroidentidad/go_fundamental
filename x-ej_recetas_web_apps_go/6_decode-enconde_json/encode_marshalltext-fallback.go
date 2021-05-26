package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type myFloat float64

func (f myFloat) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", f)), nil
}

func main() {
	f := myFloat(1.0 / 3.0)
	// var f2 myFloat = 1.0 / 3.0

	js, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", js)
}
