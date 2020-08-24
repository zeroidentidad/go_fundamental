package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Setup configura la base de datos junto con el número de
// conexiones de los grupos y más
func Setup() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/gomysqldb?parseTime=true", "remoto", "x1234567", "192.168.0.100"))
	if err != nil {
		return nil, err
	}

	// solo habrá 24 conexiones abiertas
	db.SetMaxOpenConns(24)

	// MaxIdleConns nunca puede ser inferior al
	// máximo de SetMaxOpenConns abiertas de lo contrario,
	// tendrá ese valor predeterminado
	db.SetMaxIdleConns(24)

	return db, nil
}
