package main

import (
	"log"

	"github.com/zeroidentidad/twittor/db"
	"github.com/zeroidentidad/twittor/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexion a BD")
		return
	}

	handlers.Handlers()
}
