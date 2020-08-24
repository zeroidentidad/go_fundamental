package database

import _ "github.com/go-sql-driver/mysql" // importar libreria compatible para database/sql

// Create hace una tabla llamada example y la pobla
func Create(db DB) error {
	// create table
	if _, err := db.Exec("CREATE TABLE example (name VARCHAR(20), created DATETIME)"); err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO example (name, created) values ("Zero", NOW())`); err != nil {
		return err
	}

	return nil
}
