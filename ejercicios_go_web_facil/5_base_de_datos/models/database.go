package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const username = "remoto"
const password = "x1234567"
const host = "localhost"
const port = 3306
const database = "gowebrestdb"

func CreateConnection() {
	if connection, err := sql.Open("mysql", generateURL()); err != nil {
		panic(err)
	} else {
		db = connection
		fmt.Println("Conexion ok")
	}
}

//Para la DB en general:

func generateURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
}

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func CloseConnection() {
	db.Close()
}

//Para las tablas en especifico:

func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, _ := db.Query(sql)
	return rows.Next()
}
