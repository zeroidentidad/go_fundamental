package main

import (
	"log"

	"./db"
)

func main() {
	log.Printf("Listo\n")
	db.Conectar()
}
