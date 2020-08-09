package db

import (
	"database/sql"
	"fmt"
	"log"

	"postgres-jwt-react/config"
)

//instancia DB
var DB *sql.DB

//Conectar a DB
func Connect() {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME)

	db, _ := sql.Open("postgres", dbinfo)
	err := db.Ping()
	if err != nil {
		log.Fatal("Error: No se pudo establecer conexión con la base de datos")
	}
	DB = db

	log.Println("Conexion ok...")

	// Crear tabla "users" si no existe de no usarse el script en .psql
	CreateUsersTable()
}
