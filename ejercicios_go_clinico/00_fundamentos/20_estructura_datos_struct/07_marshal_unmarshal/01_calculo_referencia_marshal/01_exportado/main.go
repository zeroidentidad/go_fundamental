package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre      string
	Apellido    string
	Edad        int
	noExportado int
}

func main() {
	p1 := persona{"James", "Bond", 20, 007}
	bs, _ := json.Marshal(p1) // _ omitir err
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
