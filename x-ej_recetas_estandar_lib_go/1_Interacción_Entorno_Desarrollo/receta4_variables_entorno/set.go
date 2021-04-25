package main

import (
	"log"
	"os"
)

func main() {

	key := "DB_CONN"
	// Establecer las variables env.
	os.Setenv(key, "postgres://usr:mydb@thewebapp.com/pg?sslmode=verify-full")

	val := GetEnvDefault(key, "postgres://usr:mydb@localhost/pg?sslmode=disable")

	log.Println("El valor es:" + val)

	os.Unsetenv(key)

	val = GetEnvDefault(key, "postgres://usr:mydb@127.0.0.1/pg?sslmode=disable")

	log.Println("El valor por default es:" + val)

}

// GetEnvDefault buscar las vars
func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}
