package configuration

import (
	"encoding/json"
	"log"
	"os"

	"fmt"

	"github.com/jinzhu/gorm"
	// dialect implicito para usar con github.com/lib/pq
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// con pg, necesario: go get -u github.com/lib/pq
// estructura igual al config.json
type configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
	Ssl      string
}

func getConfiguration() configuration {
	var c configuration
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

// GetConnection obtiene una conexión a la bd
// gorm.Open("postgres", "host=mihost port=miport user=miuser dbname=midb password=mipassword")
// extra: sslmode=disable,verify-full
func GetConnection() *gorm.DB {
	c := getConfiguration()
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", c.Server, c.Port, c.User, c.Database, c.Password, c.Ssl)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
