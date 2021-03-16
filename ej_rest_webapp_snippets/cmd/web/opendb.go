package web

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func opendb(dsn *string) *sql.DB {
	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		l.logs().errorLog.Fatal(err)
		return nil
	}

	if err = db.Ping(); err != nil {
		l.logs().errorLog.Fatal(err)
		return nil
	}

	return db
}
