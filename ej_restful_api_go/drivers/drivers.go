package drivers

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

//ConectarDB establece la conexion activa
func ConectarDB() *sql.DB {
	pgURL, err := pq.ParseURL(os.Getenv("DB_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgURL)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
