package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

const callPostgre = "select * from format_name($1,$2,$3)"

// mysql example
// const callMySQL = "CALL simpleproc(?)"

type Result struct {
	Name     string
	Category int
}

func main() {
	db := createConnection()
	defer db.Close()
	r := Result{}

	if err := db.QueryRow(callPostgre, "Zero", "Identidad", 27).Scan(&r.Name); err != nil {
		panic(err)
	}
	fmt.Printf("Result is: %+v\n", r)
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
