package main

import (
	"transaction-mysql/database"

	_ "github.com/go-sql-driver/mysql" // importar libreria compatible para database/sql
)

func main() {
	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	// esto no hará nada si el commit es exitoso
	defer tx.Rollback()

	if err := database.Exec(db); err != nil {
		panic(err)
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}
}
