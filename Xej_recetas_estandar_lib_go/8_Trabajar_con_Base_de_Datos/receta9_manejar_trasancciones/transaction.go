package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const selOne = "SELECT id,title,content FROM post WHERE ID = $1;"
const insert = "INSERT INTO post(ID,TITLE,CONTENT) VALUES (4,'Transaction Title','Transaction Content');"

type Post struct {
	ID      int
	Title   string
	Content string
}

func main() {
	db := createConnection()
	defer db.Close()

	tx, err := db.Begin() // init transacción
	if err != nil {
		panic(err)
	}
	_, err = tx.Exec(insert)
	if err != nil {
		panic(err)
	}
	p := Post{}
	// Consulta en otra sesión/transacción
	if err := db.QueryRow(selOne, 4).Scan(&p.ID, &p.Title, &p.Content); err != nil {
		fmt.Println("Got error for db.Query:" + err.Error())
	}
	fmt.Println(p)
	// Consulta dentro de la transacción
	if err := tx.QueryRow(selOne, 4).Scan(&p.ID, &p.Title, &p.Content); err != nil {
		fmt.Println("Got error for db.Query:" + err.Error())
	}
	fmt.Println(p)
	// Después de confirmar o revertir la transacción debe recrearse.
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}

	fmt.Println("\nTransaction with context")
	ctx, canc := context.WithCancel(context.Background())
	tx, err = db.BeginTx(ctx, &sql.TxOptions{sql.LevelReadUncommitted, false})
	if err != nil {
		panic(err)
	}
	_, err = tx.Exec(insert)
	if err != nil {
		panic(err)
	}
	canc()
	err = tx.Commit()
	if err != nil {
		fmt.Println("Error occured:" + err.Error())
	}

}

func createConnection() *sql.DB {
	connStr := "postgres://postgres:x1234567@192.168.0.100:5432/dbejemplox?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
