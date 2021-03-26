package web

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Opendb(dsn *string) *sql.DB {
	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		App.Logs().ErrorLog.Fatal(err)
		return nil
	}

	if err = db.Ping(); err != nil {
		App.Logs().ErrorLog.Fatal(err)
		return nil
	}

	return db
}
