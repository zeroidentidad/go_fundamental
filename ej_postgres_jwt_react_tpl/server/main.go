package main

import (
	_ "database/sql"

	"postgres-jwt-react/db"
	"postgres-jwt-react/router"

	_ "github.com/lib/pq"
)

func init() {
	db.Connect()
}

func main() {
	r := router.SetupRouter()

	// Listen&Serve en 0.0.0.0:8081 -> *:8081
	r.Run(":8081")
}
