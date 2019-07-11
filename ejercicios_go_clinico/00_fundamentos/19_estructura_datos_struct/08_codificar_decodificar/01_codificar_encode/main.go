package main

import (
	"encoding/json"
	"os"
)

type persona struct {
	Nombre      string
	Apellido    string
	Edad        int
	noExportado int
}

func main() {
	p1 := persona{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p1) // salida formato json
}
