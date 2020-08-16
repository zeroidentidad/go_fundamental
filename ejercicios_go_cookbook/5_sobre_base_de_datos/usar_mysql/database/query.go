package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // importar libreria compatible para database/sql
)

// Query toma una nueva conexión
// crea tablas en memoria y luego las descarta
// y emite consultas
func Query(db *sql.DB) error {
	name := "Zero"
	rows, err := db.Query("SELECT name, created FROM example where name=?", name)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var e Example
		if err := rows.Scan(&e.Name, &e.Created); err != nil {
			return err
		}
		fmt.Printf("Results:\n\tName: %s\n\tCreated: %v\n", e.Name, e.Created)
	}
	return rows.Err()
}
