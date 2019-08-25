package main

import (
	"log"
	"os" // <-
)

func main() {
	connStr := os.Getenv("DB_CONN")
	log.Printf("DB connection string: %s\n", connStr)
}
