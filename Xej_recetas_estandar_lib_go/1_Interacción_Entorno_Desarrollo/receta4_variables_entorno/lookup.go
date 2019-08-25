package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	key := "DB_CONN"
	connStr, ex := os.LookupEnv(key)
	if !ex {
		log.Printf("La variable env %s no está establecida.\n", key)
	}
	fmt.Println(connStr)
}
