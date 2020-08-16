package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // importar libreria compatible para database/sql
)

// Exec toma una nueva conexión
// crea tablas en memoria y luego descarta
// y emite resultados de consulta
func Exec(db *sql.DB) error {
	// error no detectado en limpieza, pero
	// quiero limpiar
	defer db.Exec("DROP TABLE example")

	if err := Create(db); err != nil {
		return err
	}

	if err := Query(db); err != nil {
		return err
	}
	return nil
}
