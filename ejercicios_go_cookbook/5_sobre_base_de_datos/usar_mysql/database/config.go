package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // importar libreria compatible para database/sql
)

// Example para contener los resultados de las consultas
// ./models
type Example struct {
	Name    string
	Created *time.Time
}

// Setup configura y devuelve
// grupo de conexiones a base de datos
func Setup() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/gomysqldb?parseTime=true", "remoto", "x1234567", "192.168.0.100"))
	if err != nil {
		return nil, err
	}
	return db, nil
}
