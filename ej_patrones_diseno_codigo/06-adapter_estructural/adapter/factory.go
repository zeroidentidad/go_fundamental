package adapter

import (
	"../auto"
	"../bici"
)

func Factory(s string) Adapter {
	switch s {
	case "automovil":
		d := auto.Automovil{}
		return &AutomovilAdapter{&d}
	case "bicicleta":
		d := bici.Bicicleta{}
		return &BicicletaAdapter{&d}
	}

	return nil
}
