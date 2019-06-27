package db

import (
	"log"

	"github.com/go-pg/pg"
)

type Params struct {
	Param1 string
	Param2 string
}

//PlaceHolderDemo func exportada para paso de parametros
func PlaceHolderDemo(db *pg.DB) error {
	var valor string
	params := Params{
		Param1: "valor param1",
		Param2: "valor param2",
	}
	var query string = "select ?param2 "
	_, selectErr := db.Query(pg.Scan(&valor), query, params)
	if selectErr != nil {
		log.Printf("Error en select query. Detalle: %v\n", selectErr)
		return selectErr
	}
	log.Printf("Escaneo correcto, valor: %s\n", valor)

	return nil
}
