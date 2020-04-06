package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// simulacion
const sel = "SELECT * FROM post p CROSS JOIN (SELECT 1 FROM generate_series(1,1000000)) tbl"

func main() {
	db := createConnection()
	defer db.Close()

	ctx, canc := context.WithTimeout(context.Background(), 20*time.Microsecond)
	rows, err := db.QueryContext(ctx, sel)
	canc() //cancelar la consulta
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		if rows.Err() != nil {
			fmt.Println(rows.Err())
			continue
		}
		count++
	}

	fmt.Printf("%d rows returned\n", count)

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
