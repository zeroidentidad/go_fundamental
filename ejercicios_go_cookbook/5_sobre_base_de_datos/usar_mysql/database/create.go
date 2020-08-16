package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // importar libreria compatible para database/sql
)

// Create crea una tabla llamada example y la pobla con registro
func Create(db *sql.DB) error {
	// crear base de datos
	if _, err := db.Exec("CREATE TABLE example (name VARCHAR(20), created DATETIME)"); err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO example (name, created) values ("Zero", NOW())`); err != nil {
		return err
	}

	return nil
}
