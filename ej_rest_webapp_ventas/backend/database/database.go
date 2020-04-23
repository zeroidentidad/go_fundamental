package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	connString := "root:passwordx@tcp(localhost:3306)/northwind"

	dbConn, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}

	return dbConn
}
