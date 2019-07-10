package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func main() {
	p1 := persona{"James", "Bond", 20}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}
