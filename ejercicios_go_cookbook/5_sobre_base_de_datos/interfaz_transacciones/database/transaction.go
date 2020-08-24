package database

import "database/sql"

// DB es una interfaz que se satisface
// por un sql.DB o un sql.Transaction
type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// Transaction puede hacer cualquier cosa que Query
// más Commit, Rollback o Stmt
type Transaction interface {
	DB
	Commit() error
	Rollback() error
	Stmt(stmt *sql.Stmt) *sql.Stmt
}
