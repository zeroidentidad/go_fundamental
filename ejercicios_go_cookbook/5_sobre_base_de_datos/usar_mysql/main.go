package main

import (
	"use-mysql/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	if err := database.Create(db); err != nil {
		panic(err)
	}

	if err := database.Query(db); err != nil {
		panic(err)
	}

	// a lo ultimo eliminar example table
	if err := database.Exec(db); err != nil {
		panic(err)
	}
}
