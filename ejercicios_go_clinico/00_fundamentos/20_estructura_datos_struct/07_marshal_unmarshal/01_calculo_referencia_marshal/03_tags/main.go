package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string
	Apellido string `json:"-"` // <- etiqueta omite campo
	Edad     int    `json:"wisdom score"`
}

func main() {
	p1 := persona{"James", "Bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}
